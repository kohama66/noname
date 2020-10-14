package repository

import (
	"context"

	"github.com/myapp/noname/api/domain/entity"
)

// Salon DIInterface
type Salon interface {
	GetByRandID(ctx context.Context, randID string) (*entity.Salon, error)
	Find(ctx context.Context, beauticianID int64) (entity.SalonSlice, error)
}
