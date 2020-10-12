package response

import (
	"github.com/myapp/noname/api/application/usecase/responsemodel"
	"github.com/myapp/noname/api/domain/entity"
)

// Salon DIInterface
type Salon interface {
	NewSalonFind(ents []*entity.Salon) *responsemodel.SalonFind
}

type salon struct{}

// NewSalon DI初期化関数
func NewSalon() Salon {
	return &salon{}
}

// NewSalonResponseModel エンティティーをレスポンスへ変換
func NewSalonResponseModel(ent *entity.Salon) *responsemodel.Salon {
	return &responsemodel.Salon{
		ID:        ent.ID,
		RandID:    ent.RandID,
		Name:      ent.Name,
		CreatedAt: ent.CreatedAt,
		UpdatedAt: ent.UpdatedAt,
	}
}

func (s *salon) NewSalonFind(ents []*entity.Salon) *responsemodel.SalonFind {
	sl := make([]*responsemodel.Salon, len(ents))
	for i, v := range ents {
		sl[i] = NewSalonResponseModel(v)
	}
	return &responsemodel.SalonFind{
		Salons: sl,
	}
}
