package model

import (
	"time"
)

type User struct {
	ID int64 `gorm:"primaryKey;autoIncrement"`

	CreatedBy int64
	CreatedAt time.Time
	UpdatedBy int64
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
}

type UserDevice struct {
	ID int64 `gorm:"primaryKey;autoIncrement"`

	UserID int64
	User          User `gorm:"foreignKey:UserID"`
	Device        string
	UserAgent     string
	Platform      string
	ClientVersion string
	DeviceToken   string

	CreatedBy int64
	CreatedAt time.Time
	UpdatedBy int64
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
}
