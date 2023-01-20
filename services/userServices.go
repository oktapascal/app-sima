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
	GetUserProfile(ctx context.Context, nik string) web.UserProfileResponses
	UpdateUserProfile(ctx context.Context, request web.UpdateUserProfileRequest)
}
