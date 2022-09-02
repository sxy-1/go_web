package main

import (
	"shuxiaoyan.com/m/dao"
	models "shuxiaoyan.com/m/model"
)

func main() {
	user := models.User{
		Username: "sxy",
		Password: "123456",
	}
	dao.Mgr.AddUser(&user)
}
