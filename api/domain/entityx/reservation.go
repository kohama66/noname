package entityx

import (
	"time"

	"github.com/myapp/noname/api/domain/entity"
)

func NewReservationMenus(beaiticianID int64, menuIDs []int64) []*entity.ReservationMenu {
	rms := make([]*entity.ReservationMenu, len(menuIDs))
	for i, menuID := range menuIDs {
		rm := &entity.ReservationMenu{
			BeauticianMenuID: menuID,
		}
		rms[i] = rm
	}
	return rms
}

func NewReservationMenu(reservationID int64, menuID int64) *entity.ReservationMenu {
	return &entity.ReservationMenu{
		ReservationID:    reservationID,
		BeauticianMenuID: menuID,
	}
}

// ReservationGetByUser 用構造体
type ReservationGetByUser struct {
	ID                  int64                    `boil:"reservation_id" json:"id"`
	Date                time.Time                `boil:"date" json:"date"`
	UserID              int64                    `boil:"user_id" json:"user_id"`
	SalonName           string                   `boil:"salon_name" json:"salon_name"`
	BeauticianFirstName string                   `boil:"first_name" json:"beautician_first_name"`
	BeauticianLatsName  string                   `boil:"last_name" json:"beautician_last_name"`
	Menus               []*entity.BeauticianMenu `json:"menus"`
	CreatedAt           time.Time                `boil:"created_at" json:"created_at"`
	UpdatedAt           time.Time                `boil:"updated_at" json:"updated_at"`
}

// ReservationGetByUserSlice 用構造体
type ReservationGetByUserSlice struct {
	Reservations []*ReservationGetByUser
}

// ReservationInfo 予約詳細
type ReservationInfo struct {
	ID         int64              `boil:"id" json:"id"`
	RandID     string             `boil:"randId" json:"randId"`
	Date       time.Time          `boil:"date" json:"date"`
	Holiday    bool               `boil:"holiday" json:"holiday"`
	Salon      *entity.Salon      `boil:"salon" json:"salon"`
	Beautician *entity.Beautician `boil:"beautician" json:"beautician"`
	User       *entity.User       `boil:"user" json:"user"`
}
