package web

import "github.com/golang-jwt/jwt/v4"

type JwtClaims struct {
	IdUser     uint   `json:"id_user"`
	KodeLokasi string `json:"kode_lokasi"`
	Role       string `json:"role"`
	jwt.RegisteredClaims
}
