package web

import "github.com/oktapascal/app-barayya/models/domain"

type UserResponses struct {
	Id         uint
	KodeLokasi string
	Password   string
	Role       string
}

func ConvertToUserResponse(user domain.User) UserResponses {
	return UserResponses{
		Id:         user.Id,
		KodeLokasi: user.Karyawan.KodeLokasi,
		Password:   user.Password,
		Role:       user.Role,
	}
}
