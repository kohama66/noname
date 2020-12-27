package repository

import (
	"context"

	"github.com/myapp/noname/api/domain/entity"
)

// Menu DIInterface
type Menu interface {
	GetByRandID(ctx context.Context, randID string) (*entity.Menu, error)
	GetAll(ctx context.Context) (entity.MenuSlice, error)
	FindByBeautician(ctx context.Context, beauticianID int64) (entity.MenuSlice, error)
	FindByRandID(ctx context.Context, randIDs []string) (entity.MenuSlice, error)
	FindByBeauticianWithMenuRandIDs(ctx context.Context, beauticianID int64, menuIDs []string) (entity.BeauticianMenuSlice, error)
	ExistsByBeauticianIDWithMenuIDs(ctx context.Context, beauticianID int64, menuIDs []int64) (bool, error)
	GetBeauticianMenusByReservationID(ctx context.Context, reservationID int64) (entity.BeauticianMenuSlice, error)
}
