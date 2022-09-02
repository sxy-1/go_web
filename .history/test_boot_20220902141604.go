package main

import "github.com/gin-gonic/gin"


func 


func main() {
	e := gin.Default()
	e.Static("/assets","./assets")
	e.LoadHTMLGlob("templates/*")

}
