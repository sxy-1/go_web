package controller

import "github.com/gin-gonic/gin"

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := model.User

}
