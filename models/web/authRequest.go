package web

type RegisterRequest struct {
	Username     string `json:"username" validate:"required"`
	Password     string `json:"password" validate:"required"`
	Role         string `json:"role" validate:"required"`
	IdLocation   string `json:"id_location" validate:"required"`
	Nik          string `json:"nik" validate:"required"`
	Name         string `json:"name" validate:"required"`
	RedirectView string `json:"redirect_view" validate:"required"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SessionRequest struct {
	IdUser    int
	Nik       string
	AuthToken string
}
