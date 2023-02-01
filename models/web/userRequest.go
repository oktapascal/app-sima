package web

type UpdateUserProfileRequest struct {
	Nik     string
	Name    string `json:"name" validate:"required"`
	NoTelp  string `json:"no_telp" validate:"required,number,min=11,max=12"`
	Address string `json:"address" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
}

type UploadUserPhoto struct {
	Nik   string
	Photo string
}
