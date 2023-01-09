package services

import (
	"context"
	"github.com/oktapascal/app-barayya/models/web"
)

type UserServices interface {
	Save(ctx context.Context, request web.RegisterRequest)
	CheckUsername(ctx context.Context, username string) web.UserResponses
	SaveSessionUser(ctx context.Context, request web.SessionRequest)
	DestroySessionUser(ctx context.Context, authToken string)
	GetSessionUser(ctx context.Context, authToken string) web.UserResponses
	GetUserProfile(ctx context.Context, IdUser uint) web.UserProfileResponses
}
