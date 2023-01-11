package repository

import (
	"context"
	"github.com/oktapascal/app-barayya/models/domain"
	"gorm.io/gorm"
	"strconv"
)

// KaryawanRepositoryImpl struct represents a repository for data related to Karyawan (employees)
type KaryawanRepositoryImpl struct {
}

// NewKaryawanRepositoryImpl returns a new instance of KaryawanRepositoryImpl struct
// that can be used to interact with the underlying data storage and perform CRUD operations on Karyawan data
func NewKaryawanRepositoryImpl() *KaryawanRepositoryImpl {
	// create and return a pointer to a new instance of KaryawanRepositoryImpl
	return &KaryawanRepositoryImpl{}
}

func (repository *KaryawanRepositoryImpl) Update(ctx context.Context, db *gorm.DB, karyawan domain.Karyawan) {
	//  update the karyawan record in the database using the gorm.DB instance
	//  and update it with the values from the karyawan struct passed as argument
	db.Table("karyawan").Where("id_user = ?", strconv.FormatUint(uint64(karyawan.IdUser), 10)).
		UpdateColumns(domain.Karyawan{
			Nik:    karyawan.Nik,
			Nama:   karyawan.Nama,
			Alamat: karyawan.Alamat,
			NoTelp: karyawan.NoTelp,
			Email:  karyawan.Email,
		})
}
