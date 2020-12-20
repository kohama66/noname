package requestmodel

import (
	"github.com/myapp/noname/api/domain/entity"
	"github.com/volatiletech/null"
)

// BeauticianCreate 美容師登録構造体
type BeauticianCreate struct {
	AuthID      string  `json:"-"`
	LineID      *string `json:"lineId"`
	InstagramID *string `json:"instagramId"`
}

// NewBeautician 美容師登録モデル変換メソッド
func (b BeauticianCreate) NewBeautician(ID int64) *entity.Beautician {
	return &entity.Beautician{
		UserID:      ID,
		LineID:      null.StringFromPtr(b.LineID),
		InstagramID: null.StringFromPtr(b.InstagramID),
	}
}

// BeauticianFind 美容師検索
type BeauticianFind struct {
	SalonRandID string   `schema:"salonRandId"`
	MenuRandIDs []string `schema:"menuRandIds"`
	Date        string   `schema:"date"`
}

// BeauticianGet 美容師情報取得構造体
type BeauticianGet struct {
	AuthID string `json:"-"`
}

type BeauticianUpdate struct {
	AuthID        string  `json:"-"`
	LastName      string  `json:"lastName" validate:"required"`
	FirstName     string  `json:"firstName" validate:"required"`
	LastNameKana  string  `json:"lastNameKana" validate:"required"`
	FirstNameKana string  `json:"firstNameKana" validate:"required"`
	PhoneNumber   string  `json:"phoneNumber" validate:"required,len=11"`
	LineID        *string `json:"lineId"`
	InstagramID   *string `json:"instagramId"`
}

func (b *BeauticianUpdate) NewBeautician(user *entity.User, beautician *entity.Beautician) (*entity.User, *entity.Beautician) {
	user.FirstName = b.FirstName
	user.FirstNameKana = b.FirstNameKana
	user.LastName = b.LastName
	user.LastNameKana = b.LastNameKana
	user.PhoneNumber = b.PhoneNumber
	beautician.LineID = null.StringFromPtr(b.LineID)
	beautician.InstagramID = null.StringFromPtr(b.InstagramID)
	return user, beautician
}
