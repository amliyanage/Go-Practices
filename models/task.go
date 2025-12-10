package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Title       string         `gorm:"not null" json:"title"`
	Description string         `json:"description"`
	Completed   bool           `gorm:"default:false" json:"completed"`
	UserID      uint           `gorm:"not null" json:"user_id"`
	User        User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
}
