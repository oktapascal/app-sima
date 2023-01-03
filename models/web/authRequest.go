package web

type RegisterRequest struct {
	Username   string `json:"username" validate:"required"`
	Password   string `json:"password" validate:"required"`
	Role       string `json:"role" validate:"required"`
	KodeLokasi string `json:"kode_lokasi" validate:"required"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SessionRequest struct {
	IdUser    uint
	AuthToken string
}
