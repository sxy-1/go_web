package main

import (
	"shuxiaoyan.com/m/dao"
	"shux
)

func main() {
	user := model.User{
		Username: "sxy",
		Password: 123456,
	}
	dao.Mgr.AddUser(&user)
}
