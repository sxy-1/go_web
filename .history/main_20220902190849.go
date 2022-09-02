package main

import (
	"shuxiaoyan.com/m/dao"
	models "shuxiaoyan.com/m/model"
)

func main() {
	user := mode.User{
		Username: "sxy",
		Password: 123456,
	}
	dao.Mgr.AddUser(&user)
}
