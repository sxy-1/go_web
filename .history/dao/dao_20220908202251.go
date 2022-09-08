package dao

import (
	"fmt"
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
	GetAllPost(username string) []model.Post
	ChangePost(id int,content string) 
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
	db.AutoMigrate(&model.Post{})
}

func (mgr *manager) Register(user *model.User) {
	mgr.db.Create(user)
}

func (mgr *manager) Login(username string) model.User {
	var user model.User
	mgr.db.Where("username=?", username).First(&user)
	return user

}

//问题操作
func (mgr *manager) AddPost(post *model.Post) {
	mgr.db.Create(post)
}

func (mgr *manager) GetAllPost(username string) []model.Post {
	var posts = make([]model.Post, 10)

	mgr.db.Where("writer=?", username).Find(&posts)

	fmt.Printf("%v",posts)

	return posts
}

/* type UserInfo struct {
	Username string
} */

/* func (mgr *manager) GetAllPost(username string) []UserInfo {
	var userinfo []UserInfo
	mgr.db.Where("writer=?", username).Find(&userinfo)
	fmt.Println(userinfo)
	return userinfo
} */


func (mgr *manager)ChangePost(id int,content string) {
	 mgr.db.Table("user").Where("id=?", 1).Update("id", 1234)
	 mgr.db.W
}