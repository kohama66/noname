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
		Date:         ent.Date,
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
