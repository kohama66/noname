package responsemodel

import (
	"time"

	"github.com/volatiletech/null"
)

// Reservation response構造体
type Reservation struct {
	Date         null.Time   `json:"date"`
	Time         null.String `json:"time"`
	SpaceID      int64       `json:"spaceId"`
	GuestID      int64       `json:"guestId"`
	BeauticianID int64       `json:"beauticiaId"`
	MenuID       int64       `json:"menuId"`
	CreatedAt    time.Time   `json:"createdAt"`
	UpdatedAt    time.Time   `json:"updatedAt"`
}

// ReservationCreate 予約作成response構造体
type ReservationCreate struct {
	*Reservation
}
