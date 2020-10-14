package usecase

import (
	"context"
	"errors"

	"github.com/myapp/noname/api/application/usecase/requestmodel"
	"github.com/myapp/noname/api/application/usecase/responsemodel"
	"github.com/myapp/noname/api/domain/entity"
	"github.com/myapp/noname/api/domain/repository"
	"github.com/myapp/noname/api/infrastructure/response"
)

// Reservation DIInterface
type Reservation interface {
	Create(ctx context.Context, req *requestmodel.ReservationCreate) (*responsemodel.ReservationCreate, error)
	FindByBeautician(ctx context.Context, req *requestmodel.ReservationFindByBeautician) (*responsemodel.ReservationFindByBeautician, error)
	Find(ctx context.Context, req *requestmodel.ReservationFind) (*responsemodel.ReservationFind, error)
}

type reservation struct {
	guestRepository       repository.Guest
	reservationRepository repository.Reservation
	reservationResponse   response.Reservation
	beauticianRepository  repository.Beautician
}

// NewReservation DI初期化
func NewReservation(
	guestRepository repository.Guest,
	reservationRepository repository.Reservation,
	reservationResponse response.Reservation,
	beauticianRepository repository.Beautician,
) Reservation {
	return &reservation{
		guestRepository:       guestRepository,
		reservationRepository: reservationRepository,
		reservationResponse:   reservationResponse,
		beauticianRepository:  beauticianRepository,
	}
}

func (r *reservation) Create(ctx context.Context, req *requestmodel.ReservationCreate) (*responsemodel.ReservationCreate, error) {
	gs, err := r.guestRepository.GetByAuthID(ctx, req.AuthID)
	if err != nil {
		return nil, err
	}
	bk, err := r.reservationRepository.ExistsSpaceDoubleBooking(ctx, req.Date, req.SpaceID)
	if err != nil {
		return nil, err
	}
	if bk {
		return nil, errors.New("このスペースの予約が重複しています")
	}
	bk, err = r.reservationRepository.ExistsBeauticianDoubleBooking(ctx, req.Date, req.BeauticianID)
	if err != nil {
		return nil, err
	}
	if bk {
		return nil, errors.New("美容師の予約が重複しています")
	}
	ent := req.NewReservation(gs.ID)
	if err = r.reservationRepository.Create(ctx, ent); err != nil {
		return nil, err
	}
	return r.reservationResponse.NewReservationCreate(ent), nil
}

func (r *reservation) FindByBeautician(ctx context.Context, req *requestmodel.ReservationFindByBeautician) (*responsemodel.ReservationFindByBeautician, error) {
	bt, err := r.beauticianRepository.GetByAuthID(ctx, req.AuthID)
	if err != nil {
		return nil, err
	}
	rv, err := r.reservationRepository.FindByBeautician(ctx, bt.ID)
	if err != nil {
		return nil, err
	}
	return r.reservationResponse.NewReservationFindByBeautician(rv), nil
}

func (r *reservation) Find(ctx context.Context, req *requestmodel.ReservationFind) (*responsemodel.ReservationFind, error) {
	var reva []*entity.Reservation
	if req.BeauticianRandID != "" {
		bt, err := r.beauticianRepository.GetByRandID(ctx, req.BeauticianRandID)
		if err != nil {
			return nil, err
		}
		rv, err := r.reservationRepository.FindByBeautician(ctx, bt.ID)
		if err != nil {
			return nil, err
		}
		reva = rv
	}
	return r.reservationResponse.NewReservationFind(reva), nil
}
