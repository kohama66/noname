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
