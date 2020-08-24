package usecase

import (
	"context"

	"github.com/myapp/noname/api/application/usecase/request"
	"github.com/myapp/noname/api/application/usecase/response"
)

type Test interface {
	Get(ctx context.Context, r *request.TestGet) (*response.TestGet, error)
}

type test struct {
}

func NewTest() Test {
	return &test{}
}

func (t *test) Get(ctx context.Context, r *request.TestGet) (*response.TestGet, error) {
	return response.NewTestGet(r.Test), nil
}
