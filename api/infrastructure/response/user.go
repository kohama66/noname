package response

import (
	"github.com/myapp/noname/api/application/usecase/responsemodel"
	"github.com/myapp/noname/api/domain/entity"
)

// User DIInterface
type User interface {
	NewUserGet(ent *entity.User) *responsemodel.UserGet
	NewUserCreate(ent *entity.User) *responsemodel.UserCreate
}

type user struct{}

// NewUser DI初期化関数
func NewUser() User {
	return &user{}
}

// NewUserResponsemodel レスポンスモデル変換関数
func NewUserResponsemodel(ent *entity.User) *responsemodel.User {
	return &responsemodel.User{
		RandID:          ent.RandID,
		FirstName:       ent.FirstName,
		LastName:        ent.LastName,
		FirstNameKana:   ent.FirstNameKana,
		LastNameKana:    ent.LastNameKana,
		Email:           ent.Email,
		PhoneNumber:     ent.PhoneNumber,
		BeauticianInfo:  ent.R.Beautician,
		BeauticianMenus: NewBeauticianMenusResponsemodel(ent.R.BeauticianBeauticianMenus),
		CreatedAt:       ent.CreatedAt,
		UpdatedAt:       ent.CreatedAt,
	}
}

func (g *user) NewUserGet(ent *entity.User) *responsemodel.UserGet {
	return &responsemodel.UserGet{
		User: NewUserResponsemodel(ent),
	}
}

func (g *user) NewUserCreate(ent *entity.User) *responsemodel.UserCreate {
	return &responsemodel.UserCreate{
		User: NewUserResponsemodel(ent),
	}
}