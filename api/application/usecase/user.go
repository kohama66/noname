package usecase

import (
	"context"

	"github.com/myapp/noname/api/application/usecase/requestmodel"
	"github.com/myapp/noname/api/application/usecase/responsemodel"
	"github.com/myapp/noname/api/domain/repository"
	"github.com/myapp/noname/api/infrastructure/response"
	"github.com/rs/xid"
)

// User DIInterface
type User interface {
	Get(ctx context.Context, r *requestmodel.UserGet) (*responsemodel.UserGet, error)
	Create(ctx context.Context, r *requestmodel.UserCreate) (*responsemodel.UserCreate, error)
}

type user struct {
	userRepository        repository.User
	userResponse          response.User
	reservationRepository repository.Reservation
	salonRepository       repository.Salon
}

// NewUser DI初期化関数
func NewUser(
	userRepository repository.User,
	userResponse response.User,
	reservationRepository repository.Reservation,
	salonRepository repository.Salon,
) User {
	return &user{
		userRepository:        userRepository,
		userResponse:          userResponse,
		reservationRepository: reservationRepository,
		salonRepository:       salonRepository,
	}
}

func (g *user) Get(ctx context.Context, r *requestmodel.UserGet) (*responsemodel.UserGet, error) {
	u, err := g.userRepository.GetByAuthID(ctx, r.AuthID)
	if err != nil {
		return nil, err
	}
	return g.userResponse.NewUserGet(u), nil
}

func (g *user) Create(ctx context.Context, r *requestmodel.UserCreate) (*responsemodel.UserCreate, error) {
	gs := r.NewUser(xid.New().String())
	if err := g.userRepository.Create(ctx, gs); err != nil {
		return nil, err
	}
	ent, err := g.userRepository.GetByAuthID(ctx, r.AuthID)
	if err != nil {
		return nil, err
	}
	return g.userResponse.NewUserCreate(ent), nil
}
