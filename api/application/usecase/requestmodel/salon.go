package requestmodel

import (
	"time"

	"github.com/myapp/noname/api/domain/entity"
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

// BeauticianSalonDelete 美容師の行ける美容院削除
type BeauticianSalonDelete struct {
	AuthID      string `json:"-"`
	SalonRandID string `json:"salonRandId"`
}

// NewBeauticianSalon ファクトリ関数
func (b *BeauticianSalonDelete) NewBeauticianSalon(beauticianID, salonID int64) *entity.BeauticianSalon {
	return &entity.BeauticianSalon{
		BeauticianID: beauticianID,
		SalonID:      salonID,
	}
}
