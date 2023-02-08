package services

import (
	"context"
	"github.com/oktapascal/app-sima/models/web"
)

type MenuServices interface {
	GetAll(ctx context.Context) []web.MenuResponses
}
