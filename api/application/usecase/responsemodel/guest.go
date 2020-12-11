package responsemodel

import "time"

// Guest ゲストレスポンス構造体
type Guest struct {
	RandID        string    `json:"randId"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	FirstNameKana string    `json:"firstNameKana"`
	LastNameKana  string    `json:"lastNameKana"`
	Email         string    `json:"email"`
	PhoneNumber   string    `json:"phoneNumber"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	// Reservations []*Reservation `json:"reservations"`
}

// GuestGet ゲスト情報取得レスポンス構造体
// type GuestGet struct {
// 	RandID       string                   `json:"randId"`
// 	FirstName    string                   `json:"firstName"`
// 	LastName     string                   `json:"lastName"`
// 	Reservations []*ReservationGetByGuest `json:"reservations"`
// 	CreatedAt    time.Time                `json:"createdAt"`
// 	UpdatedAt    time.Time                `json:"updatedAt"`
// }

// GuestGet ゲスト情報取得レスポンス構造体
type GuestGet struct {
	*Guest `json:"guest"`
}

// GuestCreate ゲスト新規登録レスポンス構造体
type GuestCreate struct {
	Guest *Guest `json:"guest"`
}
