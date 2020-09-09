package handler

import (
	"net/http"

	"github.com/myapp/noname/api/application/usecase"
	"github.com/myapp/noname/api/infrastructure/request"
	"github.com/myapp/noname/api/pkg/log"
	"github.com/myapp/noname/api/presentation/v1/resource/factory"
)

type Beautician interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type beautician struct {
	beauticianUsecase usecase.Beautician
}

func NewBeautician(
	beauticianUsecase usecase.Beautician,
) Beautician {
	return &beautician{
		beauticianUsecase: beauticianUsecase,
	}
}

// Create
// @Summary 美容師登録
// @Accept  json
// @Produce  json
// @Param data body requestmodel.BeauticianCreate true "Request body"
// @Success 200 {object} responsemodel.BeauticianCreate
// @Failure 500 {object} resource.Error "Something went wrong"
// @Router /api/v1/beautician [post]
func (b beautician) Create(w http.ResponseWriter, r *http.Request) {
	req, err := request.NewBeauticianCreate(r)
	if err != nil {
		log.Warningf(r.Context(), "BeauticianCreate.Request %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := b.beauticianUsecase.Create(r.Context(), req)
	if err != nil {
		log.Errorf(r.Context(), "BeauticianCreate: %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	factory.JSON(w, res)
	return
}

// Create
// @Summary 美容師情報取得
// @Accept  json
// @Produce  json
// @Param data body requestmodel.BeauticianGet true "Request body"
// @Success 200 {object} responsemodel.BeauticianGet
// @Failure 500 {object} resource.Error "Something went wrong"
// @Router /api/v1/beautician [get]
func (b beautician) Get(w http.ResponseWriter, r *http.Request) {
	req, err := request.NewBeauticianGet(r)
	if err != nil {
		log.Warningf(r.Context(), "BeauticianGet.Request %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := b.beauticianUsecase.Get(r.Context(), req)
	if err != nil {
		log.Errorf(r.Context(), "BeauticianGet: %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	factory.JSON(w, res)
	return
}
