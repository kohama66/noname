package requestmodel

import (
	"time"

	"github.com/myapp/noname/api/domain/entity"
)

// ReservationCreate request構造体
type ReservationCreate struct {
	AuthID       string    `json:"-"`
	BeauticianID int64     `json:"beauticiaRandId"`
	SalonID      int64     `json:"salonRandId"`
	MenuIDs      []int64   `json:"menuIds"`
	Date         time.Time `json:"date"`
	// SpaceID      int64     `json:"spaceId"`
}

// NewReservation 予約モデル変換メソッド
func (r *ReservationCreate) NewReservation(guestID int64) *entity.Reservation {
	return &entity.Reservation{
		Date: r.Date,
		// SpaceID:      r.SpaceID,
		BeauticianID: r.BeauticianID,
		GuestID:      guestID,
	}
}

// ReservationFindByBeautician 美容師予約情報取得 構造体
type ReservationFindByBeautician struct {
	AuthID string `json:"-"`
	Offset int64  `schema:"offset"`
}

// ReservationFind 予約検索
type ReservationFind struct {
	BeauticianRandID string `schema:"beauticianRandId"`
}
