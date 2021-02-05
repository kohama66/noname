package repository

import (
	"context"

	"github.com/myapp/noname/api/domain/entity"
)

// User DIInterface
type User interface {
	GetByAuthID(ctx context.Context, authID string) (*entity.User, error)
	Create(ctx context.Context, ent *entity.User) error
	GetByRandID(ctx context.Context, randID string) (*entity.User, error)
}
