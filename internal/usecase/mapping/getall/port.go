package getall

import (
	"context"

	"github.com/enesanbar/go-service/router"
)

type Repository interface {
	FindAll(ctx context.Context, request *Request) (*router.PagedResponse, error)
}
