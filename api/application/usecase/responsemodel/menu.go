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

// BeauticianMenu response構造体
type BeauticianMenu struct {
	Price        int64     `json:"price"`
	BeauticianID int64     `json:"beauticianId"`
	MenuID       int64     `json:"menuId"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// MenuFindByBeauticianWithMenuRandIDs 美容師の詳細メニュー取得
type MenuFindByBeauticianWithMenuRandIDs struct {
	// Price        int64  `json:"price"`
	// Name         string `json:"name"`
	// BeauticianID int64  `json:"beauticianId"`
	// MenuID       int64  `json:"menuId"`
	BeauticianMenus []*BeauticianMenu `json:"beauticianMenus"`
}
