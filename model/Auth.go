package model

type AuthUser struct {
	User  *User
	Token string
}

type Login struct {
	Identifier string `json:"email" binding:"required`
	Password   string `json:"password" binding:"required`
}

type Registration struct {
	Firstname string `json:"firstname" Form:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Age       uint8  `json:"age" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}
