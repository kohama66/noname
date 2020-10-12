package requestmodel

// SalonFind 美容師検索リクエスト構造体
type SalonFind struct {
	BeauticianRandID string `schema:"beauticianRandId"`
	Date             string `schema:"date"`
}
