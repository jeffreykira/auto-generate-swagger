package model

import "github.com/dgrijalva/jwt-go"

//Claims ...
type Claims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}

//Signin ...
type Signin struct {
	Username string `json:"username" binding:"required" example:"jeffrey"`
	Password string `json:"password" binding:"required" example:"kira"`
}
