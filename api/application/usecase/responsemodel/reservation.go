package responsemodel

import (
	"time"
)

// Reservation response構造体
type Reservation struct {
	Date         time.Time `json:"date"`
	SpaceID      int64     `json:"spaceId"`
	GuestID      int64     `json:"guestId"`
	BeauticianID int64     `json:"beauticiaId"`
	MenuID       int64     `json:"menuId"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// ReservationCreate 予約作成 構造体
type ReservationCreate struct {
	*Reservation
}

// ReservationFindByBeautician 美容師予約情報取得 構造体
type ReservationFindByBeautician struct {
	Reservations []*Reservation
}

// ReservationFind 予約検索 構造体
type ReservationFind struct {
	Reservations []*Reservation
}
