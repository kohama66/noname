package response

import (
	"github.com/myapp/noname/api/application/usecase/responsemodel"
	"github.com/myapp/noname/api/domain/entity"
)

type Beautician interface {
	NewResponseModelBeautician(ent *entity.Beautician) *responsemodel.Beautician
	NewBeauticianCreate(ent *entity.Beautician) *responsemodel.BeauticianCreate
	NewBeauticianGet(ent *entity.Beautician) *responsemodel.BeauticianGet
}

type beautician struct {
}

func NewBeautician() Beautician {
	return &beautician{}
}

// NewResponseModelBeautician エンティティーをレスポンスへ変換する
func (b *beautician) NewResponseModelBeautician(ent *entity.Beautician) *responsemodel.Beautician {
	return &responsemodel.Beautician{
		ID:          ent.ID,
		RandID:      ent.RandID,
		FirstName:   ent.FirstName,
		LastName:    ent.LastName,
		Age:         ent.Age,
		PhoneNumber: ent.PhoneNumber,
		CreatedAt:   ent.CreatedAt,
		UpdatedAt:   ent.UpdatedAt,
	}
}

// NewBeauticianCreate レスポンス変換
func (b *beautician) NewBeauticianCreate(ent *entity.Beautician) *responsemodel.BeauticianCreate {
	return &responsemodel.BeauticianCreate{
		Beautician: b.NewResponseModelBeautician(ent),
	}
}

// NewBeauticianGet レスポンス変換
func (b *beautician) NewBeauticianGet(ent *entity.Beautician) *responsemodel.BeauticianGet {
	return &responsemodel.BeauticianGet{
		Beautician: b.NewResponseModelBeautician(ent),
	}
}
