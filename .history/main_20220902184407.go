package main

import "gorm.io/gorm"

func main() {
	_, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go?charset=utf8&parseTime=True&loc=Local")
	if err != nil {

		panic(err)
	}
}
