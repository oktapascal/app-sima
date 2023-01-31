package web

import "github.com/golang-jwt/jwt/v4"

type JwtClaims struct {
	IdUser     int    `json:"id_user"`
	Nik        string `json:"nik"`
	IdLocation string `json:"id_location"`
	Role       string `json:"role"`
	jwt.RegisteredClaims
}
