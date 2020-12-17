package repository

import (
	"context"
	"time"

	"github.com/myapp/noname/api/domain/entity"
)

// Beautician DIInterface
type Beautician interface {
	Create(ctx context.Context, ent *entity.Beautician) error
	GetByUserID(ctx context.Context, ID int64) (*entity.Beautician, error)
	// GetByRandID(ctx context.Context, randID string) (*entity.Beautician, error)
	Find(ctx context.Context, date time.Time, salon *int64, menus []int64) ([]*entity.User, error)
	// GetAll(ctx context.Context) ([]*entity.Beautician, error)
	FindPossibleSalon(ctx context.Context, beauciaonID int64) (entity.BeauticianSalonSlice, error)
}
