package repository

import (
	"context"
	"github.com/oktapascal/app-barayya/models/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Store(ctx context.Context, db *gorm.DB, user domain.User)
	CheckUsername(ctx context.Context, db *gorm.DB, username string) (domain.User, error)
	StoreSessionUser(ctx context.Context, db *gorm.DB, session domain.Session)
	DeleteSessionUser(ctx context.Context, db *gorm.DB, authToken string)
	GetUserBySession(ctx context.Context, db *gorm.DB, authToken string) (domain.User, error)
}
