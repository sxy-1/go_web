package main

import "github.com/gin-gonic/gin"


func Gostatic(e *gin.Context){
	e.h
}


func main() {
	e := gin.Default()
	e.Static("/assets","./assets")
	e.LoadHTMLGlob("templates/*")

}
