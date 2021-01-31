package response

import (
	"github.com/myapp/noname/api/application/usecase/responsemodel"
	"github.com/myapp/noname/api/domain/entity"
)

// Beautician DIInterface
type Beautician interface {
	NewBeauticianCreate(ent *entity.Beautician) *responsemodel.BeauticianCreate
	NewBeauticianGet(ent *entity.Beautician) *responsemodel.BeauticianGet
	NewBeauticianFind(ents []*entity.User) *responsemodel.BeauticianFind
}

type beautician struct {
}

// NewBeautician DI初期化関数
func NewBeautician() Beautician {
	return &beautician{}
}

// NewBeauticianResponseModel エンティティーをレスポンスへ変換
func NewBeauticianResponseModel(ent *entity.Beautician) *responsemodel.Beautician {
	return &responsemodel.Beautician{
		LineID:      ent.LineID.Ptr(),
		InstagramID: ent.InstagramID.Ptr(),
		Comment:     ent.Comment.Ptr(),
		CreatedAt:   ent.CreatedAt,
		UpdatedAt:   ent.UpdatedAt,
	}
}

// NewBeauticianCreate レスポンス変換
func (b *beautician) NewBeauticianCreate(ent *entity.Beautician) *responsemodel.BeauticianCreate {
	return &responsemodel.BeauticianCreate{
		Beautician: NewBeauticianResponseModel(ent),
	}
}

// NewBeauticianGet レスポンス変換
func (b *beautician) NewBeauticianGet(ent *entity.Beautician) *responsemodel.BeauticianGet {
	return &responsemodel.BeauticianGet{
		Beautician: NewBeauticianResponseModel(ent),
	}
}

func (b *beautician) NewBeauticianFind(ents []*entity.User) *responsemodel.BeauticianFind {
	bs := make([]*responsemodel.User, len(ents))
	for i, v := range ents {
		bs[i] = NewBeauticianUserResponsemodel(v)
	}
	return &responsemodel.BeauticianFind{
		Beauticians: bs,
	}
}
