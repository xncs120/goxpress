package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"unique;not null" json:"username"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	Role      uint      `gorm:"not null;default:1" json:"role"`
	Status    uint      `gorm:"not null;default:1" json:"status"`
	CreatedAt time.Time `gorm:"null" json:"created_at"`
}

func (u *User) FormatRoleName() string {
	switch u.Role {
	case 0:
		return "suadmin"
	case 2:
		return "admin"
	default:
		return "user"
	}
}
