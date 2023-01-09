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

// UserServicesImpl is a struct that holds a UserRepository and a *gorm.DB.
type UserServicesImpl struct {
	UserRepository repository.UserRepository
	Db             *gorm.DB
}

// NewUserServicesImpl is a function that creates and returns a new *UserServicesImpl.
// It takes in a userRepository of type repository.UserRepository and a db of type *gorm.DB as input parameters.
func NewUserServicesImpl(userRepository repository.UserRepository, db *gorm.DB) *UserServicesImpl {
	// Return a new *UserServicesImpl with the given userRepository and db.
	return &UserServicesImpl{UserRepository: userRepository, Db: db}
}

// Save is a method on the *UserServicesImpl struct that saves a new user to the database.
// It takes in a context of type context.Context and a request of type web.RegisterRequest as input parameters.
func (services *UserServicesImpl) Save(context context.Context, request web.RegisterRequest) {
	// Begin a new transaction using the given context and the Db field of the *UserServicesImpl struct.
	tx := services.Db.WithContext(context).Begin()
	// Defer a call to the CommitRollback function, passing in the tx variable as an argument.
	defer utils.CommitRollback(tx)

	// Call the HashString function from the utils package, passing in the request.Password field as an argument.
	// Assign the return value to the passwordHash variable.
	// If the HashString function returns an error, call the PanicIfError function from the utils package and pass in the error as an argument.
	passwordHash, err := utils.HashString(request.Password)
	utils.PanicIfError(err)

	// Initialize a domain.User variable called user with the fields of the request,
	// as well as a hashed password and an empty domain.Karyawan struct with the KodeLokasi field set to the request.KodeLokasi field.
	user := domain.User{
		Username: request.Username,
		Password: passwordHash,
		Role:     request.Role,
		Karyawan: domain.Karyawan{
			KodeLokasi: request.KodeLokasi,
		},
	}

	// Call the Store method on the UserRepository field of the *UserServicesImpl struct, passing in the context, tx, and user variables as arguments.
	services.UserRepository.Store(context, tx, user)
}

// CheckUsername is a method on the *UserServicesImpl struct that checks if a given username is already in use.
// It takes in a context of type context.Context and a username of type string as input parameters.
// It returns a web.UserResponses.
func (services *UserServicesImpl) CheckUsername(context context.Context, username string) web.UserResponses {
	// Begin a new transaction using the given context and the Db field of the *UserServicesImpl struct.
	tx := services.Db.WithContext(context).Begin()
	// Defer a call to the CommitRollback function, passing in the tx variable as an argument.
	defer utils.CommitRollback(tx)

	// Call the CheckUsername method on the UserRepository field of the *UserServicesImpl struct,
	// passing in the context, tx, and username variables as arguments.
	// Assign the returned user and error to the user and err variables, respectively.
	user, err := services.UserRepository.CheckUsername(context, tx, username)

	// If the err variable is not nil, call the NewErrorUnprocessableEntity function from the exceptions package
	// and pass in the err.Error() and "username" as arguments. Panic with the returned error.
	if err != nil {
		panic(exceptions.NewErrorUnprocessableEntity(err.Error(), "username"))
	}

	// Return the result of calling the ConvertToUserResponse function from the web package, passing in the user variable as an argument.
	return web.ConvertToUserResponse(user)
}

// SaveSessionUser is a method on the *UserServicesImpl struct that saves a new session for a user in the database.
// It takes in a context of type context.Context and a request of type web.SessionRequest as input parameters.
func (services *UserServicesImpl) SaveSessionUser(context context.Context, request web.SessionRequest) {
	// Begin a new transaction using the given context and the Db field of the *UserServicesImpl struct.
	tx := services.Db.WithContext(context).Begin()
	// Defer a call to the CommitRollback function, passing in the tx variable as an argument.
	defer utils.CommitRollback(tx)

	// Initialize an empty domain.Session variable.
	var session domain.Session
	// Assign the IdUser field of the request variable to the IdUser field of the session variable.
	session.IdUser = request.IdUser
	// Assign the AuthToken field of the request variable to the AuthToken field of the session variable.
	session.AuthToken = request.AuthToken

	// Call the StoreSessionUser method on the UserRepository field of the *UserServicesImpl struct,
	// passing in the context, tx, and session variables as arguments.
	services.UserRepository.StoreSessionUser(context, tx, session)
}

// DestroySessionUser is a method on the *UserServicesImpl struct that removes a user's session from the database.
// It takes in a context of type context.Context and an authToken of type string as input parameters.
func (services *UserServicesImpl) DestroySessionUser(context context.Context, authToken string) {
	// Begin a new transaction using the given context and the Db field of the *UserServicesImpl struct.
	tx := services.Db.WithContext(context).Begin()
	// Defer a call to the CommitRollback function, passing in the tx variable as an argument.
	defer utils.CommitRollback(tx)

	// Call the DeleteSessionUser method on the UserRepository field of the *UserServicesImpl struct,
	// passing in the context, tx, and authToken variables as arguments.
	services.UserRepository.DeleteSessionUser(context, tx, authToken)
}

// GetSessionUser is a method on the *UserServicesImpl struct that retrieves a user from the database using a given authToken.
// It takes in a context of type context.Context and an authToken of type string as input parameters.
// It returns a web.UserResponses.
func (services *UserServicesImpl) GetSessionUser(context context.Context, authToken string) web.UserResponses {
	// Begin a new transaction using the given context and the Db field of the *UserServicesImpl struct.
	tx := services.Db.WithContext(context).Begin()
	// Defer a call to the CommitRollback function, passing in the tx variable as an argument.
	defer utils.CommitRollback(tx)

	// Call the GetUserBySession method on the UserRepository field of the *UserServicesImpl struct,
	// passing in the context, tx, and authToken variables as arguments.
	// Assign the returned domain.User and error to user and err variables, respectively.
	user, err := services.UserRepository.GetUserBySession(context, tx, authToken)
	// If the error is not nil, panic with a new exceptions.NewErrorNotFound error that has the error message as its message.
	if err != nil {
		panic(exceptions.NewErrorNotFound(err.Error()))
	}

	// Return the result of calling the ConvertToUserResponse function from the web package, passing in the user variable as an argument.
	return web.ConvertToUserResponse(user)
}

