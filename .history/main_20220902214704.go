package main

import (
	"shuxiaoyan.com/m/dao"
	"shuxiaoyan.com/m/model"
	"shuxiaoyan.com/m/router"
)

func main() {
	/* user := model.User{
		Username: "张三",
		Password: "123456",
	}
	dao.Mgr.Register(&user) */
	router.Start
}
