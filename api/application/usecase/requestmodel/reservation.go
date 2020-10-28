package requestmodel

import (
	"time"

	"github.com/myapp/noname/api/domain/entity"
)

// ReservationCreate request構造体
type ReservationCreate struct {
	AuthID       string    `json:"-"`
	BeauticianID int64     `json:"beauticiaId"`
	SpaceID      int64     `json:"spaceId"`
	MenuID       int64     `json:"menuId"`
	Date         time.Time `json:"date"`
}

// NewReservation 予約モデル変換メソッド
func (r *ReservationCreate) NewReservation(guestID int64) *entity.Reservation {
	return &entity.Reservation{
		Date:         r.Date,
		SpaceID:      r.SpaceID,
		BeauticianID: r.BeauticianID,
		GuestID:      guestID,
		MenuID:       r.MenuID,
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
