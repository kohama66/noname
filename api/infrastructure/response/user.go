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
		RandID:        ent.RandID,
		FirstName:     ent.FirstName,
		LastName:      ent.LastName,
		FirstNameKana: ent.FirstNameKana,
		LastNameKana:  ent.LastNameKana,
		Email:         ent.Email,
		PhoneNumber:   ent.PhoneNumber,
		IsBeauticina:  ent.IsBeautician,
		CreatedAt:     ent.CreatedAt,
		UpdatedAt:     ent.CreatedAt,
	}
}

// NewBeauticianUserResponsemodel レスポンスモデル変換関数
func NewBeauticianUserResponsemodel(ent *entity.User) *responsemodel.User {
	return &responsemodel.User{
		RandID:          ent.RandID,
		FirstName:       ent.FirstName,
		LastName:        ent.LastName,
		FirstNameKana:   ent.FirstNameKana,
		LastNameKana:    ent.LastNameKana,
		Email:           ent.Email,
		PhoneNumber:     ent.PhoneNumber,
		IsBeauticina:    ent.IsBeautician,
		BeauticianInfo:  NewBeauticianResponseModel(ent.R.Beautician),
		BeauticianMenus: NewBeauticianMenusResponsemodel(ent.R.BeauticianBeauticianMenus),
		// BeauticianSalons: NewSalonsResponseModel(salons),
		CreatedAt: ent.CreatedAt,
		UpdatedAt: ent.CreatedAt,
	}
}

func (u *user) NewUserGet(ent *entity.User) *responsemodel.UserGet {
	if ent.IsBeautician {
		return &responsemodel.UserGet{
			User: NewBeauticianUserResponsemodel(ent),
		}
	}
	return &responsemodel.UserGet{
		User: NewUserResponsemodel(ent),
	}
}

func (u *user) NewUserCreate(ent *entity.User) *responsemodel.UserCreate {
	return &responsemodel.UserCreate{
		User: NewUserResponsemodel(ent),
	}
}
