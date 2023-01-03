package services

import (
	"context"
	"github.com/oktapascal/app-barayya/exceptions"
	"github.com/oktapascal/app-barayya/models/domain"
	"github.com/oktapascal/app-barayya/models/web"
	"github.com/oktapascal/app-barayya/repository"
	"github.com/oktapascal/app-barayya/utils"
	"gorm.io/gorm"
)

type UserServicesImpl struct {
	UserRepository repository.UserRepository
	Db             *gorm.DB
}

func NewUserServicesImpl(userRepository repository.UserRepository, db *gorm.DB) *UserServicesImpl {
	return &UserServicesImpl{UserRepository: userRepository, Db: db}
}

func (services *UserServicesImpl) Save(context context.Context, request web.RegisterRequest) {
	tx := services.Db.WithContext(context).Begin()
	defer utils.CommitRollback(tx)

	passwordHash, err := utils.HashString(request.Password)
	utils.PanicIfError(err)

	user := domain.User{
		Username: request.Username,
		Password: passwordHash,
		Role:     request.Role,
		Karyawan: domain.Karyawan{
			KodeLokasi: request.KodeLokasi,
		},
	}

	services.UserRepository.Store(context, tx, user)
}

func (services *UserServicesImpl) CheckUsername(context context.Context, username string) web.UserResponses {
	tx := services.Db.WithContext(context).Begin()
	defer utils.CommitRollback(tx)

	user, err := services.UserRepository.CheckUsername(context, tx, username)

	if err != nil {
		panic(exceptions.NewErrorUnprocessableEntity(err.Error(), "username"))
	}

	return web.ConvertToUserResponse(user)
}

func (services *UserServicesImpl) SaveSessionUser(context context.Context, request web.SessionRequest) {
	tx := services.Db.WithContext(context).Begin()
	defer utils.CommitRollback(tx)

	var session domain.Session
	session.IdUser = request.IdUser
	session.AuthToken = request.AuthToken

	services.UserRepository.StoreSessionUser(context, tx, session)
}

func (services *UserServicesImpl) DestroySessionUser(context context.Context, authToken string) {
	tx := services.Db.WithContext(context).Begin()
	defer utils.CommitRollback(tx)

	services.UserRepository.DeleteSessionUser(context, tx, authToken)
}

func (services *UserServicesImpl) GetSessionUser(context context.Context, authToken string) web.UserResponses {
	tx := services.Db.WithContext(context).Begin()
	defer utils.CommitRollback(tx)

	user, err := services.UserRepository.GetUserBySession(context, tx, authToken)
	if err != nil {
		panic(exceptions.NewErrorNotFound(err.Error()))
	}

	return web.ConvertToUserResponse(user)
}
