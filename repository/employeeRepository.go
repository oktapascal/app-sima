package repository

import (
	"github.com/oktapascal/app-sima/models/domain"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Store(db *gorm.DB, employee domain.Employee)
	GetOne(db *gorm.DB, nik string) (domain.Employee, error)
	Update(db *gorm.DB, employee domain.Employee)
}
