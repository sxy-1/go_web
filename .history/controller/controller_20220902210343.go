package controller

import (
	"github.com/gin-gonic/gin"
	"shuxiaoyan.com/m/dao"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := model.User{
		Username : username,
		Password : password,
	}

	dao.Mgr.Register()

}
