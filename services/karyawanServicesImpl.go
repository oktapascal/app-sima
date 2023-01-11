package services

import (
	"context"
	"github.com/oktapascal/app-barayya/models/domain"
	"github.com/oktapascal/app-barayya/models/web"
	"github.com/oktapascal/app-barayya/repository"
	"github.com/oktapascal/app-barayya/utils"
	"gorm.io/gorm"
)

// KaryawanServicesImpl struct is used to define services related to Karyawan
type KaryawanServicesImpl struct {
	// KaryawanRepository is an interface or struct that contains methods to interact with the underlying data storage
	KaryawanRepository repository.KaryawanRepository
	// Db is a pointer to a gorm instance, it is used to interact with the database
	Db *gorm.DB
}

// NewKaryawanServicesImpl returns a new instance of KaryawanServicesImpl struct that can be used to interact with Karyawan data
// it takes two arguments as input :
// - karyawanRepository: an instance of the karyawanRepository struct to pass to the service
// - db : a pointer to a gorm instance, represents a connection to the database
func NewKaryawanServicesImpl(karyawanRepository repository.KaryawanRepository, db *gorm.DB) *KaryawanServicesImpl {
	// create and return a pointer to a new instance of KaryawanServicesImpl
	return &KaryawanServicesImpl{KaryawanRepository: karyawanRepository, Db: db}
}

func (services *KaryawanServicesImpl) Update(ctx context.Context, request web.KaryawanUpdateRequest) {
	// starts a new database transaction
	tx := services.Db.WithContext(ctx).Begin()
	//defer the commit or rollback of the transaction
	defer utils.CommitRollback(tx)

	//create a new karyawan struct and set the value of its fields
	karyawan := domain.Karyawan{
		IdUser: request.IdUser,
		Nik:    &request.Nik,
		Nama:   &request.Nama,
		Alamat: &request.Alamat,
		NoTelp: &request.NoTelp,
		Email:  &request.Email,
	}

	// call the update method on the karyawan repository with the transaction and the karyawan struct
	services.KaryawanRepository.Update(ctx, tx, karyawan)
}
