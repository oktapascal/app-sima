package web

import "github.com/oktapascal/app-sima/models/domain"

type UserResponses struct {
	Nik        string `json:"nik"`
	KodeLokasi string `json:"kode_lokasi"`
	Password   string `json:"-"`
	Role       string `json:"role"`
}

type UserProfileResponses struct {
	Username   string  `json:"username"`
	KodeLokasi string  `json:"kode_lokasi"`
	Role       string  `json:"role"`
	Nik        string  `json:"nik"`
	Nama       string  `json:"nama"`
	Alamat     *string `json:"alamat"`
	NoTelp     *string `json:"no_telp"`
	Email      *string `json:"email"`
	Foto       *string `json:"foto"`
}

func ConvertToUserResponse(user domain.User) UserResponses {
	return UserResponses{
		Nik:        user.Nik,
		KodeLokasi: user.KodeLokasi,
		Password:   user.Password,
		Role:       user.Role,
	}
}

func ConvertToUserProfileResponse(user domain.User) UserProfileResponses {
	return UserProfileResponses{
		Username:   user.Username,
		KodeLokasi: user.KodeLokasi,
		Role:       user.Role,
		Nik:        user.Nik,
		Nama:       user.Nama,
		Alamat:     user.Alamat,
		NoTelp:     user.NoTelp,
		Email:      user.Email,
		Foto:       user.Foto,
	}
}
