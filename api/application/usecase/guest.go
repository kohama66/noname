package usecase

import (
	"context"

	"github.com/myapp/noname/api/application/usecase/requestmodel"
	"github.com/myapp/noname/api/application/usecase/responsemodel"
	"github.com/myapp/noname/api/domain/repository"
	"github.com/myapp/noname/api/infrastructure/response"
	"github.com/rs/xid"
)

// Guest DIInterface
type Guest interface {
	Get(ctx context.Context, r *requestmodel.GuestGet) (*responsemodel.GuestGet, error)
	Create(ctx context.Context, r *requestmodel.GuestCreate) (*responsemodel.GuestCreate, error)
}

type guest struct {
	guestRepository       repository.Guest
	guestResponse         response.Guest
	reservationRepository repository.Reservation
}

// NewGuest DI初期化関数
func NewGuest(
	guestRepository repository.Guest,
	guestResponse response.Guest,
	reservationRepository repository.Reservation,
) Guest {
	return &guest{
		guestRepository:       guestRepository,
		guestResponse:         guestResponse,
		reservationRepository: reservationRepository,
	}
}

func (g *guest) Get(ctx context.Context, r *requestmodel.GuestGet) (*responsemodel.GuestGet, error) {
	gs, err := g.guestRepository.GetByAuthID(ctx, r.AuthID)
	if err != nil {
		return nil, err
	}
	rs, err := g.reservationRepository.FindByGuest(ctx, gs.ID)
	if err != nil {
		return nil, err
	}
	return g.guestResponse.NewGuestGet(gs, rs), nil
}

func (g *guest) Create(ctx context.Context, r *requestmodel.GuestCreate) (*responsemodel.GuestCreate, error) {
	gs := r.NewGuest(xid.New().String())
	if err := g.guestRepository.Create(ctx, gs); err != nil {
		return nil, err
	}
	return g.guestResponse.NewGuestCreate(gs), nil
}
