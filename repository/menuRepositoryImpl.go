package repository

import (
	"github.com/oktapascal/app-sima/models/domain"
	"github.com/oktapascal/app-sima/utils"
	"gorm.io/gorm"
)

type MenuRepositoryImpl struct {
}

func NewMenuRepositoryImpl() *MenuRepositoryImpl {
	return &MenuRepositoryImpl{}
}

func (repository *MenuRepositoryImpl) GetAll(db *gorm.DB) []domain.Menu {
	var menus []domain.Menu

	rows, err := db.Table("klp_menu").Select(`kode_klp_menu, name_klp_menu, menu`).Rows()
	//noinspection GoUnhandledErrorResult
	defer rows.Close()
	utils.PanicIfError(err)

	for rows.Next() {
		var menu domain.Menu

		err := rows.Scan(&menu.KodeKlpMenu, &menu.NameKlpMenu, &menu.ListMenuJson)
		utils.PanicIfError(err)

		menus = append(menus, menu)
	}

	return menus
}

func (repository *MenuRepositoryImpl) Store(db *gorm.DB, menu domain.Menu) {
	//TODO implement me
	panic("implement me")
}

func (repository *MenuRepositoryImpl) Update(db *gorm.DB, menu domain.Menu) {
	//TODO implement me
	panic("implement me")
}
