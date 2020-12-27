package repository

import (
	"context"
	"time"

	"github.com/myapp/noname/api/domain/entity"
	"github.com/myapp/noname/api/domain/entityx"
)

// Reservation DIInterface
type Reservation interface {
	ExistsSpaceDoubleBooking(ctx context.Context, date time.Time, spaceID int64) (bool, error)
	ExistsBeauticianDoubleBooking(ctx context.Context, date time.Time, beauticianID int64) (bool, error)
	FindByBeautician(ctx context.Context, beauticianID int64) (entity.ReservationSlice, error)
	Create(ctx context.Context, re *entity.Reservation, menuIDs []int64) error
	FindByUser(ctx context.Context, userID int64) ([]*entityx.ReservationGetByUser, error)
	GetByRandID(ctx context.Context, randID string) (*entity.Reservation, error)
}
