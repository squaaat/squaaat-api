package model

import "time"

type User struct {
	ID int64 `gorm:"primaryKey;autoIncrement"`

	CreatedAt time.Time
	CreatedBy int64
	UpdatedAt time.Time
	UpdatedBy int64
}

type UserDevice struct {
	ID int64 `gorm:"primaryKey;autoIncrement"`

	User          User
	Device        string
	UserAgent     string
	Platform      string
	ClientVersion string
	DeviceToken   string
}
