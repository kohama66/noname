package usecase

import (
	"context"

	"github.com/myapp/noname/api/application/usecase/requestmodel"
	"github.com/myapp/noname/api/application/usecase/responsemodel"
	"github.com/myapp/noname/api/domain/repository"
	"github.com/myapp/noname/api/infrastructure/response"
)

// Salon DIInterface
type Salon interface {
	Find(ctx context.Context, r *requestmodel.SalonFind) (*responsemodel.SalonFind, error)
}

type salon struct {
	salonRepository repository.Salon
	salonResponse   response.Salon
	userRepository  repository.User
}

// NewSalon DI初期化関数
func NewSalon(
	salonRepository repository.Salon,
	salonResponse response.Salon,
	userRepository repository.User,
) Salon {
	return &salon{
		salonRepository: salonRepository,
		salonResponse:   salonResponse,
		userRepository:  userRepository,
	}
}

func (s *salon) Find(ctx context.Context, r *requestmodel.SalonFind) (*responsemodel.SalonFind, error) {
	var beauticianID *int64
	if r.BeauticianRandID != nil {
		us, err := s.userRepository.GetByRandID(ctx, *r.BeauticianRandID)
		if err != nil {
			return nil, err
		}
		beauticianID = &us.ID
	}
	sl, err := s.salonRepository.Find(ctx, beauticianID, r.Date)
	if err != nil {
		return nil, err
	}
	return s.salonResponse.NewSalonFind(sl), nil
}
