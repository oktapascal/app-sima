package services

import (
	"context"
	"github.com/oktapascal/app-sima/models/web"
	"github.com/oktapascal/app-sima/repository"
	"github.com/oktapascal/app-sima/utils"
	"gorm.io/gorm"
)

type MenuServicesImpl struct {
	MenuRepository repository.MenuRepository
	Db             *gorm.DB
}

func NewMenuServicesImpl(menuRepository repository.MenuRepository, db *gorm.DB) *MenuServicesImpl {
	return &MenuServicesImpl{MenuRepository: menuRepository, Db: db}
}

func (services *MenuServicesImpl) GetAll(ctx context.Context) []web.MenuResponses {
	tx := services.Db.WithContext(ctx).Begin()
	defer utils.CommitRollback(tx)

	menus := services.MenuRepository.GetAll(tx)

	return web.ConvertToMenusResponses(menus)
}
