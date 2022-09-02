package main

func main(){
	dbConn, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test1?charset=utf8&parseTime=True&loc=Local")

if err != nil {

		panic(err)

}
}