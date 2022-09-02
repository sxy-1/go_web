package dao

import (
	"log"

	"shuxiaoyan.com/m/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Manager interface {
	//用户
	Register(user *model.User)
	Login(username string) model.User

	//问题
	AddPost(post *model.Post)
	GetAllPost() []
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
	db.AutoMigrate(&model.User{})
}

func (mgr *manager) Register(user *model.User) {
	mgr.db.Create(user)
}

func (mgr *manager) Login(username string) model.User {
	var user model.User
	mgr.db.Where("username=?",username).First(&user)
	return user

}
