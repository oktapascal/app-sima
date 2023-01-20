package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"github.com/oktapascal/app-barayya/models/domain"
	"github.com/oktapascal/app-barayya/utils"
	"google.golang.org/api/iterator"
	"reflect"
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

func (repository *UserRepositoryImpl) Store(ctx context.Context, db *firestore.Client, user domain.User) {
	// Set the data from the user struct in the "users" collection, using the user's NIK as the document ID
	_, err := db.Collection("users").Doc(user.Nik).Set(ctx, domain.User{
		Username:   user.Username,
		Password:   user.Password,
		Role:       user.Role,
		KodeLokasi: user.KodeLokasi,
		Nik:        user.Nik,
		Nama:       user.Nama,
		Alamat:     user.Alamat,
		NoTelp:     user.NoTelp,
		Email:      user.Email,
		Foto:       user.Foto,
	})
	// Panic if there is an error during the set operation
	utils.PanicIfError(err)
}

func (repository *UserRepositoryImpl) CheckUsername(ctx context.Context, db *firestore.Client, username string) (domain.User, error) {
	//Create an empty user variable
	var user domain.User

	//Query the "users" collection for documents where the "username" field is equal to the provided "username" string
	iter := db.Collection("users").Where("username", "==", username).Documents(ctx)

	//Iterate through the returned documents
	for {
		document, err := iter.Next()
		//Break the loop if no more documents are returned
		if err == iterator.Done {
			break
		}
		//Populate the user variable with the data from the document
		_ = document.DataTo(&user)
	}

	//Check if the user variable is zero value, indicating that the username does not exist
	if reflect.ValueOf(user).IsZero() {
		return user, errors.New("username tidak terdaftar")
	}
	//Return the user if it is not zero value
	return user, nil
}

func (repository *UserRepositoryImpl) StoreSessionUser(ctx context.Context, db *firestore.Client, session domain.Session, user domain.User) {
	_, _, err := db.Collection("users").Doc(user.Nik).Collection("sessions").Add(ctx, domain.Session{
		AuthToken: session.AuthToken,
		CreatedAt: session.CreatedAt,
		DeletedAt: session.DeletedAt,
	})
	utils.PanicIfError(err)
}

func (repository *UserRepositoryImpl) DeleteSessionUser(ctx context.Context, db *firestore.Client, session domain.Session, user domain.User) {
	// Initialize a variable to hold the session ID
	var id string

	// Get the iterator for the matching documents in the "sessions" subcollection where the auth token matches the one provided
	iter := db.Collection("users").Doc(user.Nik).Collection("sessions").Where("auth_token", "==", session.AuthToken).Documents(ctx)

	// Iterate through the documents
	for {
		document, err := iter.Next()
		if err == iterator.Done {
			// Break the loop if there are no more documents
			break
		}

		// Get the ID of the session document
		id = document.Ref.ID
	}

	// Update the deleted_at field in the session document with the value provided
	db.Collection("users").Doc(user.Nik).Collection("sessions").Doc(id).Update(ctx, []firestore.Update{
		{
			Path:  "deleted_at",
			Value: session.DeletedAt,
		},
	})
}

// GetUserBySession is a method on the *UserRepositoryImpl struct that retrieves a domain.User from the database using a given authToken.
// It takes in a ctx of type context.Context, a db of type *firestore.Client, and an authToken of type string as input parameters.
// It returns a domain.User and an error.
func (repository *UserRepositoryImpl) GetUserBySession(ctx context.Context, db *firestore.Client, authToken string) (domain.User, error) {
	// Initialize an empty domain.User variable
	var user domain.User

	// Create a row variable using a firestore.Client method called Table, which allows you to specify the table name to be used in the query.
	// Call the Select method on row, which specifies the columns that should be returned in the query.
	// Call the Joins method on row multiple times, each time specifying an inner join with a different table.
	// Call the Where method on row, which specifies a condition for the query. Pass in the authToken input parameter as the condition.
	//row := db.Table("user").Select("user.id, user.role, karyawan.kode_lokasi").
	//	Joins("inner join karyawan on user.id=karyawan.id_user").
	//	Joins("inner join session on user.id=session.id_user").
	//	Where("session.auth_token = ?", authToken).Row()

	// Check if the Err method of row returns gorm.ErrRecordNotFound.
	// If it does, return the empty user variable and a new errors.New error with the message "authentication token invalid".
	//if errors.Is(row.Err(), gorm.ErrRecordNotFound) {
	//	return user, errors.New("authentication token invalid")
	//}

	// If the Err method of row does not return gorm.ErrRecordNotFound, call the Scan method on row,
	// passing in pointers to the fields of the user variable as arguments.
	// If the Scan method returns an error, call the PanicIfError function from the utils package and pass in the error as an argument.
	//err := row.Scan(&user.Id, &user.Role, &user.Karyawan.KodeLokasi)
	//if err != nil {
	//	utils.PanicIfError(err)
	//}

	// Return the user variable and a nil error.
	return user, nil
}

// GetUserProfile retrieves a user's profile information from the database using the given IdUser.
// It returns a domain.User value and an error.
func (repository *UserRepositoryImpl) GetUserProfile(ctx context.Context, db *firestore.Client, IdUser uint) (domain.User, error) {
	var user domain.User

	//row := db.Table("user as a").Select("a.id, a.username, a.role, b.kode_lokasi, b.foto, b.email, b.no_telp, b.alamat, b.nama, b.nik").
	//	InnerJoins("inner join karyawan as b on a.id=b.id_user").
	//	Where("a.id = ?", strconv.FormatUint(uint64(IdUser), 10)).Row()
	//
	//if errors.Is(row.Err(), gorm.ErrRecordNotFound) {
	//	return user, errors.New("user not found")
	//}
	//
	//err := row.Scan(&user.Id, &user.Username, &user.Role, &user.Karyawan.KodeLokasi, &user.Karyawan.Foto, &user.Karyawan.Email,
	//	&user.Karyawan.NoTelp, &user.Karyawan.Alamat, &user.Karyawan.Nama, &user.Karyawan.Nik)
	//
	//if err != nil {
	//	utils.PanicIfError(err)
	//}

	return user, nil
}

func (repository *UserRepositoryImpl) CheckDuplicateUsername(ctx context.Context, db *firestore.Client, username string) error {
	//Create an empty user variable
	var user domain.User

	//Query the "users" collection for documents where the "username" field is equal to the provided "username" string
	iter := db.Collection("users").Where("username", "==", username).Documents(ctx)

	//Iterate through the returned documents
	for {
		document, err := iter.Next()
		//Break the loop if no more documents are returned
		if err == iterator.Done {
			break
		}
		//Populate the user variable with the data from the document
		_ = document.DataTo(&user)
	}

	//Check if the user variable is not zero value, indicating that the username does is exist
	if !reflect.ValueOf(user).IsZero() {
		return errors.New("username sudah terdaftar")
	}

	return nil
}
