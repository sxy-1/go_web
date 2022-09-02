package dao

import "gorm.io/gorm"

type Manager interface {
	AddUser(user *model.User)
}

type Manager struct {
	db *gorm.DB
}



func main() {

}
