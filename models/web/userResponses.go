package web

import "github.com/oktapascal/app-barayya/models/domain"

type UserResponses struct {
	Id         uint   `json:"id_user"`
	KodeLokasi string `json:"kode_lokasi"`
	Password   string `json:"-"`
	Role       string `json:"role"`
}

func ConvertToUserResponse(user domain.User) UserResponses {
	return UserResponses{
		Id:         user.Id,
		KodeLokasi: user.Karyawan.KodeLokasi,
		Password:   user.Password,
		Role:       user.Role,
	}
}
