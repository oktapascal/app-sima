package web

type UpdateUserProfileRequest struct {
	Nik     string
	Nama    string `json:"nama" validate:"required"`
	NoTelp  string `json:"no_telp" validate:"required,number,min=11,max=12"`
	Jabatan string `json:"jabatan" validate:"required"`
	Alamat  string `json:"alamat" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
}
