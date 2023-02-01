package services

import (
	"context"
	"github.com/oktapascal/app-sima/exceptions"
	"github.com/oktapascal/app-sima/models/domain"
	"github.com/oktapascal/app-sima/models/web"
	"github.com/oktapascal/app-sima/repository"
	"github.com/oktapascal/app-sima/utils"
	"gorm.io/gorm"
	"time"
)

type AuthServicesImpl struct {
	UserRepository     repository.UserRepository
	EmployeeRepository repository.EmployeeRepository
	SessionRepository  repository.SessionRepository
	Db                 *gorm.DB
}

func NewAuthServicesImpl(userRepository repository.UserRepository, employeeRepository repository.EmployeeRepository, sessionRepository repository.SessionRepository, db *gorm.DB) *AuthServicesImpl {
	return &AuthServicesImpl{UserRepository: userRepository, EmployeeRepository: employeeRepository, SessionRepository: sessionRepository, Db: db}
}

func (services *AuthServicesImpl) Register(ctx context.Context, request web.RegisterRequest) {
	tx := services.Db.WithContext(ctx).Begin()
	defer utils.CommitRollback(tx)

	_, err := services.UserRepository.GetUser(tx, request.Username)
	if err == nil {
		panic(exceptions.NewErrorUnprocessableEntity("username sudah terdaftar", "username"))
	}

	hash, err := utils.HashString(request.Password)
	utils.PanicIfError(err)

	user := domain.User{
		Username:     request.Username,
		Password:     hash,
		Role:         request.Role,
		FlagInput:    1,
		FlagEdit:     1,
		FlagDelete:   1,
		RedirectView: request.RedirectView,
	}

	userId := services.UserRepository.Store(tx, user)

	employee := domain.Employee{
		Nik:         request.Nik,
		IdLocation:  request.IdLocation,
		Name:        request.Name,
		ActiveFlags: 1,
		IdUser:      userId,
	}

	services.EmployeeRepository.Store(tx, employee)
}

func (services *AuthServicesImpl) Login(ctx context.Context, request web.LoginRequest) web.UserResponses {
	tx := services.Db.WithContext(ctx).Begin()
	defer utils.CommitRollback(tx)

	user, err := services.UserRepository.GetUser(tx, request.Username)
	if err != nil {
		panic(exceptions.NewErrorUnprocessableEntity("username tidak terdaftar", "username"))
	}

	isValidPassword := utils.ValidatePassword(request.Password, user.Password)
	if !isValidPassword {
		panic(exceptions.NewErrorUnprocessableEntity("password tidak sesuai", "password"))
	}

	employee, err := services.EmployeeRepository.GetOne(tx, request.Username)
	if err != nil {
		panic(exceptions.NewErrorNotFound(err.Error()))
	}

	return web.ConvertToUserResponse(user, employee)
}

func (services *AuthServicesImpl) Logout(ctx context.Context, request web.SessionRequest) {
	tx := services.Db.WithContext(ctx).Begin()
	defer utils.CommitRollback(tx)

	now := time.Now()
	session := domain.Session{
		IdUser:    request.IdUser,
		AuthToken: request.AuthToken,
		DeletedAt: &now,
	}

	services.SessionRepository.Delete(tx, session)
}

func (services *AuthServicesImpl) StoreSessionUser(ctx context.Context, request web.SessionRequest) {
	tx := services.Db.WithContext(ctx).Begin()
	defer utils.CommitRollback(tx)

	session := domain.Session{
		AuthToken: request.AuthToken,
		CreatedAt: time.Now(),
		IdUser:    request.IdUser,
	}

	services.SessionRepository.Create(tx, session)
}

func (services *AuthServicesImpl) GetUserProfile(ctx context.Context, nik string) web.UserProfileResponses {
	tx := services.Db.WithContext(ctx).Begin()
	defer utils.CommitRollback(tx)

	user, err := services.UserRepository.GetUser(tx, nik)
	if err != nil {
		panic(exceptions.NewErrorNotFound(err.Error()))
	}

	employee, err := services.EmployeeRepository.GetOne(tx, nik)
	if err != nil {
		panic(exceptions.NewErrorNotFound(err.Error()))
	}

	return web.ConvertToUserProfileResponse(user, employee)
}

func (services *AuthServicesImpl) UpdateUserProfile(ctx context.Context, request web.UpdateUserProfileRequest) {
	tx := services.Db.WithContext(ctx).Begin()
	defer utils.CommitRollback(tx)

	employee, err := services.EmployeeRepository.GetOne(tx, request.Nik)
	if err != nil {
		panic(exceptions.NewErrorNotFound(err.Error()))
	}

	employee.Name = request.Name
	employee.Address = &request.Address
	employee.NoTelp = &request.NoTelp
	employee.Email = &request.Email

	services.EmployeeRepository.Update(tx, employee)
}

func (services *AuthServicesImpl) UploadUserPhoto(ctx context.Context, request web.UploadUserPhoto) {
	tx := services.Db.WithContext(ctx).Begin()
	defer utils.CommitRollback(tx)

	user, err := services.UserRepository.GetUser(tx, request.Nik)
	if err != nil {
		panic(exceptions.NewErrorNotFound(err.Error()))
	}

	user.Photo = &request.Photo

	services.UserRepository.StorePhoto(tx, user)
}
