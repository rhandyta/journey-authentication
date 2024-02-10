package model

import "github.com/golang-jwt/jwt/v5"

type AuthUser struct {
	User  User
	Token string
}

type Login struct {
	Email    string `json:"email" validate:"required_if=Username '' "`
	Username string `json:"username" validate:"required_if=Email '' "`
	Password string `json:"password" validate:"required"`
}

type AuthUserClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}
