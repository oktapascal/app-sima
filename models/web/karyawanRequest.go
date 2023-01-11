package web

type KaryawanUpdateRequest struct {
	Nik    string `json:"nik" validate:"required,number,len=5"`
	Nama   string `json:"nama" validate:"required"`
	Alamat string `json:"alamat" validate:"required"`
	NoTelp string `json:"no_telp" validate:"required,number,min=11,max=12"`
	Email  string `json:"email" validate:"required,email"`
	IdUser uint
}
