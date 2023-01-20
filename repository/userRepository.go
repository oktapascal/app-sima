package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/oktapascal/app-barayya/models/domain"
)

type UserRepository interface {
	Store(ctx context.Context, db *firestore.Client, user domain.User)
	CheckUsername(ctx context.Context, db *firestore.Client, username string) (domain.User, error)
	CheckDuplicateUsername(ctx context.Context, db *firestore.Client, username string) error
	StoreSessionUser(ctx context.Context, db *firestore.Client, session domain.Session, user domain.User)
	DeleteSessionUser(ctx context.Context, db *firestore.Client, session domain.Session, user domain.User)
	GetUserProfile(ctx context.Context, db *firestore.Client, nik string) (domain.User, error)
}
