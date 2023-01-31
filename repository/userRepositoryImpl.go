package repository

import (
	"database/sql"
	"errors"
	"github.com/oktapascal/app-sima/models/domain"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Store(db *gorm.DB, user domain.User) int {
	var result = domain.User{
		Username:     user.Username,
		Password:     user.Password,
		Role:         user.Role,
		RedirectView: user.RedirectView,
		FlagInput:    user.FlagInput,
		FlagEdit:     user.FlagEdit,
		FlagDelete:   user.FlagDelete,
		Photo:        user.Photo,
		Provider:     user.Provider,
		IDProvider:   user.IDProvider,
	}
	db.Create(&result)

	var userId int
	row := db.Table("user a").Select("id_user").Order("id_user desc").Row()
	_ = row.Scan(&userId)

	return userId
}

func (repository *UserRepositoryImpl) Update(db *gorm.DB, user domain.User) {
	//TODO implement me
	panic("implement me")
}

func (repository *UserRepositoryImpl) GetUser(db *gorm.DB, username string) (domain.User, error) {
	var user domain.User

	row := db.Table("user a").Select("a.id_user, a.username, a.password, a.role, a.redirect_view, a.flag_input, a.flag_edit, a.flag_delete, a.photo, a.provider, a.id_provider").
		Where("a.username = ?", username).Row()

	err := row.Scan(&user.IdUser, &user.Username, &user.Password, &user.Role, &user.RedirectView, &user.FlagInput, &user.FlagEdit, &user.FlagDelete, &user.Photo, &user.Provider, &user.IDProvider)

	if errors.Is(err, sql.ErrNoRows) {
		return user, errors.New("data not found")
	}

	return user, nil
}
