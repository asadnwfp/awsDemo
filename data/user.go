package data

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User

func AddUser(u User) []User {
	users = append(users, u)
	return users
}

func GetUsers() []User {
	return users
}
