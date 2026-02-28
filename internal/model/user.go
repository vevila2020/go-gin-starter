package model

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type User struct {
	Model
	Username string `gorm:"size:64;uniqueIndex;not null" json:"username"`
	Email    string `gorm:"size:128;uniqueIndex;not null" json:"email"`
	Password string `gorm:"size:256;not null" json:"-"`
	Nickname string `gorm:"size:64" json:"nickname"`
	Avatar   string `gorm:"size:256" json:"avatar"`
	Status   int    `gorm:"default:1" json:"status"`
}

func (User) TableName() string {
	return "user"
}
