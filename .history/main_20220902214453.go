package main

import (
	"shuxiaoyan.com/m/dao"
	"shuxiaoyan.com/m/model"
)

func main() {
	user := model.User{
		Username: "sxy",
		Password: "123456",
	}
	dao.Mgr.Register(&user)
}
