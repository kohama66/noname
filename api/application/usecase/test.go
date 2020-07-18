package usecase

import (
	"github.com/myapp/noname/api/application/usecase/request"
	"github.com/myapp/noname/api/application/usecase/response"
)

type Test interface {
}

type test struct {
}

func NewTest() Test {
	return test{}
}

func (t *test) Get(r *request.TestGet) *response.TestGet {
	return response.NewTestGet(r.Test)
}
