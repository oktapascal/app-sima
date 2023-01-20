package services

import (
	"context"
	"github.com/oktapascal/app-barayya/models/web"
)

type UserServices interface {
	Register(ctx context.Context, request web.RegisterRequest)
	Login(ctx context.Context, request web.LoginRequest) web.UserResponses
	Logout(ctx context.Context, request web.SessionRequest)
	StoreSessionUser(ctx context.Context, request web.SessionRequest)
	GetSessionUser(ctx context.Context, authToken string) web.UserResponses
}
