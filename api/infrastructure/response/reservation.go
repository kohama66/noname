package response

import (
	"github.com/myapp/noname/api/application/usecase/responsemodel"
	"github.com/myapp/noname/api/domain/entity"
)

// Reservation DIInterface
type Reservation interface {
	NewReservationCreate(ent *entity.Reservation) *responsemodel.ReservationCreate
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
		Time:         ent.Time,
		SpaceID:      ent.SpaceID,
		BeauticianID: ent.BeauticianID,
		GuestID:      ent.GuestID,
		MenuID:       ent.MenuID,
		CreatedAt:    ent.CreatedAt,
		UpdatedAt:    ent.UpdatedAt,
	}
}

func (r *reservation) NewReservationCreate(ent *entity.Reservation) *responsemodel.ReservationCreate {
	return &responsemodel.ReservationCreate{
		Reservation: NewResponseModelReservation(ent),
	}
}
