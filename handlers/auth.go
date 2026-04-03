package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/xncs120/goxpress/config"
	"github.com/xncs120/goxpress/internal/security"
	"github.com/xncs120/goxpress/models"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type AuthHandler struct {
	db *gorm.DB
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{
		db: db,
	}
}

func (h *AuthHandler) Health(c *echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{
		"status": "live",
	})
}

func (h *AuthHandler) Register(c *echo.Context) error {
	var request struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"error": "Invalid JSON format"})
	}
	if err := c.Validate(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"error": "Missing required fields"})
	}

	var existingUser models.User
	if err := h.db.Where("username = ? OR email = ?", request.Username, request.Email).First(&existingUser).Error; err == nil {
		return c.JSON(http.StatusConflict, map[string]any{"error": "Username or Email already exists"})
	} else if err != gorm.ErrRecordNotFound {
		return c.JSON(http.StatusInternalServerError, map[string]any{"error": "Database error"})
	}

	hashedPassword, err := security.HashPassword(request.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"error": "Fail to encrypt password"})
	}

	user := models.User{
		Username: request.Username,
		Email:    request.Email,
		Password: string(hashedPassword),
	}

	if err := h.db.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"error": "Fail to create user"})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"id":         user.ID,
		"username":   user.Username,
		"email":      user.Email,
		"role":       user.FormatRoleName(),
		"created_at": user.CreatedAt,
	})
}

func (h *AuthHandler) Login(c *echo.Context) error {
	var request struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"error": "Invalid JSON format"})
	}
	if err := c.Validate(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"error": "Missing required fields"})
	}

	var user models.User
	if err := h.db.Where("username = ? AND status = 1", request.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]any{"error": "Username and Password not matched"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]any{"error": "Database error"})
	}

	passwordVerification := security.VerifyPassword(request.Password, user.Password)
	if !passwordVerification {
		return c.JSON(http.StatusNotFound, map[string]any{"error": "Username and Password not matched"})
	}

	expireAt := time.Now().Add(time.Minute * time.Duration(config.App.ExpMins))
	token, err := security.JWTGenerateToken(strconv.Itoa(int(user.ID)), user.Username, user.Email, expireAt)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{"error": "Fail to generate token"})
	}

	return c.JSON(http.StatusOK, map[string]any{"token": token, "expired_at": expireAt})
}

func (h *AuthHandler) GetAuthUser(c *echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := security.JWTParseToken(tokenString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"error": "Invalid token"})
	}

	id := claims["id"].(string)
	var user models.User
	if err := h.db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]any{"error": "User not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]any{"error": "Database error"})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"id":         user.ID,
		"username":   user.Username,
		"email":      user.Email,
		"role":       user.FormatRoleName(),
		"created_at": user.CreatedAt,
	})
}

func (h *AuthHandler) UpdateAuthUser(c *echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := security.JWTParseToken(tokenString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"error": "Invalid token"})
	}

	id := claims["id"].(string)
	var request struct {
		OldPassword string `json:"old_password" validate:"required"`
		Username    string `json:"username"`
		Email       string `json:"email" validate:"omitempty,email"`
		Password    string `json:"password"`
	}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"error": "Invalid JSON format"})
	}
	if err := c.Validate(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"error": "Missing required fields"})
	}

	var user models.User
	if err := h.db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]any{"error": "User not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]any{"error": "Database error"})
	}

	passwordVerification := security.VerifyPassword(request.OldPassword, user.Password)
	if !passwordVerification {
		return c.JSON(http.StatusNotFound, map[string]any{"error": "Old Password not matched"})
	}

	updateData := make(map[string]any)
	if request.Username != "" {
		updateData["username"] = request.Username
	}
	if request.Email != "" {
		updateData["email"] = request.Email
	}
	if request.Password != "" {
		hashedPassword, err := security.HashPassword(request.Password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{"error": "Fail to encrypt password"})
		}
		updateData["password"] = string(hashedPassword)
	}

	if err := h.db.Model(&models.User{}).Where("id = ?", id).Updates(updateData).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"error": "Fail to update user, username or email might already exist"})
	}
	var updatedUser models.User
	h.db.First(&updatedUser, id)

	return c.JSON(http.StatusOK, map[string]any{
		"id":         updatedUser.ID,
		"username":   updatedUser.Username,
		"email":      updatedUser.Email,
		"role":       updatedUser.FormatRoleName(),
		"created_at": user.CreatedAt,
	})
}
