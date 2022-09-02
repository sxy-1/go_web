package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id   int
	Name string
	Age  int
	Addr string
	Pic  string
}

func main() {
	
}
