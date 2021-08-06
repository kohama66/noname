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
func (r *ReservationCreate) NewReservation(userID, spaceID, beauticanID int64, randID string, date time.Time) *entity.Reservation {
	return &entity.Reservation{
		RandID:       randID,
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

// ReservationGetInfo 予約詳細取得
type ReservationGetInfo struct {
	RandID string `json:"randId"`
}

// ReservationSetHoliday 美容師休日設定
type ReservationSetHoliday struct {
	AuthID  string    `json:"-"`
	Holiday time.Time `json:"holiday"`
}

// NewReservationSetHoliday 美容師休日エンティティ変換関数
func (r *ReservationSetHoliday) NewReservationSetHoliday(randID string, beauticianID int64) *entity.Reservation {
	return &entity.Reservation{
		RandID:       randID,
		Date:         r.Holiday,
		Holiday:      true,
		SpaceID:      1,
		BeauticianID: beauticianID,
		UserID:       1,
	}
}

// ReservationCreateBySalonDayOff 美容院休日設定
type ReservationCreateBySalonDayOff struct {
	RandID string    `json:"randId"`
	AuthID string    `json:"-"`
	Date   time.Time `json:"date"`
}

func (r *ReservationCreateBySalonDayOff) NewReservation() {
	return &entity.Reservation{}
}
