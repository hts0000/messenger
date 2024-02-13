package model

import "gorm.io/gorm"

type User struct {
	ID                 uint64 `gorm:"primarykey"`
	Email              string
	Username           string
	Password           []byte
	Salt               []byte
	AvatarUrl          string
	BackgroundImageUrl string
	Signature          string

	gorm.Model
}

func (*User) TableName() string {
	return "user"
}
