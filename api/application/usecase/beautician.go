package usecase

import (
	"context"

	"github.com/myapp/noname/api/application/usecase/request"
)

type Beautician interface {
	Create(ctx context.Context, r *request.BeauticianCreate) (string, error)
}

type beautician struct {
}

func NewBeautician() Beautician {
	return &beautician{}
}

func (b *beautician) Create(ctx context.Context, r *request.BeauticianCreate) (string, error) {
	return "", nil
}
