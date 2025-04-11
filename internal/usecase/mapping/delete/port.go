package delete

import (
	"context"
)

type Repository interface {
	Delete(ctx context.Context, code string) error
}
