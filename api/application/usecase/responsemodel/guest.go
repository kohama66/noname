package responsemodel

import "time"

// Guest ゲストレスポンス構造体
type Guest struct {
	RandID    string    `json:"randId"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// GuestGet ゲスト情報取得レスポンス構造体
type GuestGet struct {
	*Guest `json:"guest"`
}
