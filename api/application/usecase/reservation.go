package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

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
	salonRepository       repository.Salon
	menuRepository        repository.Menu
}

// NewReservation DI初期化
func NewReservation(
	guestRepository repository.Guest,
	reservationRepository repository.Reservation,
	reservationResponse response.Reservation,
	beauticianRepository repository.Beautician,
	salonRepository repository.Salon,
	menuRepository repository.Menu,
) Reservation {
	return &reservation{
		guestRepository:       guestRepository,
		reservationRepository: reservationRepository,
		reservationResponse:   reservationResponse,
		beauticianRepository:  beauticianRepository,
		salonRepository:       salonRepository,
		menuRepository:        menuRepository,
	}
}

func (r *reservation) Create(ctx context.Context, req *requestmodel.ReservationCreate) (*responsemodel.ReservationCreate, error) {
	convDate, err := time.Parse("2006-01-02 15:04:05", req.Date)
	if err != nil {
		return nil, err
	}
	fmt.Println(req.AuthID)
	gs, err := r.guestRepository.GetByAuthID(ctx, req.AuthID)
	if err != nil {
		return nil, err
	}
	bt, err := r.beauticianRepository.GetByRandID(ctx, req.BeauticianID)
	if err != nil {
		return nil, err
	}
	mns, err := r.menuRepository.FindByRandID(ctx, req.MenuIDs)
	if err != nil {
		return nil, err
	}
	mnids := make([]int64, len(mns))
	for i, v := range mns {
		mnids[i] = v.ID
	}
	ok, err := r.menuRepository.ExistsByBeauticianIDWithMenuIDs(ctx, bt.ID, mnids)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("menu error")
	}
	sl, err := r.salonRepository.GetByRandID(ctx, req.SalonID)
	if err != nil {
		return nil, err
	}
	ok, err = r.salonRepository.ExistsByBeauticianWithSalon(ctx, bt.ID, sl.ID)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("salon error")
	}
	sp, err := r.salonRepository.GetVacantSpace(ctx, convDate, sl.ID)
	if err != nil {
		return nil, err
	}
	ng, err := r.reservationRepository.ExistsBeauticianDoubleBooking(ctx, convDate, bt.ID)
	if err != nil {
		return nil, err
	}
	if ng {
		return nil, errors.New("美容師の予約が重複しています")
	}
	ent := req.NewReservation(gs.ID, sp.ID, bt.ID, convDate)
	if err = r.reservationRepository.Create(ctx, ent, mnids); err != nil {
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
