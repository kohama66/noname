package responsemodel

import (
	"time"
)

// Reservation response構造体
type Reservation struct {
	RandID       string    `json:"randId"`
	Date         time.Time `json:"date"`
	SpaceID      int64     `json:"spaceId"`
	UserID       int64     `json:"userId"`
	BeauticianID int64     `json:"beauticiaId"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// ReservationCreate 予約作成 構造体
type ReservationCreate struct {
	*Reservation `json:"reservation"`
}

// ReservationFindByBeautician 美容師予約情報取得 構造体
type ReservationFindByBeautician struct {
	Reservations []*Reservation `json:"reservations"`
}

// ReservationFind 予約検索 構造体
type ReservationFind struct {
	Reservations []*Reservation `json:"reservations"`
}

// ReservationGetByUser 用構造体
type ReservationGetByUser struct {
	ID                  int64             `json:"id"`
	Date                time.Time         `json:"date"`
	UserID              int64             `json:"userId"`
	SalonName           string            `json:"salonName"`
	BeauticianFirstName string            `json:"beauticianFirstName"`
	BeauticianLatsName  string            `json:"beauticianLastName"`
	Menus               []*BeauticianMenu `json:"menus"`
}

type ReservationFindByUser struct {
	Reservations []*ReservationGetByUser `json:"reservations"`
}

// ReservationInfo 予約詳細取得
type ReservationInfo struct {
	RandID    string            `json:"randId"`
	Date      time.Time         `json:"date"`
	Salon     *Salon            `json:"salon"`
	User      *User             `json:"user"`
	Menus     []*BeauticianMenu `json:"menus"`
	CreatedAt time.Time         `json:"createdAt"`
	UpdatedAt time.Time         `json:"updatedAt"`
}

// ReservationGetInfo 予約詳細取得
type ReservationGetInfo struct {
	ReservationInfo *ReservationInfo `json:"reservationInfo"`
}

// ReservationSetHoliday 美容師休日設定
type ReservationSetHoliday struct {
	*Reservation `json:"reservation"`
}
