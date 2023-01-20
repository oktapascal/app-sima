package web

import "github.com/golang-jwt/jwt/v4"

type JwtClaims struct {
	Nik        string `json:"nik"`
	KodeLokasi string `json:"kode_lokasi"`
	Role       string `json:"role"`
	jwt.RegisteredClaims
}
