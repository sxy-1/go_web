package models


type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"passowrd"`
}