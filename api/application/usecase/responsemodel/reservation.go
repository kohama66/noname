package responsemodel

import (
	"time"
)

// Reservation response構造体
type Reservation struct {
	Date         time.Time `json:"date"`
	SpaceID      int64     `json:"spaceId"`
	UserID       int64     `json:"userId"`
	BeauticianID int64     `json:"beauticiaId"`
	MenuID       int64     `json:"menuId"`
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
