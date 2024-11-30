package user

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/xncs120/goxpress/internal/base/auth"
	"github.com/xncs120/goxpress/internal/base/config"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) CreateUser(c echo.Context) error {
	user := new(Model)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	var existingUser Model
	if err := h.db.Where("username = ? OR email = ?", user.Username, user.Email).First(&existingUser).Error; err == nil {
		return c.JSON(http.StatusConflict, echo.Map{"error": "Username or Email already exists"})
	} else if err != gorm.ErrRecordNotFound {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Database error"})
	}

	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Fail to encrypt password"})
	}
	user.Password = hashedPassword

	if err := h.db.Create(user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Fail to create user"})
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *Handler) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var user Model
	result := h.db.First(&user, id)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Database error"})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var existingUser Model
	result := h.db.First(&existingUser, id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Database error"})
	}

	var request struct {
		Password  string  `json:"password"`
		Username  *string `json:"username"`
		Email     *string `json:"email"`
		Status    *uint   `json:"status"`
		NPassword *string `json:"new_password"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid JSON input"})
	}

	passwordVerification := auth.VerifyPassword(existingUser.Password, request.Password)
	if !passwordVerification {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "User Email and Password not matched"})
	}

	if request.Username != nil {
		existingUser.Username = *request.Username
	}
	if request.Email != nil {
		existingUser.Email = *request.Email
	}
	if request.Status != nil {
		existingUser.Status = *request.Status
	}
	if request.NPassword != nil {
		hashedPassword, err := auth.HashPassword(*request.NPassword)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error hashing password"})
		}
		existingUser.Password = hashedPassword
	}

	if err := h.db.Save(&existingUser).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, existingUser)
}

func (h *Handler) GenerateToken(c echo.Context) error {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid JSON input"})
	}

	var user Model
	if err := h.db.Where("email = ? AND status = 1", request.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Database error"})
	}

	passwordVerification := auth.VerifyPassword(user.Password, request.Password)
	if !passwordVerification {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "User Email and Password not matched"})
	}

	expireAt := time.Now().Add(time.Hour * time.Duration(config.Secret.JwtExpireHrs))
	token, err := auth.JWTGenerateToken(strconv.Itoa(int(user.ID)), user.Email, expireAt)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Fail to generate token"})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token, "expired_at": expireAt})
}
