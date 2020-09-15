package repository

import (
	"context"

	"github.com/myapp/noname/api/domain/entity"
)

// Beautician DIInterface
type Beautician interface {
	Create(ctx context.Context, ent *entity.Beautician) error
	GetByAuthID(ctx context.Context, authID string) (*entity.Beautician, error)
	GetAll(ctx context.Context) (entity.BeauticianSlice, error)
}
