package repository

import (
	"context"

	"github.com/myapp/noname/api/domain/entity"
)

type Salon interface {
	GetByRandID(ctx context.Context, randID string) (*entity.Salon, error)
}
