package response

import (
	"github.com/myapp/noname/api/application/usecase/responsemodel"
	"github.com/myapp/noname/api/domain/entity"
)

// Beautician DIInterface
type Beautician interface {
	NewBeauticianCreate(ent *entity.Beautician) *responsemodel.BeauticianCreate
	NewBeauticianGet(ent *entity.Beautician) *responsemodel.BeauticianGet
}

type beautician struct {
}

// NewBeautician DI初期化関数
func NewBeautician() Beautician {
	return &beautician{}
}

// NewResponseModelBeautician エンティティーをレスポンスへ変換
func NewResponseModelBeautician(ent *entity.Beautician) *responsemodel.Beautician {
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
		Beautician: NewResponseModelBeautician(ent),
	}
}

// NewBeauticianGet レスポンス変換
func (b *beautician) NewBeauticianGet(ent *entity.Beautician) *responsemodel.BeauticianGet {
	return &responsemodel.BeauticianGet{
		Beautician: NewResponseModelBeautician(ent),
	}
}
