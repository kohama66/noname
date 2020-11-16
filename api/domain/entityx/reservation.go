package entityx

import (
	"time"

	"github.com/myapp/noname/api/domain/entity"
)

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

// ReservationGetByGuest 用構造体
type ReservationGetByGuest struct {
	ID                  int64                    `boil:"reservation_id" json:"id"`
	Date                time.Time                `boil:"date" json:"date"`
	GuestID             int64                    `boil:"guest_id" json:"guest_id"`
	SalonName           string                   `boil:"salon_name" json:"salon_name"`
	BeauticianFirstName string                   `boil:"first_name" json:"beautician_first_name"`
	BeauticianLatsName  string                   `boil:"last_name" json:"beautician_last_name"`
	Menus               []*entity.BeauticianMenu `json:"menus"`
	CreatedAt           time.Time                `boil:"created_at" json:"created_at"`
	UpdatedAt           time.Time                `boil:"updated_at" json:"updated_at"`
}

// ReservationGetByGuest 用構造体
type ReservationGetByGuestSlice struct {
	Reservations []*ReservationGetByGuest
}
