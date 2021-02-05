package response

import (
	"github.com/myapp/noname/api/application/usecase/responsemodel"
	"github.com/myapp/noname/api/domain/entity"
	"github.com/myapp/noname/api/domain/entityx"
)

// Reservation DIInterface
type Reservation interface {
	NewReservationCreate(ent *entity.Reservation) *responsemodel.ReservationCreate
	NewReservationFindByBeautician(ents []*entity.Reservation) *responsemodel.ReservationFindByBeautician
	NewReservationFind(ents []*entity.Reservation) *responsemodel.ReservationFind
	NewReservationFindByUser(ents []*entityx.ReservationGetByUser) *responsemodel.ReservationFindByUser
	NewReservationInfo(rs *entity.Reservation, sl *entity.Salon, mn []*entity.BeauticianMenu) *responsemodel.ReservationGetInfo
	NewSetHoliday(ent *entity.Reservation) *responsemodel.ReservationSetHoliday
}

type reservation struct {
}

// NewReservation DI初期化関数
func NewReservation() Reservation {
	return &reservation{}
}

// NewResponseModelReservation エンティティーをレスポンスへ変換
func NewResponseModelReservation(ent *entity.Reservation) *responsemodel.Reservation {
	return &responsemodel.Reservation{
		RandID:       ent.RandID,
		Date:         ent.Date,
		Holiday:      ent.Holiday,
		SpaceID:      ent.SpaceID,
		BeauticianID: ent.BeauticianID,
		UserID:       ent.UserID,
		CreatedAt:    ent.CreatedAt,
		UpdatedAt:    ent.UpdatedAt,
	}
}

func (r *reservation) NewReservationCreate(ent *entity.Reservation) *responsemodel.ReservationCreate {
	return &responsemodel.ReservationCreate{
		Reservation: NewResponseModelReservation(ent),
	}
}

func (r *reservation) NewReservationFindByBeautician(ents []*entity.Reservation) *responsemodel.ReservationFindByBeautician {
	rv := make([]*responsemodel.Reservation, len(ents))
	for i, v := range ents {
		rv[i] = NewResponseModelReservation(v)
	}
	return &responsemodel.ReservationFindByBeautician{
		Reservations: rv,
	}
}

func (r *reservation) NewReservationFind(ents []*entity.Reservation) *responsemodel.ReservationFind {
	rv := make([]*responsemodel.Reservation, len(ents))
	for i, v := range ents {
		rv[i] = NewResponseModelReservation(v)
	}
	return &responsemodel.ReservationFind{
		Reservations: rv,
	}
}

func (r *reservation) NewReservationFindByUser(ents []*entityx.ReservationGetByUser) *responsemodel.ReservationFindByUser {
	rv := make([]*responsemodel.ReservationGetByUser, len(ents))
	for i, v := range ents {
		rv[i] = &responsemodel.ReservationGetByUser{
			ID:                  v.ID,
			Date:                v.Date,
			UserID:              v.UserID,
			SalonName:           v.SalonName,
			BeauticianFirstName: v.BeauticianFirstName,
			BeauticianLatsName:  v.BeauticianLatsName,
			Menus:               NewBeauticianMenusResponsemodel(v.Menus),
		}
	}
	return &responsemodel.ReservationFindByUser{
		Reservations: rv,
	}
}

// NewReservationInfo 予約詳細レスポンスモデル変換
func NewReservationInfo(rs *entity.Reservation, sl *entity.Salon, mn []*entity.BeauticianMenu) *responsemodel.ReservationInfo {
	return &responsemodel.ReservationInfo{
		RandID:    rs.RandID,
		Date:      rs.Date,
		User:      NewUserResponsemodel(rs.R.User),
		Salon:     NewSalonResponseModel(sl),
		Menus:     NewBeauticianMenusResponsemodel(mn),
		CreatedAt: rs.CreatedAt,
		UpdatedAt: rs.UpdatedAt,
	}
}

func (r *reservation) NewReservationInfo(rs *entity.Reservation, sl *entity.Salon, mn []*entity.BeauticianMenu) *responsemodel.ReservationGetInfo {
	return &responsemodel.ReservationGetInfo{
		ReservationInfo: NewReservationInfo(rs, sl, mn),
	}
}

func (r *reservation) NewSetHoliday(ent *entity.Reservation) *responsemodel.ReservationSetHoliday {
	return &responsemodel.ReservationSetHoliday{
		Reservation: NewResponseModelReservation(ent),
	}
}
