package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"passowrd"`
}

type Post struct{
	gorm.Model
	Content string `gorm:"type:text"`
	Writer string
}

type Post struct{
	gorm.Model
	Content string `gorm:"type:text"`
	Writer string
}