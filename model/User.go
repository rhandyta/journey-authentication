package model

type User struct {
	Id        int
	Firstname string
	Lastname  string
	Age       int
	Email     string
	Password  string
}

type UserResponse struct {
	Id        int    `json:"id"`
	Firstname string `json:"string"`
	Lastname  string `json:"string"`
	Age       int    `json:"int"`
	Email     string `json:"email"`
}
