package requestmodel

import "github.com/myapp/noname/api/domain/entity"

// MenuFind メニュー検索
type MenuFind struct {
	BeauticianRandID string `schema:"beauticianRandId"`
}

// MenuFindByBeauticianWithMenuRandIDs 美容師の詳細メニュー取得
type MenuFindByBeauticianWithMenuRandIDs struct {
	BeauticianRandID string   `json:"-"`
	MenuRandIDs      []string `schema:"menuRandIds"`
}

// BeauticianMenuCreate 美容師のメニュー作成
type BeauticianMenuCreate struct {
	AuthID       string `json:"-"`
	MenuName     string `json:"menuName" validate:"required"`
	MenuCategory string `json:"menuCategory" validate:"required"`
	Price        int64  `json:"price" validate:"required"`
}

// NewBeauticianMenu エンティティ変換関数
func (m *BeauticianMenuCreate) NewBeauticianMenu(beauticianID, menuID int64) *entity.BeauticianMenu {
	return &entity.BeauticianMenu{
		Name:         m.MenuName,
		Price:        m.Price,
		BeauticianID: beauticianID,
		MenuID:       menuID,
	}
}
