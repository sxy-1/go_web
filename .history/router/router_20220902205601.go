package router

import "github.com/gin-gonic/gin"

func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("tmplates/*")
	e.Static("/assets","./assets")
	e.
	e.Run()

}