package repository

import (
	"context"

	"github.com/myapp/noname/api/domain/entity"
)

type Menu interface {
	GetByRandID(ctx context.Context, randID string) (*entity.Menu, error)
}
