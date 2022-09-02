package controller

import (
	"github.com/gin-gonic/gin"
	"shuxiaoyan.com/m/dao"
	"shuxiaoyan.com/m/model"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := model.User{
		Username : username,
		Password : password,
	}

	dao.Mgr.Register(&user)

}

func Login(c *gin.Context){
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(username)
	u := dao.Mgr.Login(username)

	if u.Username==""{
		c.HTML(200,"login.html","用户名不存在！")
		fmt.Println("")
	}
}