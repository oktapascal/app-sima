package web

import "github.com/oktapascal/app-sima/models/domain"

type UserResponses struct {
	IdUser     int    `json:"-"`
	Nik        string `json:"nik"`
	IdLocation string `json:"id_location"`
	Password   string `json:"-"`
	Role       string `json:"role"`
}

type UserProfileResponses struct {
	Username   string  `json:"username"`
	IdLocation string  `json:"id_location"`
	Role       string  `json:"role"`
	Nik        string  `json:"nik"`
	Name       string  `json:"name"`
	Address    *string `json:"address"`
	NoTelp     *string `json:"no_telp"`
	Email      *string `json:"email"`
	Photo      *string `json:"photo"`
}

func ConvertToUserResponse(user domain.User, employee domain.Employee) UserResponses {
	return UserResponses{
		IdUser:     user.IdUser,
		Nik:        employee.Nik,
		IdLocation: employee.IdLocation,
		Password:   user.Password,
		Role:       user.Role,
	}
}

func ConvertToUserProfileResponse(user domain.User, employee domain.Employee) UserProfileResponses {
	return UserProfileResponses{
		Username:   user.Username,
		IdLocation: employee.IdLocation,
		Role:       user.Role,
		Nik:        employee.Nik,
		Name:       employee.Name,
		Address:    employee.Address,
		NoTelp:     employee.NoTelp,
		Email:      employee.Email,
		Photo:      user.Photo,
	}
}
