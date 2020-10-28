package requestmodel

// MenuFind メニュー検索
type MenuFind struct {
	BeauticianRandID string `schema:"beauticianRandId"`
}

// MenuFindByBeauticianWithMenuRandIDs 美容師の詳細メニュー取得
type MenuFindByBeauticianWithMenuRandIDs struct {
	BeauticianRandID string   `json:"-"`
	MenuRandIDs      []string `schema:"menuRandIds"`
}
