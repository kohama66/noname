package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/myapp/noname/api/application/usecase/requestmodel"
	"github.com/myapp/noname/api/domain/repository"
	"github.com/volatiletech/null"
)

// Reservation DIInterface
type Reservation interface {
	Create(ctx context.Context, req *requestmodel.ReservationCreate) (error, error)
}

type reservation struct {
	guestRepository       repository.Guest
	reservationRepository repository.Reservation
}

// NewReservation DI初期化
func NewReservation(
	guestRepository repository.Guest,
	reservationRepository repository.Reservation,
) Reservation {
	return &reservation{
		guestRepository:       guestRepository,
		reservationRepository: reservationRepository,
	}
}

func (r *reservation) Create(ctx context.Context, req *requestmodel.ReservationCreate) (error, error) {
	// me, err := r.guestRepository.GetByAuthID(ctx, req.AuthID)
	// if err != nil {
	// 	return
	// }
	dt, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, err
	}
	bk, err := r.reservationRepository.ExistsSpaceDoubleBooking(ctx, null.TimeFrom(dt), null.StringFrom(req.Time), req.SpaceID)
	if err != nil {
		return nil, err
	}
	if bk {
		return nil, errors.New("このスペースの予約が重複しています")
	}
	bk, err = r.reservationRepository.ExistsBeauticianDoubleBooking(ctx, null.TimeFrom(dt), null.StringFrom(req.Time), req.BeauticianID)
	if err != nil {
		return nil, err
	}
	if bk {
		return nil, errors.New("美容師の予約が重複しています")
	}
	return nil, nil
}
