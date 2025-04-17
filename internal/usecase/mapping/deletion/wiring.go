package deletion

import (
	"context"
)

type Service interface {
	Execute(ctx context.Context, code string) error
}
