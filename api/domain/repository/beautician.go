package repository

import (
	"context"
	"time"

	"github.com/myapp/noname/api/domain/entity"
)

// Beautician DIInterface
type Beautician interface {
	Create(ctx context.Context, ent *entity.Beautician) error
	GetByAuthID(ctx context.Context, authID string) (*entity.Beautician, error)
	Find(ctx context.Context, date time.Time, menu, salon *int64) ([]*entity.Beautician, error)
	GetAll(ctx context.Context) ([]*entity.Beautician, error)
}
