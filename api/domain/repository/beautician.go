package repository

import (
	"context"

	"github.com/myapp/noname/api/domain/entity"
)

// Beautician DIInterface
type Beautician interface {
	Create(ctx context.Context, ent *entity.Beautician) error
	Get(ctx context.Context, id int64) (*entity.Beautician, error)
}
