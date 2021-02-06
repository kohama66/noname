package usecase

import (
	"context"
	"fmt"

	"github.com/myapp/noname/api/application/usecase/requestmodel"
	"github.com/myapp/noname/api/application/usecase/responsemodel"
	"github.com/myapp/noname/api/domain/entityx"
	"github.com/myapp/noname/api/domain/repository"
	"github.com/myapp/noname/api/infrastructure/response"
	"github.com/rs/xid"
)

// Salon DIInterface
type Salon interface {
	Find(ctx context.Context, r *requestmodel.SalonFind) (*responsemodel.SalonFind, error)
	FindNotBelongs(ctx context.Context, r *requestmodel.SalonFindNotBelongs) (*responsemodel.SalonFindNotBelongs, error)
	CreateToBeautician(ctx context.Context, r *requestmodel.BeauticianSalonCreata) error
	DeleteToBeautician(ctx context.Context, r *requestmodel.BeauticianSalonDelete) error
	Create(ctx context.Context, r *requestmodel.SalonCreate) (*responsemodel.SalonCreate, error)
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

func (s *salon) FindNotBelongs(ctx context.Context, r *requestmodel.SalonFindNotBelongs) (*responsemodel.SalonFindNotBelongs, error) {
	me, err := s.userRepository.GetByAuthID(ctx, r.AuthID)
	if err != nil {
		return nil, err
	}
	if !me.IsBeautician {
		return nil, fmt.Errorf("You are not beautician")
	}
	sa, err := s.salonRepository.FindNotBelongs(ctx, me.ID)
	if err != nil {
		return nil, err
	}
	return s.salonResponse.NewSalonFindNotBelongs(sa), nil
}

func (s *salon) CreateToBeautician(ctx context.Context, r *requestmodel.BeauticianSalonCreata) error {
	me, err := s.userRepository.GetByAuthID(ctx, r.AuthID)
	if err != nil {
		return err
	}
	if !me.IsBeautician {
		return fmt.Errorf("You are not beautician")
	}
	sl, err := s.salonRepository.GetByRandID(ctx, r.SalonRandID)
	if err != nil {
		return err
	}
	ng, err := s.salonRepository.ExistsByBeauticianWithSalon(ctx, me.ID, sl.ID)
	if err != nil {
		return err
	}
	if ng {
		return fmt.Errorf("Already exists")
	}
	if err := s.salonRepository.CreateBeauticianSalon(ctx, entityx.NewBeauticianSalon(me.ID, sl.ID)); err != nil {
		return err
	}
	return nil
}

func (s *salon) DeleteToBeautician(ctx context.Context, r *requestmodel.BeauticianSalonDelete) error {
	me, err := s.userRepository.GetByAuthID(ctx, r.AuthID)
	if err != nil {
		return err
	}
	sa, err := s.salonRepository.GetByRandID(ctx, r.SalonRandID)
	if err != nil {
		return err
	}
	if _, err := s.salonRepository.DeleteBeauticianSalon(ctx, r.NewBeauticianSalon(me.ID, sa.ID)); err != nil {
		return err
	}
	return nil
}

func (s *salon) Create(ctx context.Context, r *requestmodel.SalonCreate) (*responsemodel.SalonCreate, error) {
	sa := r.NewSalon(xid.New().String())
	if err := s.salonRepository.Create(ctx, sa); err != nil {
		return nil, err
	}
	return s.salonResponse.NewSalonCreate(sa), nil
}
