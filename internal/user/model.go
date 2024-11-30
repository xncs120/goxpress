package user

import "time"

type Model struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"unique;not null" json:"username"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	Status    uint      `gorm:"not null;default:1" json:"status"`
	CreatedAt time.Time `gorm:"null" json:"created_at"`
}

func (Model) TableName() string {
	return "users"
}
