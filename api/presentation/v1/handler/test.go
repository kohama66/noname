package handler

import (
	"net/http"

	"github.com/myapp/noname/api/application/usecase"
	"github.com/myapp/noname/api/application/usecase/request"
	"github.com/myapp/noname/api/pkg/log"
	"github.com/myapp/noname/api/presentation/v1/resource/factory"
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
// @Summary swaggerTest
// @Accept  json
// @Produce  json
// @Param data body request.TestGet true "Request body"
// @Success 200 {object} response.TestGet
// @Failure 500 {object} resource.Error "Soeventthing went wrong"
// @Router /api/v1/hello [get]
func (t test) Get(w http.ResponseWriter, r *http.Request) {
	req, err := request.NewTestGet(r)
	if err != nil {
		log.Warningf(r.Context(), "TestGet.Request %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := t.testUsecase.Get(r.Context(), req)
	if err != nil {
		log.Errorf(r.Context(), "TestGet: %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	factory.JSON(w, res)
	return
}
