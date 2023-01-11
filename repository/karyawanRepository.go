package repository

import (
	"context"
	"github.com/oktapascal/app-barayya/models/domain"
	"gorm.io/gorm"
)

type KaryawanRepository interface {
	Update(ctx context.Context, db *gorm.DB, karyawan domain.Karyawan)
}
