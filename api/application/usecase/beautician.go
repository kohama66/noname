package usecase

import (
	"context"

	"github.com/myapp/noname/api/application/usecase/requestmodel"
	"github.com/myapp/noname/api/application/usecase/responsemodel"
	"github.com/myapp/noname/api/infrastructure/response"

	"github.com/myapp/noname/api/domain/repository"
	"github.com/rs/xid"
)

// Beautician 美容師
type Beautician interface {
	Create(ctx context.Context, r *requestmodel.BeauticianCreate) (*responsemodel.BeauticianCreate, error)
	Get(ctx context.Context, r *requestmodel.BeauticianGet) (*responsemodel.BeauticianGet, error)
}

type beautician struct {
	beauticianRepository repository.Beautician
	beauticianResponse   response.Beautician
}

// NewBeautician usecaseの初期化
func NewBeautician(
	beauticianRepository repository.Beautician,
	beauticianResponse response.Beautician,
) Beautician {
	return &beautician{
		beauticianRepository: beauticianRepository,
		beauticianResponse:   beauticianResponse,
	}
}

func (b *beautician) Create(ctx context.Context, r *requestmodel.BeauticianCreate) (*responsemodel.BeauticianCreate, error) {
	ent := r.NewBeautician(xid.New().String())
	if err := b.beauticianRepository.Create(ctx, ent); err != nil {
		return nil, err
	}
	res := b.beauticianResponse.NewBeauticianCreate(ent)
	return res, nil
}

func (b *beautician) Get(ctx context.Context, r *requestmodel.BeauticianGet) (*responsemodel.BeauticianGet, error) {
	ent, err := b.beauticianRepository.GetByAuthID(ctx, r.AuthID)
	if err != nil {
		return nil, err
	}
	res := b.beauticianResponse.NewBeauticianGet(ent)
	return res, nil
}
