package handler

import (
	"fmt"
	"net/http"

	"github.com/myapp/noname/api/application/usecase"
)

type Test interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type test struct {
	testUsecase usecase.Test
}

func NewTest(
	testUsecase usecase.Test,
) Test {
	return test{
		testUsecase: testUsecase,
	}
}

// Test
// @Summary swaggerテスト
// @Accept  json
// @Produce  json
// @Param data body request.TestGet true "Request body"
// @Success 200 {object} response.TestGet
// @Failure 500 {object} resource.Error "Soeventthing went wrong"
// @Router /api/v1/hello [get]
func (t test) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ok golang!")
}
