package dao

type Manager interface {
	AddUser(user *model.User)
}

type Manager struct{
	db 
}

func main() {

}
