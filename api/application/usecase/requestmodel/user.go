package requestmodel

import "github.com/myapp/noname/api/domain/entity"

// UserGet ユーザー情報取得
type UserGet struct {
	AuthID string `json:"-"`
}

// UserCreate ゲスト新規登録リクエスト構造体
type UserCreate struct {
	AuthID        string `json:"-"`
	LastName      string `json:"lastName" validate:"required"`
	FirstName     string `json:"firstName" validate:"required"`
	LastNameKana  string `json:"lastNameKana" validate:"required"`
	FirstNameKana string `json:"firstNameKana" validate:"required"`
	Email         string `json:"email" validate:"required,email"`
	// PhoneNumber   string `json:"phoneNumber" validate:"required,len=11"`
}

// NewUser リクエストモデルエンティティ変換メソッド
func (g *UserCreate) NewUser(randID string) *entity.User {
	return &entity.User{
		AuthID:        g.AuthID,
		RandID:        randID,
		LastName:      g.LastName,
		FirstName:     g.FirstName,
		LastNameKana:  g.LastNameKana,
		FirstNameKana: g.FirstNameKana,
		Email:         g.Email,
	}
}
