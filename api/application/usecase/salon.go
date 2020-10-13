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
}

// NewSalon DI初期化関数
func NewSalon(
	salonRepository repository.Salon,
	salonResponse response.Salon,
) Salon {
	return &salon{
		salonRepository: salonRepository,
		salonResponse:   salonResponse,
	}
}

func (s *salon) Find(ctx context.Context, r *requestmodel.SalonFind) (*responsemodel.SalonFind, error) {
	ents, err := s.salonRepository.Find(ctx)
	if err != nil {
		return nil, err
	}
	return s.salonResponse.NewSalonFind(ents), nil
}
