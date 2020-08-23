package handler

import (
	"net/http"

	"github.com/myapp/noname/api/application/usecase"
	"github.com/myapp/noname/api/application/usecase/request"
	"github.com/myapp/noname/api/presentation/v1/resource/factory"
)

type Beautician interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type beautician struct {
	beauticianUsecase usecase.Beautician
}

func NewBeautician(
	beauticianUsecase usecase.Beautician,
) Beautician {
	return beautician{
		beauticianUsecase: beauticianUsecase,
	}
}

// Create
// @Summary 美容師登録
// @Accept  json
// @Produce  json
// @Param data body request.BeauticianCreate true "Request body"
// @Success 200 {object} response.BeauticianCreate
// @Failure 500 {object} resource.Error "Soeventthing went wrong"
// @Router /api/v1/beautician [post]
func (b beautician) Create(w http.ResponseWriter, r *http.Request) {
	req, err := request.NewBeauticianCreate(r)
	if err != nil {
		return
	}
	res, err := b.beauticianUsecase.Create(r.Context(), req)
	if err != nil {
		return
	}
	factory.JSON(w, res)
	return
}
