package requestmodel

import (
	"github.com/myapp/noname/api/domain/entity"
	"github.com/volatiletech/null"
)

// ReservationCreate 予約構造体
type ReservationCreate struct {
	AuthID       string `json:"-"`
	BeauticianID int64  `json:"beauticiaId"`
	SpaceID      int64  `json:"spaceId"`
	MenuID       int64  `json:"menuId"`
	Date         string `json:"date"`
	Time         string `json:"time"`
}

// NewReservation 予約モデル変換メソッド
func (r *ReservationCreate) NewReservation(date null.Time, time null.String, guestID int64) *entity.Reservation {
	return &entity.Reservation{
		Date:         date,
		Time:         time,
		SpaceID:      r.SpaceID,
		BeauticianID: r.BeauticianID,
		GuestID:      guestID,
		MenuID:       r.MenuID,
	}
}
