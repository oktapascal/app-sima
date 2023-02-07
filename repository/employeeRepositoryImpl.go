package repository

import (
	"database/sql"
	"errors"
	"github.com/oktapascal/app-sima/models/domain"
	"gorm.io/gorm"
)

type EmployeeRepositoryImpl struct {
}

func NewEmployeeRepositoryImpl() *EmployeeRepositoryImpl {
	return &EmployeeRepositoryImpl{}
}

func (repository *EmployeeRepositoryImpl) Store(db *gorm.DB, employee domain.Employee) {
	var result = domain.Employee{
		Nik:         employee.Nik,
		IdLocation:  employee.IdLocation,
		Name:        employee.Name,
		Address:     employee.Address,
		NoTelp:      employee.NoTelp,
		Email:       employee.Email,
		ActiveFlags: employee.ActiveFlags,
		IdArea:      employee.IdArea,
		IdFm:        employee.IdFm,
		IdBm:        employee.IdBm,
		IdPosition:  employee.IdPosition,
		IdUser:      employee.IdUser,
	}

	db.Create(&result)
}

func (repository *EmployeeRepositoryImpl) GetOne(db *gorm.DB, nik string) (domain.Employee, error) {
	var employee domain.Employee

	row := db.Table("employee a").Select(`a.id_employee, a.nik, a.id_location, a.name, a.address, 
	a.email, a.active_flags, a.id_position, a.id_area, a.id_fm, a.id_bm, a.no_telp`).
		Where("a.nik = ?", nik).Row()

	err := row.Scan(&employee.IdEmployee, &employee.Nik, &employee.IdLocation, &employee.Name, &employee.Address, &employee.Email,
		&employee.ActiveFlags, &employee.IdPosition, &employee.IdArea, &employee.IdFm, &employee.IdBm, &employee.NoTelp)

	if errors.Is(err, sql.ErrNoRows) {
		return employee, errors.New("data employee not found")
	}

	return employee, nil
}

func (repository *EmployeeRepositoryImpl) Update(db *gorm.DB, employee domain.Employee) {
	db.Table("employee").Where("nik", employee.Nik).Updates(domain.Employee{
		Name:    employee.Name,
		Address: employee.Address,
		NoTelp:  employee.NoTelp,
		Email:   employee.Email,
	})
}
