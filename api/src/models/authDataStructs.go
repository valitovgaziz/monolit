package models

import (
	"github.com/golang-jwt/jwt/v4"
)

type Credentials struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}

type Claims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
}

type ShortCredentials struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
