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
	dsn := "root:root@tcp(127.0.0.1:3306)/test_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 自动迁移
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})

	// 增
	db.Create(&User{
		Name: "张三",
		Age:  18,
		Addr: "北京市",
		Pic:  "/static/img.png",
	})

	// 查
	var user User
	db.First(&user)
	fmt.Println(user) // {1 张三 18 北京市 /static/img.png}

	// 改
	user.Name = "lisi"
	db.Save(&user)
	fmt.Println(user) // {1 lisi 18 北京市 /static/img.png}

	// 删
	db.Delete(&user)
}
