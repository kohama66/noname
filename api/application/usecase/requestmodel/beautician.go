package requestmodel

import (
	"github.com/myapp/noname/api/domain/entity"
)

// BeauticianCreate 美容師登録構造体
type BeauticianCreate struct {
	AuthID        string `json:"-"`
	LastName      string `json:"lastName" validate:"required"`
	FirstName     string `json:"firstName" validate:"required"`
	LastNameKana  string `json:"lastNameKana" validate:"required"`
	FirstNameKana string `json:"firstNameKana" validate:"required"`
	Email         string `json:"email" validate:"required,email"`
	// PhoneNumber   string `json:"phoneNumber" validate:"required,len=11"`
}

// NewBeautician 美容師登録モデル変換メソッド
func (b BeauticianCreate) NewBeautician(randID string) *entity.Beautician {
	return &entity.Beautician{
		AuthID:        b.AuthID,
		RandID:        randID,
		LastName:      b.LastName,
		FirstName:     b.FirstName,
		LastNameKana:  b.LastNameKana,
		FirstNameKana: b.FirstNameKana,
		Email:         b.Email,
	}
}

// // BeauticianGet 美容師情報取得構造体
// type BeauticianGet struct {
// 	AuthID string `json:"-"`
// }

type BeauticianFind struct {
	SalonRandID string   `schema:"salonRandId"`
	MenuRandIDs []string `schema:"menuRandIds"`
	Date        string   `schema:"date"`
}

// BeauticianGet 美容師情報取得構造体
type BeauticianGet struct {
	RandID string `json:"-"`
}
