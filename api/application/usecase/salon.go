package usecase

import (
	"context"

	"github.com/myapp/noname/api/application/usecase/requestmodel"
	"github.com/myapp/noname/api/application/usecase/responsemodel"
	"github.com/myapp/noname/api/domain/entity"
	"github.com/myapp/noname/api/domain/repository"
	"github.com/myapp/noname/api/infrastructure/response"
)

// Salon DIInterface
type Salon interface {
	Find(ctx context.Context, r *requestmodel.SalonFind) (*responsemodel.SalonFind, error)
}

type salon struct {
	salonRepository      repository.Salon
	salonResponse        response.Salon
	beauticianRepository repository.Beautician
}

// NewSalon DI初期化関数
func NewSalon(
	salonRepository repository.Salon,
	salonResponse response.Salon,
	beauticianRepository repository.Beautician,
) Salon {
	return &salon{
		salonRepository:      salonRepository,
		salonResponse:        salonResponse,
		beauticianRepository: beauticianRepository,
	}
}

func (s *salon) Find(ctx context.Context, r *requestmodel.SalonFind) (*responsemodel.SalonFind, error) {
	var sls entity.SalonSlice
	// var i int64
	// i = 3
	if !r.Date.IsZero() {
		sl, err := s.salonRepository.Find(ctx, nil, &r.Date)
		if err != nil {
			return nil, err
		}
		sls = append(sls, sl...)
	}
	// sl, err := s.salonRepository.Find(ctx, &i, nil)
	// if err != nil {
	// 	return nil, err
	// }
	// sls = append(sls, sl...)
	return s.salonResponse.NewSalonFind(sls), nil

	// var salons entity.SalonSlice
	// if r.BeauticianRandID != "" {
	// 	bt, err := s.beauticianRepository.GetByRandID(ctx, r.BeauticianRandID)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	sl, err := s.salonRepository.FindByBeautician(ctx, bt.ID)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	salons = sl
	// } else {
	// 	sl, err := s.salonRepository.GetAll(ctx)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	salons = sl
	// }
	// return s.salonResponse.NewSalonFind(salons), nil
}
