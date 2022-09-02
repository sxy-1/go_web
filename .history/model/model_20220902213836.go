package model

import (
	"github.com/sirupsen/logrus/hooks/writer"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"passowrd"`
}

type Post struct{
	gorm.Model
	Content string `gorm:"type:text`
	writer
}