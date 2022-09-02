package main

import "shuxiaoyan.com/m/dao"
import "shuxiaoyan.com/m/model"

func main() {
	user := model.User{
		Username: "sxy",
		Password: 123456,
	}
	dao.Mgr.AddUser(&user)
}
