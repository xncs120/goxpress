package handlers

import (
	"net/http"
	"strconv"

	"goxpress/internal/security"
	"goxpress/models"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		db: db,
	}
}

func (h *UserHandler) GetUser(c *echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

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

func (h *UserHandler) UpdateUser(c *echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var request struct {
		OldPassword string `json:"old_password" validate:"required"`
		Username    string `json:"username"`
		Email       string `json:"email" validate:"email"`
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
		return c.JSON(http.StatusNotFound, map[string]any{"error": "Invalid Old Password"})
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
