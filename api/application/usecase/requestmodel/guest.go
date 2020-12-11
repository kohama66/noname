package requestmodel

import "github.com/myapp/noname/api/domain/entity"

// GuestGet ゲスト情報取得
type GuestGet struct {
	AuthID string `json:"-"`
}

// GuestCreate ゲスト新規登録リクエスト構造体
type GuestCreate struct {
	AuthID        string `json:"-"`
	LastName      string `json:"lastName" validate:"required"`
	FirstName     string `json:"firstName" validate:"required"`
	LastNameKana  string `json:"lastNameKana" validate:"required"`
	FirstNameKana string `json:"firstNameKana" validate:"required"`
	Email         string `json:"email" validate:"required,email"`
	PhoneNumber   string `json:"phoneNumber" validate:"required,len=11"`
}

// NewGuest リクエストモデルエンティティ変換メソッド
func (g *GuestCreate) NewGuest(randID string) *entity.Guest {
	return &entity.Guest{
		AuthID:        g.AuthID,
		RandID:        randID,
		LastName:      g.LastName,
		FirstName:     g.FirstName,
		LastNameKana:  g.LastNameKana,
		FirstNameKana: g.FirstNameKana,
		Email:         g.Email,
	}
}
