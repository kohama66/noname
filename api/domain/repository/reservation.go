package repository

import (
	"context"
	"time"

	"github.com/myapp/noname/api/domain/entity"
	"github.com/volatiletech/null"
)

type Reservation interface {
	ExistsSpaceDoubleBooking(ctx context.Context, date null.Time, time null.String, spaceID int64) (bool, error)
	ExistsBeauticianDoubleBooking(ctx context.Context, date null.Time, time null.String, beauticianID int64) (bool, error)
	Create(ctx context.Context, ent *entity.Reservation) error
	FindByBeautician(ctx context.Context, beauticianID int64, date time.Time) (entity.ReservationSlice, error)
}
