package responsemodel

import "time"

// Guest ゲストレスポンス構造体
type Guest struct {
	RandID       string         `json:"randId"`
	FirstName    string         `json:"firstName"`
	LastName     string         `json:"lastName"`
	Reservations []*Reservation `json:"reservations"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
}

// GuestGet ゲスト情報取得レスポンス構造体
type GuestGet struct {
	RandID       string                   `json:"randId"`
	FirstName    string                   `json:"firstName"`
	LastName     string                   `json:"lastName"`
	Reservations []*ReservationGetByGuest `json:"reservations"`
	CreatedAt    time.Time                `json:"createdAt"`
	UpdatedAt    time.Time                `json:"updatedAt"`
}
