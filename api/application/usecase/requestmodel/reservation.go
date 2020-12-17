package requestmodel

import (
	"time"

	"github.com/myapp/noname/api/domain/entity"
)

// ReservationCreate request構造体
type ReservationCreate struct {
	AuthID       string   `json:"-"`
	BeauticianID string   `json:"beauticianRandId"`
	SalonID      string   `json:"salonRandId"`
	MenuIDs      []string `json:"menuIds"`
	Date         string   `json:"date"`
}

// NewReservation 予約モデル変換メソッド
func (r *ReservationCreate) NewReservation(userID, spaceID, beauticanID int64, date time.Time) *entity.Reservation {
	return &entity.Reservation{
		Date:         date,
		SpaceID:      spaceID,
		BeauticianID: beauticanID,
		UserID:       userID,
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

// ReservationFindByUser ゲスト予約履歴取得構造体
type ReservationFindByUser struct {
	AuthID string `json:"-"`
}
