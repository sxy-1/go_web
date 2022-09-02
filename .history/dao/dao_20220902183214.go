package dao

import "fmt"

package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"passowrd"`
}
func main() {

}
