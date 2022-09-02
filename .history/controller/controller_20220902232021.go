package controller

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"shuxiaoyan.com/m/dao"
	"shuxiaoyan.com/m/model"
)

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	user := model.User{
		Username: username,
		Password: password,
	}

	dao.Mgr.Register(&user)

}

func Login(c *gin.Context) {
	// localhost:8080/login?username=sxy&password=123
	username := c.Query("username")
	password := c.Query("password")
	fmt.Println(username)
	u := dao.Mgr.Login(username)

	if u.Username == "" {
		c.String(200, "用户名不存在！")
	} else {
		if u.Password != password {
			c.String(200, "密码错误")
		} else {
			c.String(200, "登录成功")
		}

	}
}

func GetPost(c *gin.Context) {
	// localhost:8080/getpost?username=sxy
	username := c.Query("username")

	u := dao.Mgr.GetAllPost(username)
	res, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	//c.String(200,append(u, -1))
	c.JSON(200, res)
}

func AddPost(c *gin.Context) {
	// localhost:8080/getpost?username=sxy
	username := c.Query("username")

	u := dao.Mgr.GetAllPost(username)
	res, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	//c.String(200,append(u, -1))
	c.JSON(200, res)
}
