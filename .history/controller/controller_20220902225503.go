package controller

import (
	"fmt"

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
	// localhost:8080/login?username=sxy&password=123
	username := c.Query("username")
	password := c.Query("password")
	fmt.Println(username)
	u := dao.Mgr.Login(username)

	if u.Username==""{
		c.HTML(200,"login.html","用户名不存在！")
		c.str
		fmt.Println("用户名不存在！")
	}else{
		if u.Password !=password{
			fmt.Println("密码错误")
			c.HTML(200,"login.html","密码错误")
		}else {
			fmt.Println("登陆成功")
			c.Redirect(301,"/")
		}

	}
}