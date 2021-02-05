package response

import (
	"github.com/myapp/noname/api/application/usecase/responsemodel"
	"github.com/myapp/noname/api/domain/entity"
)

// Salon DIInterface
type Salon interface {
	NewSalonFind(ents []*entity.Salon) *responsemodel.SalonFind
	NewSalonFindNotBelongs(ents []*entity.Salon) *responsemodel.SalonFindNotBelongs
}

type salon struct{}

// NewSalon DI初期化関数
func NewSalon() Salon {
	return &salon{}
}

// NewSalonResponseModel エンティティーをレスポンスへ変換
func NewSalonResponseModel(ent *entity.Salon) *responsemodel.Salon {
	return &responsemodel.Salon{
		ID:           ent.ID,
		RandID:       ent.RandID,
		Name:         ent.Name,
		PhoneNumber:  ent.PhoneNumber,
		OpeningHours: ent.OpeningHours,
		ClosingHours: ent.ClosingHours,
		PostalCode:   ent.PostalCode,
		Prefectures:  ent.Prefectures,
		City:         ent.City,
		Town:         ent.Town,
		AddressCode:  ent.AddressCode,
		AddressOther: ent.AddressOther,
		CreatedAt:    ent.CreatedAt,
		UpdatedAt:    ent.UpdatedAt,
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

// NewSalonsResponseModel エンティティースライスをレスポンススライスへ変換
func NewSalonsResponseModel(ents []*entity.Salon) []*responsemodel.Salon {
	sls := make([]*responsemodel.Salon, len(ents))
	for i, v := range ents {
		sls[i] = NewSalonResponseModel(v)
	}
	return sls
}

// NewSalonFindNotBelongs 美容師が所属してない美容院検索
func (s *salon) NewSalonFindNotBelongs(ents []*entity.Salon) *responsemodel.SalonFindNotBelongs {
	sls := make([]*responsemodel.Salon, len(ents))
	for i, v := range ents {
		sls[i] = NewSalonResponseModel(v)
	}
	return &responsemodel.SalonFindNotBelongs{
		Salons: sls,
	}
}
