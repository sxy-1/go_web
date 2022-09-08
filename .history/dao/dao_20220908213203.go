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
	GetAllPost(username string) []model.Post
	ChangePost(id int, content string)

	//回答
	AddAnswer(answer *model.Answer)
	GetAnswer(username string) []model.Answer
	ChangeAnswer(id int, content string)

	//评论
	AddComment(comment *model.Comment)
	GetComment(username string) []model.Comment
	ChangeComment(id int, content string)
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
	db.AutoMigrate(&model.Answer{})
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

	/* 	fmt.Printf("%v", posts)
	 */
	return posts
}

func (mgr *manager) ChangePost(id int, content string) {
	mgr.db.Table("posts").Where("id=?", id).Update("content", content)
}



//回答操作
func (mgr *manager) AddAnswer(answer *model.Answer) {
	mgr.db.Create(answer)
}

func (mgr *manager) GetAnswer(username string) []model.Answer {
	var answer = make([]model.Answer, 10)
	mgr.db.Table("answers").Where("writer=?", username).Find(&answer)
	return answer
}

func (mgr *manager) ChangeAnswer(id int, content string) {
	mgr.db.Table("answers").Where("id=?", id).Update("content", content)
}
