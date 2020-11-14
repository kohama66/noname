package response

import (
	"github.com/myapp/noname/api/application/usecase/responsemodel"
	"github.com/myapp/noname/api/domain/entity"
	"github.com/myapp/noname/api/domain/entityx"
)

// Guest DIInterface
type Guest interface {
	NewGuestGet(ent *entity.Guest, rent []*entityx.ReservationGetByGuest) *responsemodel.GuestGet
}

type guest struct{}

// NewGuest DI初期化関数
func NewGuest() Guest {
	return &guest{}
}

// NewGuestResponsemodel レスポンスモデル変換関数
func NewGuestResponsemodel(ent *entity.Guest) *responsemodel.Guest {
	return &responsemodel.Guest{
		RandID:    ent.RandID,
		FirstName: ent.FirstName,
		LastName:  ent.LastName,
		CreatedAt: ent.CreatedAt,
		UpdatedAt: ent.CreatedAt,
	}
}

func (g *guest) NewGuestGet(ent *entity.Guest, rent []*entityx.ReservationGetByGuest) *responsemodel.GuestGet {
	rs := make([]*responsemodel.ReservationGetByGuest, len(rent))
	for i, v := range rent {
		rs[i] = &responsemodel.ReservationGetByGuest{
			ID:                  v.ID,
			Date:                v.Date,
			GuestID:             v.GuestID,
			SalonName:           v.SalonName,
			BeauticianFirstName: v.BeauticianFirstName,
			BeauticianLatsName:  v.BeauticianLatsName,
		}
	}
	return &responsemodel.GuestGet{
		RandID:       ent.RandID,
		FirstName:    ent.FirstName,
		LastName:     ent.LastName,
		Reservations: rs,
		CreatedAt:    ent.CreatedAt,
		UpdatedAt:    ent.CreatedAt,
	}
}
