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
		Jabatan:    user.Jabatan,
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
	// Store the session data in the sessions sub collection of the user's document in the "users" collection
	_, _, err := db.Collection("users").Doc(user.Nik).Collection("sessions").Add(ctx, domain.Session{
		AuthToken: session.AuthToken,
		CreatedAt: session.CreatedAt,
		DeletedAt: session.DeletedAt,
	})

	// Panic if there is an error
	utils.PanicIfError(err)
}

func (repository *UserRepositoryImpl) DeleteSessionUser(ctx context.Context, db *firestore.Client, session domain.Session, user domain.User) {
	// Initialize a variable to hold the session ID
	var id string

	// Get the iterator for the matching documents in the "sessions" sub collection where the auth token matches the one provided
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
	_, err := db.Collection("users").Doc(user.Nik).Collection("sessions").Doc(id).Update(ctx, []firestore.Update{
		{
			Path:  "deleted_at",
			Value: session.DeletedAt,
		},
	})
	utils.PanicIfError(err)
}

func (repository *UserRepositoryImpl) GetUserProfile(ctx context.Context, db *firestore.Client, nik string) (domain.User, error) {
	// Initialize a variable to hold the user data
	var user domain.User

	// Get the user document from the "users" collection using the provided NIK
	document, err := db.Collection("users").Doc(nik).Get(ctx)
	utils.PanicIfError(err)

	// Check if the document exists
	if !document.Exists() {
		// Return an error if the document does not exist
		return user, errors.New("user tidak ditemukan")
	}

	// Convert the document data to the user struct
	err = document.DataTo(&user)
	utils.PanicIfError(err)

	// Return the user struct and a nil error
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

func (repository *UserRepositoryImpl) Update(ctx context.Context, db *firestore.Client, user domain.User) {
	// Update the user's data in the "users" collection using the provided NIK
	_, err := db.Collection("users").Doc(user.Nik).Update(ctx, []firestore.Update{
		{
			Path:  "nama",
			Value: user.Nama,
		},
		{
			Path:  "no_telp",
			Value: user.NoTelp,
		},
		{
			Path:  "email",
			Value: user.Email,
		},
		{
			Path:  "alamat",
			Value: user.Alamat,
		},
		{
			Path:  "jabatan",
			Value: user.Jabatan,
		},
	})
	// Panic if there is an error
	utils.PanicIfError(err)
}
