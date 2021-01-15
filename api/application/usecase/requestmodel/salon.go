package requestmodel

import "time"

// SalonFind 美容師検索リクエスト構造体
type SalonFind struct {
	BeauticianRandID *string    `schema:"beauticianRandId"`
	Date             *time.Time `schema:"date"`
}

// SalonFindNotBelongs 美容師が所属してない美容院検索
type SalonFindNotBelongs struct {
	AuthID string `json:"-"`
}
