package repository

import (
	"github.com/oktapascal/app-sima/models/domain"
	"gorm.io/gorm"
)

type SessionRepositoryImpl struct {
}

func NewSessionRepositoryImpl() *SessionRepositoryImpl {
	return &SessionRepositoryImpl{}
}

func (repository *SessionRepositoryImpl) Create(db *gorm.DB, session domain.Session) {
	var result = domain.Session{
		AuthToken: session.AuthToken,
		CreatedAt: session.CreatedAt,
		DeletedAt: session.DeletedAt,
		IdUser:    session.IdUser,
	}

	db.Create(&result)
}

func (repository *SessionRepositoryImpl) Delete(db *gorm.DB, session domain.Session) {
	db.Table("session").Where("auth_token = ? and id_user = ?", session.AuthToken, session.IdUser).Update("deleted_at", session.DeletedAt)
}
