package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/oktapascal/app-barayya/models/domain"
	"github.com/oktapascal/app-barayya/utils"
	"gorm.io/gorm"
	"time"
)

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Store(ctx context.Context, db *gorm.DB, user domain.User) {
	db.Create(&domain.User{
		Username: user.Username,
		Password: user.Password,
		Role:     user.Role,
	})

	var idUser uint
	row := db.Table("user").Select("id").Order("id desc").Limit(1).Row()
	_ = row.Scan(&idUser)

	karyawan := domain.Karyawan{
		IdUser:     idUser,
		KodeLokasi: user.Karyawan.KodeLokasi,
	}

	db.Create(&karyawan)
}

func (repository *UserRepositoryImpl) CheckUsername(ctx context.Context, db *gorm.DB, username string) (domain.User, error) {
	var user domain.User

	row := db.Table("user").Select("user.id, user.password, user.role, karyawan.kode_lokasi").
		Joins("inner join karyawan on user.id=karyawan.id_user").
		Where("user.username = ?", username).Row()

	err := row.Scan(&user.Id, &user.Password, &user.Role, &user.Karyawan.KodeLokasi)

	if errors.Is(err, sql.ErrNoRows) {
		return user, errors.New("username tidak terdaftar")
	}

	return user, nil
}

func (repository *UserRepositoryImpl) StoreSessionUser(ctx context.Context, db *gorm.DB, session domain.Session) {
	db.Create(&domain.Session{
		IdUser:    session.IdUser,
		AuthToken: session.AuthToken,
		CreatedAt: time.Now(),
	})
}

func (repository *UserRepositoryImpl) DeleteSessionUser(ctx context.Context, db *gorm.DB, authToken string) {
	db.Model(&domain.Session{}).Where("auth_token = ?", authToken).Update("deleted_at", time.Now())
}

func (repository *UserRepositoryImpl) GetUserBySession(ctx context.Context, db *gorm.DB, authToken string) (domain.User, error) {
	var user domain.User

	row := db.Table("user").Select("user.id, user.role, karyawan.kode_lokasi").
		Joins("inner join karyawan on user.id=karyawan.id_user").
		Joins("inner join session on user.id=session.id_user").
		Where("session.auth_token = ?", authToken).Row()

	if errors.Is(row.Err(), gorm.ErrRecordNotFound) {
		return user, errors.New("authentication token invalid")
	}

	err := row.Scan(&user.Id, &user.Role, &user.Karyawan.KodeLokasi)
	if err != nil {
		utils.PanicIfError(err)
	}

	return user, nil
}
