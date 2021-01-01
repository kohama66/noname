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
	"github.com/rs/xid"
)

// Reservation DIInterface
type Reservation interface {
	Create(ctx context.Context, req *requestmodel.ReservationCreate) (*responsemodel.ReservationCreate, error)
	FindByBeautician(ctx context.Context, req *requestmodel.ReservationFindByBeautician) (*responsemodel.ReservationFindByBeautician, error)
	Find(ctx context.Context, req *requestmodel.ReservationFind) (*responsemodel.ReservationFind, error)
	FindByUser(ctx context.Context, req *requestmodel.ReservationFindByUser) (*responsemodel.ReservationFindByUser, error)
	GetInfo(ctx context.Context, re *requestmodel.ReservationGetInfo) (*responsemodel.ReservationGetInfo, error)
	SetHoliday(ctx context.Context, req *requestmodel.ReservationSetHoliday) (*responsemodel.ReservationSetHoliday, error)
}

type reservation struct {
	userRepository        repository.User
	reservationRepository repository.Reservation
	reservationResponse   response.Reservation
	beauticianRepository  repository.Beautician
	salonRepository       repository.Salon
	menuRepository        repository.Menu
}

// NewReservation DI初期化
func NewReservation(
	userRepository repository.User,
	reservationRepository repository.Reservation,
	reservationResponse response.Reservation,
	beauticianRepository repository.Beautician,
	salonRepository repository.Salon,
	menuRepository repository.Menu,
) Reservation {
	return &reservation{
		userRepository:        userRepository,
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
	gs, err := r.userRepository.GetByAuthID(ctx, req.AuthID)
	if err != nil {
		return nil, err
	}
	bt, err := r.userRepository.GetByRandID(ctx, req.BeauticianID)
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
	ent := req.NewReservation(gs.ID, sp.ID, bt.ID, xid.New().String(), convDate)
	if err = r.reservationRepository.Create(ctx, ent, mnids); err != nil {
		return nil, err
	}
	return r.reservationResponse.NewReservationCreate(ent), nil
}

func (r *reservation) FindByBeautician(ctx context.Context, req *requestmodel.ReservationFindByBeautician) (*responsemodel.ReservationFindByBeautician, error) {
	bt, err := r.userRepository.GetByAuthID(ctx, req.AuthID)
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
		bt, err := r.userRepository.GetByRandID(ctx, req.BeauticianRandID)
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

func (r *reservation) FindByUser(ctx context.Context, req *requestmodel.ReservationFindByUser) (*responsemodel.ReservationFindByUser, error) {
	gs, err := r.userRepository.GetByAuthID(ctx, req.AuthID)
	if err != nil {
		return nil, err
	}
	rv, err := r.reservationRepository.FindByUser(ctx, gs.ID)
	if err != nil {
		return nil, err
	}
	return r.reservationResponse.NewReservationFindByUser(rv), nil
}

func (r *reservation) GetInfo(ctx context.Context, re *requestmodel.ReservationGetInfo) (*responsemodel.ReservationGetInfo, error) {
	rs, err := r.reservationRepository.GetByRandID(ctx, re.RandID)
	if err != nil {
		return nil, err
	}
	sl, err := r.salonRepository.GetBySpaceID(ctx, rs.SpaceID)
	if err != nil {
		return nil, err
	}
	mn, err := r.menuRepository.GetBeauticianMenusByReservationID(ctx, rs.ID)
	if err != nil {
		return nil, err
	}
	fmt.Println(mn)
	return r.reservationResponse.NewReservationInfo(rs, sl, mn), nil
}

func (r *reservation) SetHoliday(ctx context.Context, req *requestmodel.ReservationSetHoliday) (*responsemodel.ReservationSetHoliday, error) {
	bt, err := r.userRepository.GetByAuthID(ctx, req.AuthID)
	if err != nil {
		return nil, err
	}
	if !bt.IsBeautician {
		return nil, fmt.Errorf("You are not beautician")
	}
	ex, err := r.reservationRepository.ExistsByDate(ctx, req.Holiday)
	if err != nil {
		return nil, err
	}
	var holiday *entity.Reservation
	if ex {
		rs, err := r.reservationRepository.FindByDate(ctx, req.Holiday)
		if err != nil {
			return nil, err
		}
		if rs.Holiday { // 既に休日に場合
			if _, err := r.reservationRepository.Delete(ctx, rs); err != nil {
				return nil, err
			}
			holiday = rs
		} else { // 既にお客様の予約が入っている場合
			return nil, fmt.Errorf("Reservation already exists")
		}
	} else { // 休日に設定
		hd := req.NewReservationSetHoliday(xid.New().String(), bt.ID)
		if err := r.reservationRepository.CreateHoliday(ctx, hd); err != nil {
			return nil, err
		}
		holiday = hd
	}
	return r.reservationResponse.NewSetHoliday(holiday), nil
}
