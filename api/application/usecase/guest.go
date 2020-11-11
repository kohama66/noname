package usecase

import (
	"context"

	"github.com/myapp/noname/api/application/usecase/requestmodel"
	"github.com/myapp/noname/api/application/usecase/responsemodel"
	"github.com/myapp/noname/api/domain/repository"
	"github.com/myapp/noname/api/infrastructure/response"
)

// Guest DIInterface
type Guest interface {
	Get(ctx context.Context, r *requestmodel.GuestGet) (*responsemodel.GuestGet, error)
}

type guest struct {
	guestRepository repository.Guest
	guestResponse   response.Guest
}

// NewGuest DI初期化関数
func NewGuest(
	guestRepository repository.Guest,
	guestResponse response.Guest,
) Guest {
	return &guest{
		guestRepository: guestRepository,
		guestResponse:   guestResponse,
	}
}

func (g *guest) Get(ctx context.Context, r *requestmodel.GuestGet) (*responsemodel.GuestGet, error) {
	gs, err := g.guestRepository.GetByAuthID(ctx, r.AuthID)
	if err != nil {
		return nil, err
	}
	return g.guestResponse.NewGuestGet(gs), nil
}
