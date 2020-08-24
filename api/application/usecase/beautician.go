package usecase

import (
	"context"

	"github.com/myapp/noname/api/application/usecase/request"
	"github.com/myapp/noname/api/application/usecase/response"

	"github.com/myapp/noname/api/domain/repository"
	"github.com/rs/xid"
)

// Beautician 美容師
type Beautician interface {
	Create(ctx context.Context, r *request.BeauticianCreate) (*response.Beautician, error)
}

type beautician struct {
	repositoryBeautician repository.Beautician
}

// NewBeautician ユースケースの初期化
func NewBeautician(
	repositoryBeautician repository.Beautician,
) Beautician {
	return &beautician{
		repositoryBeautician: repositoryBeautician,
	}
}

func (b *beautician) Create(ctx context.Context, r *request.BeauticianCreate) (*response.Beautician, error) {
	ent := r.NewBeautician(xid.New().String())
	if err := b.repositoryBeautician.Create(ctx, ent); err != nil {
		return nil, err
	}
	res := response.NewBeautician(ent)
	return res, nil
}
