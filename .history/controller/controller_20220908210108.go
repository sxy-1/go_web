package controller

import (
	"encoding/json"
	"fmt"
	"strconv"

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
		fmt.Println("出问题了")
		panic(err)

	}
	//c.String(200,append(u, -1))

	c.JSON(200, string(res))

}

func AddPost(c *gin.Context) {
	// localhost:8080/getpost?username=sxy
	username := c.Query("username")
	content := c.Query("content")

	u := new(model.Post)
	u.Writer = username
	u.Content = content
	dao.Mgr.AddPost(u)
	c.String(200, "添加成功！")
}

func ChangePost(c *gin.Context) {
	id,_ := strconv.Atoi(c.Query("id"))
	content := c.Query("content")

	dao.Mgr.ChangePost(id, content)
	c.String(200, "修改成功！")
}


//回答
func GetPost(c *gin.Context) {
	// localhost:8080/getpost?username=sxy

	username := c.Query("username")

	u := dao.Mgr.GetAllPost(username)

	res, err := json.Marshal(u)
	if err != nil {
		fmt.Println("出问题了")
		panic(err)

	}
	//c.String(200,append(u, -1))

	c.JSON(200, string(res))

}

func AddAnswer(c *gin.Context) {
	// localhost:8080/getpost?username=sxy
	username := c.Query("username")
	content := c.Query("content")
	post_id := c.Query("post_id")
	u := new(model.Answer)
	u.Content = content
	u.Writer = username
	u.Post_id = post_id
	dao.Mgr.AddPost(u)
	c.String(200, "添加成功！")
}

func ChangePost(c *gin.Context) {
	id,_ := strconv.Atoi(c.Query("id"))
	content := c.Query("content")

	dao.Mgr.ChangePost(id, content)
	c.String(200, "修改成功！")
}