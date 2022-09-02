package main

import "github.com/gin-gonic/gin"


func Gostatic(e *gin.Context){
	e.HTML(200,"test_static.html",)
}


func main() {
	e := gin.Default()
	e.Static("/assets","./assets")
	e.LoadHTMLGlob("templates/*")

}
