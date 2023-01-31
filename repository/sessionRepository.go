package repository

import (
	"github.com/oktapascal/app-sima/models/domain"
	"gorm.io/gorm"
)

type SessionRepository interface {
	Create(db *gorm.DB, session domain.Session)
	Delete(db *gorm.DB, session domain.Session)
}
