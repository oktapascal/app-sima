package services

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/oktapascal/app-barayya/exceptions"
	"github.com/oktapascal/app-barayya/models/domain"
	"github.com/oktapascal/app-barayya/models/web"
	"github.com/oktapascal/app-barayya/repository"
	"github.com/oktapascal/app-barayya/utils"
	"time"
)

type UserServicesImpl struct {
	//UserRepository is an implementation of the UserRepository interface
	UserRepository repository.UserRepository
	//Db is a pointer to a Firestore client
	Db *firestore.Client
}

// NewUserServicesImpl returns a new instance of UserServicesImpl
func NewUserServicesImpl(userRepository repository.UserRepository, db *firestore.Client) *UserServicesImpl {
	return &UserServicesImpl{UserRepository: userRepository, Db: db}
}

func (services *UserServicesImpl) Register(ctx context.Context, request web.RegisterRequest) {
	// Retrieve a database client from the `services` struct.
	client := services.Db

	// Check if the provided username already exists in the database
	// using the `CheckDuplicateUsername` method from the `UserRepository` struct.
	err := services.UserRepository.CheckDuplicateUsername(ctx, client, request.Username)

	// If the username already exists, raise an error
	if err != nil {
		panic(exceptions.NewErrorUnprocessableEntity("username sudah terdaftar", "username"))
	}

	// Hash the provided password using the `HashString` function from the `utils` package.
	passwordHash, err := utils.HashString(request.Password)
	// Raise an error if there's any error in HashString
	utils.PanicIfError(err)

	// Create a `User` struct with the provided username, hashed password, and role
	user := domain.User{
		Username:   request.Username,
		Password:   passwordHash,
		Role:       request.Role,
		KodeLokasi: request.KodeLokasi,
		Nik:        request.Nik,
		Nama:       request.Nama,
	}

	// Store the user struct in the database
	// using the `Store` method from the `UserRepository` struct.
	services.UserRepository.Store(ctx, client, user)
}

func (services *UserServicesImpl) Login(ctx context.Context, request web.LoginRequest) web.UserResponses {
	// Get the pointer to the firestore client
	client := services.Db

	// Check if the provided username exists in the "users" collection
	user, err := services.UserRepository.CheckUsername(ctx, client, request.Username)
	if err != nil {
		// Panic if the username is not registered
		panic(exceptions.NewErrorUnprocessableEntity("username tidak terdaftar", "username"))
	}

	// Validate the provided password with the one stored in the user document
	isValidPassword := utils.ValidatePassword(request.Password, user.Password)
	if !isValidPassword {
		// Panic if the passwords do not match
		panic(exceptions.NewErrorUnprocessableEntity("password tidak valid", "password"))
	}

	// Convert the user data to the format of web.UserResponses and return the converted data
	return web.ConvertToUserResponse(user)
}

func (services *UserServicesImpl) StoreSessionUser(ctx context.Context, request web.SessionRequest) {
	// Get the pointer to the firestore client
	client := services.Db

	// Create a user struct with the NIK provided in the request
	user := domain.User{
		Nik: request.Nik,
	}

	// Create a session struct with the auth token and current timestamp provided in the request
	session := domain.Session{
		AuthToken: request.AuthToken,
		CreatedAt: time.Now(),
	}

	// Store the session for the user in the "users" collection
	services.UserRepository.StoreSessionUser(ctx, client, session, user)
}

func (services *UserServicesImpl) Logout(ctx context.Context, request web.SessionRequest) {
	// Get the pointer to the firestore client
	client := services.Db

	// Create a user struct with the NIK provided in the request
	user := domain.User{Nik: request.Nik}

	// Get the current time
	timeNow := time.Now()

	// Create a session struct with the auth token and the current time as deleted at
	session := domain.Session{
		AuthToken: request.AuthToken,
		DeletedAt: &timeNow,
	}

	// Delete the session for the user in the "users" collection
	services.UserRepository.DeleteSessionUser(ctx, client, session, user)
}

func (services *UserServicesImpl) GetUserProfile(ctx context.Context, nik string) web.UserProfileResponses {
	// Get the pointer to the firestore client
	client := services.Db

	// Get the user's profile from the "users" collection using the provided NIK
	user, err := services.UserRepository.GetUserProfile(ctx, client, nik)
	if err != nil {
		// Panic if the user is not found
		panic(exceptions.NewErrorNotFound(err.Error()))
	}

	// Convert the user data to the format of web.UserProfileResponses and return the converted data
	return web.ConvertToUserProfileResponse(user)
}

func (services *UserServicesImpl) UpdateUserProfile(ctx context.Context, request web.UpdateUserProfileRequest) {
	// Get the pointer to the firestore client
	client := services.Db

	// Create a user struct with the data from the request
	user := domain.User{
		Nik:     request.Nik,
		Nama:    request.Nama,
		Alamat:  &request.Alamat,
		NoTelp:  &request.NoTelp,
		Email:   &request.Email,
		Jabatan: &request.Jabatan,
	}

	// Update the user's profile using the UserRepository struct
	services.UserRepository.Update(ctx, client, user)
}
