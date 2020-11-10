package entityx

import "github.com/myapp/noname/api/domain/entity"

// Reservation DIInterface
type Reservation interface {
	NewReservationMenus(beaiticianID int64, menuIDs []int64) []*entity.ReservationMenu
	NewReservationMenu(reservationID int64, menuID int64) *entity.ReservationMenu
}

type reservation struct{}

// NewReservation DI初期化関数
func NewReservation() Reservation {
	return &reservation{}
}

func (r *reservation) NewReservationMenus(beaiticianID int64, menuIDs []int64) []*entity.ReservationMenu {
	rms := make([]*entity.ReservationMenu, len(menuIDs))
	for i, menuID := range menuIDs {
		rm := &entity.ReservationMenu{
			BeauticianMenuID: menuID,
		}
		rms[i] = rm
	}
	return rms
}

func (r *reservation) NewReservationMenu(reservationID int64, menuID int64) *entity.ReservationMenu {
	return &entity.ReservationMenu{
		ReservationID:    reservationID,
		BeauticianMenuID: menuID,
	}
}
