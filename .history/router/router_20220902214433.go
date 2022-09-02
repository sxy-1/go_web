package router

import (
	"github.com/gin-gonic/gin"
	"shuxiaoyan.com/m/controller"
)

func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("tmplates/*")
	e.Static("/assets", "./assets")
	e.GET("/register", controller.Register)
/* 	e.POST(" /register", controller.Register)
	
	e.GET("/login", controller.GoLogin)
	e.POST(" /1ogin", controller.Login)
	e.GET("!", controller.Index) */
	e.Run()

}
