package repository

import (
	"context"

	"github.com/volatiletech/null"
)

type Reservation interface {
	ExistsSpaceDoubleBooking(ctx context.Context, date null.Time, time null.String, spaceID int64) (bool, error)
	ExistsBeauticianDoubleBooking(ctx context.Context, date null.Time, time null.String, beauticianID int64) (bool, error)
}
