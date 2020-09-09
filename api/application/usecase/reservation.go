package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/myapp/noname/api/application/usecase/requestmodel"
	"github.com/myapp/noname/api/application/usecase/responsemodel"
	"github.com/myapp/noname/api/domain/repository"
	"github.com/myapp/noname/api/infrastructure/response"
	"github.com/volatiletech/null"
)

// Reservation DIInterface
type Reservation interface {
	Create(ctx context.Context, req *requestmodel.ReservationCreate) (*responsemodel.ReservationCreate, error)
}

type reservation struct {
	guestRepository       repository.Guest
	reservationRepository repository.Reservation
	reservationResponse   response.Reservation
}

// NewReservation DI初期化
func NewReservation(
	guestRepository repository.Guest,
	reservationRepository repository.Reservation,
	reservationResponse response.Reservation,
) Reservation {
	return &reservation{
		guestRepository:       guestRepository,
		reservationRepository: reservationRepository,
		reservationResponse:   reservationResponse,
	}
}

func (r *reservation) Create(ctx context.Context, req *requestmodel.ReservationCreate) (*responsemodel.ReservationCreate, error) {
	gs, err := r.guestRepository.GetByAuthID(ctx, req.AuthID)
	if err != nil {
		return nil, err
	}
	d, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, err
	}
	dt := null.TimeFrom(d)
	tm := null.StringFrom(req.Time)
	bk, err := r.reservationRepository.ExistsSpaceDoubleBooking(ctx, dt, tm, req.SpaceID)
	if err != nil {
		return nil, err
	}
	if bk {
		return nil, errors.New("このスペースの予約が重複しています")
	}
	bk, err = r.reservationRepository.ExistsBeauticianDoubleBooking(ctx, dt, tm, req.BeauticianID)
	if err != nil {
		return nil, err
	}
	if bk {
		return nil, errors.New("美容師の予約が重複しています")
	}
	ent := req.NewReservation(dt, tm, gs.ID)
	if err = r.reservationRepository.Create(ctx, ent); err != nil {
		return nil, err
	}
	return r.reservationResponse.NewReservationCreate(ent), nil
}
