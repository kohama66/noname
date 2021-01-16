package requestmodel

import (
	"time"
)

// SalonFind 美容師検索リクエスト構造体
type SalonFind struct {
	BeauticianRandID *string    `schema:"beauticianRandId"`
	Date             *time.Time `schema:"date"`
}

// SalonFindNotBelongs 美容師が所属してない美容院検索
type SalonFindNotBelongs struct {
	AuthID string `json:"-"`
}

// BeauticianSalonCreata 美容師の行ける美容院追加
type BeauticianSalonCreata struct {
	AuthID      string `json:"-"`
	SalonRandID string `json:"salonRandId" validate:"required"`
}
