package responsemodel

import (
	"time"

	"github.com/myapp/noname/api/domain/entity"
)

// User ユーザーレスポンス構造体
type User struct {
	RandID          string             `json:"randId"`
	FirstName       string             `json:"firstName"`
	LastName        string             `json:"lastName"`
	FirstNameKana   string             `json:"firstNameKana"`
	LastNameKana    string             `json:"lastNameKana"`
	Email           string             `json:"email"`
	PhoneNumber     string             `json:"phoneNumber"`
	BeauticianInfo  *entity.Beautician `json:"beauticianInfo"`
	BeauticianMenus []*BeauticianMenu  `json:"menus"`
	CreatedAt       time.Time          `json:"createdAt"`
	UpdatedAt       time.Time          `json:"updatedAt"`
	// Reservations []*Reservation `json:"reservations"`
}

// UserGet ゲスト情報取得レスポンス構造体
// type UserGet struct {
// 	RandID       string                   `json:"randId"`
// 	FirstName    string                   `json:"firstName"`
// 	LastName     string                   `json:"lastName"`
// 	Reservations []*ReservationGetByUser `json:"reservations"`
// 	CreatedAt    time.Time                `json:"createdAt"`
// 	UpdatedAt    time.Time                `json:"updatedAt"`
// }

// UserGet ユーザー情報取得レスポンス構造体
type UserGet struct {
	*User `json:"user"`
}

// UserCreate ゲスト新規登録レスポンス構造体
type UserCreate struct {
	User *User `json:"user"`
}
