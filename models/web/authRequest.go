package web

type RegisterRequest struct {
	Username   string `json:"username" validate:"required"`
	Password   string `json:"password" validate:"required"`
	Role       string `json:"role" validate:"required"`
	KodeLokasi string `json:"kode_lokasi" validate:"required"`
	Nik        string `json:"nik" validate:"required"`
	Nama       string `json:"nama" validate:"required"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SessionRequest struct {
	Nik       string
	AuthToken string
}
