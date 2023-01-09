package web

import "github.com/oktapascal/app-barayya/models/domain"

type UserResponses struct {
	Id         uint   `json:"id_user"`
	KodeLokasi string `json:"kode_lokasi"`
	Password   string `json:"-"`
	Role       string `json:"role"`
}

type UserProfileResponses struct {
	Id         uint    `json:"id_user"`
	Username   string  `json:"username"`
	KodeLokasi string  `json:"kode_lokasi"`
	Role       string  `json:"role"`
	Nik        *string `json:"nik"`
	Nama       *string `json:"nama"`
	Alamat     *string `json:"alamat"`
	NoTelp     *string `json:"no_telp"`
	Email      *string `json:"email"`
	Foto       *string `json:"foto"`
}

func ConvertToUserResponse(user domain.User) UserResponses {
	return UserResponses{
		Id:         user.Id,
		KodeLokasi: user.Karyawan.KodeLokasi,
		Password:   user.Password,
		Role:       user.Role,
	}
}

func ConvertToUserProfileResponse(user domain.User) UserProfileResponses {
	return UserProfileResponses{
		Id:         user.Id,
		Username:   user.Username,
		KodeLokasi: user.Karyawan.KodeLokasi,
		Role:       user.Role,
		Nik:        user.Karyawan.Nik,
		Nama:       user.Karyawan.Nama,
		Alamat:     user.Karyawan.Alamat,
		NoTelp:     user.Karyawan.NoTelp,
		Email:      user.Karyawan.Email,
		Foto:       user.Karyawan.Foto,
	}
}
