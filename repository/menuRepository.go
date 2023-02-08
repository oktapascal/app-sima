package repository

import (
	"github.com/oktapascal/app-sima/models/domain"
	"gorm.io/gorm"
)

type MenuRepository interface {
	GetAll(db *gorm.DB) []domain.Menu
	Store(db *gorm.DB, menu domain.Menu)
	Update(db *gorm.DB, menu domain.Menu)
}
