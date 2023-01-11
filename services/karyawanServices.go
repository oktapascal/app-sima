package services

import (
	"context"
	"github.com/oktapascal/app-barayya/models/web"
)

type KaryawanServices interface {
	Update(ctx context.Context, request web.KaryawanUpdateRequest)
}
