package router

import (
	"github.com/gin-gonic/gin"
	"shuxiaoyan.com/m/controller"
)

func Start() {
	e := gin.Default()
	/* e.LoadHTMLGlob("tmplates/*")
	e.Static("/assets", "./assets") */
	e.GET("/register", controller.Register)
	e.GET("/login", controller.Login)

	e.GET("/addpost", controller.AddPost) // localhost:8080/addpost?username=sxy&content=1+1=
	e.GET("/getpost", controller.GetPost) // localhost:8080/getpost?username=sxy
	e.GET("/changepost", controller.ChangePost) // localhost:8080/changepost?id=1&content=2+2=


	e.GET("/addanswer", controller.AddAnswer) // localhost:8080/addanswer?username=sxy&content=2&post_id=1
	e.GET("/getanswer", controller.GetAnswer) // localhost:8080/getanswer?username=sxy
	e.GET("/changeanswer", controller.ChangeAnswer) // localhost:8080/changeanswer?id=1&content=5


	e.GET("/addcomment", controller.AddAnswer) // localhost:8080/addanswer?username=sxy&content=2&post_id=1
	e.GET("/getcomment", controller.GetAnswer) // localhost:8080/getanswer?username=sxy
	e.GET("/changecomment", controller.ChangeAnswer) // localhost:8080/changeanswer?id=1&content=5

	

	e.Run()

}
