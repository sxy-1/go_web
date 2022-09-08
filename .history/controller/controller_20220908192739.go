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
	fmt.Println("%v",res)
	res, err := json.Marshal(u)
	if err != nil {
		fmt.Println("出问题了")
		panic(err)
		
	}
	//c.String(200,append(u, -1))
	fmt.Println("%v",res)
	c.JSON(200, res)

/* 	c.String(200,username) */
}

func AddPost(c *gin.Context) {
	// localhost:8080/getpost?username=sxy
	username := c.Query("username")
	content :=c.Query("content")

	u:= new(model.Post)
	u.Writer=username
	u.Content=content
	dao.Mgr.AddPost(u)
	c.String(200,"添加成功！")
}
