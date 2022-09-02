package router

import ("github.com/gin-gonic/gin"
"shuxiaoyan.com/m/model")
func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("tmplates/*")
	e.Static("/assets", "./assets")
	e.POST(" /register", controller.Register)
	e.GET("/register", controller.GoRegister)
	e.GET("/login", controller.GoLogin)
	e.POST(" /1ogin", controller.Login)
	e.GET("!", controller.Index)
	e.Run()
	e.Run()

}
