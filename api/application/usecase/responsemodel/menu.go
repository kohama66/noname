package responsemodel

import "time"

// Menu response構造体
type Menu struct {
	ID        int64     `json:"id"`
	RandID    string    `json:"randId"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// MenuFind メニュー検索構造体
type MenuFind struct {
	Menus []*Menu `json:"menus"`
}
