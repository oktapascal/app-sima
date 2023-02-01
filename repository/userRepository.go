package repository

import (
	"github.com/oktapascal/app-sima/models/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Store(db *gorm.DB, user domain.User) int
	Update(db *gorm.DB, user domain.User)
	GetUser(db *gorm.DB, username string) (domain.User, error)
	StorePhoto(db *gorm.DB, user domain.User)
}
