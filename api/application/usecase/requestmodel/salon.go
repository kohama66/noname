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

// SalonCreate 美容院作成
type SalonCreate struct {
	AuthID       string `json:"-"`
	Name         string `json:"name" validate:"required"`
	PhoneNumber  string `json:"phoneNumber" validate:"required,len=11"`
	PostalCode   string `json:"postalCode" validate:"required,len=7,numeric"`
	Prefectures  string `json:"prefectures" validate:"required"`
	City         string `json:"city" validate:"required"`
	Town         string `json:"town" validate:"required"`
	AddressCode  string `json:"addressCode" validate:"required"`
	AddressOther string `json:"addressOther"`
	OpenHour     string `json:"openingHours" validate:"required"`
	CloseHour    string `json:"closingHours" validate:"required"`
}

// NewSalon ファクトリ関数
func (s *SalonCreate) NewSalon(randID string) *entity.Salon {
	return &entity.Salon{
		RandID:       randID,
		Name:         s.Name,
		PhoneNumber:  s.PhoneNumber,
		OpeningHours: s.OpenHour,
		ClosingHours: s.CloseHour,
		PostalCode:   s.PostalCode,
		Prefectures:  s.Prefectures,
		City:         s.City,
		Town:         s.Town,
		AddressCode:  s.AddressCode,
		AddressOther: s.AddressOther,
	}
}

// SalonMyPageGet 美容院マイページ取得
type SalonMyPageGet struct {
	RandID string `json:"randId"`
	AuthID string `json:"-"`
}
