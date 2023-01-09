package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/oktapascal/app-barayya/models/domain"
	"github.com/oktapascal/app-barayya/utils"
	"gorm.io/gorm"
	"time"
)

// UserRepositoryImpl is a struct that implements the UserRepository interface.
// This struct is responsible for performing CRUD (Create, Read, Update, Delete) operations on user data.
type UserRepositoryImpl struct {
}

// NewUserRepositoryImpl is a method used to instantiate the UserRepositoryImpl struct.
// This method will return a pointer to the newly instantiated UserRepositoryImpl struct.
func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

// Store is a method in the UserRepository interface that is responsible for saving user data to the database.
// It accepts 4 parameters:
// - ctx: the context used to manage the request.
// - db: the *gorm.DB object used to execute queries on the database.
// - user: a domain.User object containing the user data to be saved.
func (repository *UserRepositoryImpl) Store(ctx context.Context, db *gorm.DB, user domain.User) {
	// Insert user data into the 'user' table.
	db.Create(&domain.User{
		Username: user.Username,
		Password: user.Password,
		Role:     user.Role,
	})

	// Retrieve the latest ID inserted into the 'user' table.
	var idUser uint
	row := db.Table("user").Select("id").Order("id desc").Limit(1).Row()
	_ = row.Scan(&idUser)

	// Create a 'karyawan' object using the newly acquired user ID and the location code from the 'user' object.
	karyawan := domain.Karyawan{
		IdUser:     idUser,
		KodeLokasi: user.Karyawan.KodeLokasi,
	}

	// Insert the 'karyawan' object into the 'karyawan' table.
	db.Create(&karyawan)
}

// CheckUsername is a method in the UserRepository interface that is responsible for checking if a given username is registered in the database.
// It accepts 3 parameters:
// - ctx: the context used to manage the request.
// - db: the *gorm.DB object used to execute queries on the database.
// - username: a string containing the username to be checked.
func (repository *UserRepositoryImpl) CheckUsername(ctx context.Context, db *gorm.DB, username string) (domain.User, error) {
	var user domain.User

	// Execute a SELECT query with an inner join on the 'karyawan' table to retrieve user data with the specified username.
	row := db.Table("user").Select("user.id, user.password, user.role, karyawan.kode_lokasi").
		Joins("inner join karyawan on user.id=karyawan.id_user").
		Where("user.username = ?", username).Row()

	err := row.Scan(&user.Id, &user.Password, &user.Role, &user.Karyawan.KodeLokasi)

	// If no data is found, return an error with the message "username not registered".
	if errors.Is(err, sql.ErrNoRows) {
		return user, errors.New("username tidak terdaftar")
	}

	// Return the user data and a nil error if data is found.
	return user, nil
}

// StoreSessionUser is a method in the UserRepository interface that is responsible for saving user session data to the database.
// It accepts 3 parameters:
// - ctx: the context used to manage the request.
// - db: the *gorm.DB object used to execute queries on the database.
// - session: a domain.Session object containing the user session data to be saved.
func (repository *UserRepositoryImpl) StoreSessionUser(ctx context.Context, db *gorm.DB, session domain.Session) {
	// Insert user session data into the 'session' table.
	db.Create(&domain.Session{
		IdUser:    session.IdUser,
		AuthToken: session.AuthToken,
		CreatedAt: time.Now(),
	})
}

// DeleteSessionUser is a method in the UserRepository interface that is responsible for deleting user session data from the database.
// It accepts 3 parameters:
// - ctx: the context used to manage the request.
// - db: the *gorm.DB object used to execute queries on the database.
// - authToken: a string containing the authentication token of the user session to be deleted.
func (repository *UserRepositoryImpl) DeleteSessionUser(ctx context.Context, db *gorm.DB, authToken string) {
	// Set the 'deleted_at' field of the user session with the specified authentication token to the current time.
	db.Model(&domain.Session{}).Where("auth_token = ?", authToken).Update("deleted_at", time.Now())
}

// GetUserBySession is a method on the *UserRepositoryImpl struct that retrieves a domain.User from the database using a given authToken.
// It takes in a ctx of type context.Context, a db of type *gorm.DB, and an authToken of type string as input parameters.
// It returns a domain.User and an error.
func (repository *UserRepositoryImpl) GetUserBySession(ctx context.Context, db *gorm.DB, authToken string) (domain.User, error) {
	// Initialize an empty domain.User variable
	var user domain.User

	// Create a row variable using a gorm.DB method called Table, which allows you to specify the table name to be used in the query.
	// Call the Select method on row, which specifies the columns that should be returned in the query.
	// Call the Joins method on row multiple times, each time specifying an inner join with a different table.
	// Call the Where method on row, which specifies a condition for the query. Pass in the authToken input parameter as the condition.
	row := db.Table("user").Select("user.id, user.role, karyawan.kode_lokasi").
		Joins("inner join karyawan on user.id=karyawan.id_user").
		Joins("inner join session on user.id=session.id_user").
		Where("session.auth_token = ?", authToken).Row()

	// Check if the Err method of row returns gorm.ErrRecordNotFound.
	// If it does, return the empty user variable and a new errors.New error with the message "authentication token invalid".
	if errors.Is(row.Err(), gorm.ErrRecordNotFound) {
		return user, errors.New("authentication token invalid")
	}

	// If the Err method of row does not return gorm.ErrRecordNotFound, call the Scan method on row,
	// passing in pointers to the fields of the user variable as arguments.
	// If the Scan method returns an error, call the PanicIfError function from the utils package and pass in the error as an argument.
	err := row.Scan(&user.Id, &user.Role, &user.Karyawan.KodeLokasi)
	if err != nil {
		utils.PanicIfError(err)
	}

	// Return the user variable and a nil error.
	return user, nil
}
