package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID int64 `gorm:"primaryKey;autoIncrement"`

	CreatedBy int64
	CreatedAt time.Time `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedBy int64
	UpdatedAt time.Time      `gorm:"type:timestamp;default:current_timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserDevice struct {
	ID int64 `gorm:"primaryKey;autoIncrement"`

	UserID int64
	User   User `gorm:"foreignKey:UserID"`

	DeviceToken   string
	Device        string
	UserAgent     string
	Platform      string
	ClientVersion string

	CreatedBy int64
	CreatedAt time.Time `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedBy int64
	UpdatedAt time.Time      `gorm:"type:timestamp;default:current_timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
