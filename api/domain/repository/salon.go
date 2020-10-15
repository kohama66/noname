package repository

import (
	"context"

	"github.com/myapp/noname/api/domain/entity"
)

// Salon DIInterface
type Salon interface {
	GetByRandID(ctx context.Context, randID string) (*entity.Salon, error)
	GetAll(ctx context.Context) (entity.SalonSlice, error)
	FindByBeautician(ctx context.Context, beauticianID int64) (entity.SalonSlice, error)
}
