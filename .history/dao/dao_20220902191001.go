package dao

import (
	"shuxiaoyan.com/m/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Manager interface {
	AddUser(user *models.User)
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to init db :", err)
	}
	Mgr = &manager{db: db}
	db.AutoMigrate(&models.User{})
}

func (mgr *manager) AddUSER(user *model.User) {
	mgr.db.Create(user)
}

func main() {

}
