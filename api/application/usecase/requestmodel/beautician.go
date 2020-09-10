package requestmodel

import "github.com/myapp/noname/api/domain/entity"

// BeauticianCreate 美容師登録構造体
type BeauticianCreate struct {
	AuthID      string `json:"-"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Age         int64  `json:"age"`
	PhoneNumber string `json:"phoneNumber"`
}

// NewBeautician 美容師登録モデル変換メソッド
func (b BeauticianCreate) NewBeautician(randID string) *entity.Beautician {
	return &entity.Beautician{
		AuthID:      b.AuthID,
		RandID:      randID,
		FirstName:   b.FirstName,
		LastName:    b.LastName,
		Age:         b.Age,
		PhoneNumber: b.PhoneNumber,
	}
}

// BeauticianGet 美容師情報取得構造体
type BeauticianGet struct {
	AuthID string `json:"-"`
}
