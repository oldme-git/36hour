package model

type Id uint64

type User struct {
	Id       Id     `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}
