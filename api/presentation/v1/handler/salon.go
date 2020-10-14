package handler

import (
	"net/http"

	"github.com/myapp/noname/api/application/usecase"
	"github.com/myapp/noname/api/infrastructure/request"
	"github.com/myapp/noname/api/pkg/log"
	"github.com/myapp/noname/api/presentation/v1/resource/factory"
)

type salon struct {
	salonUsecase usecase.Salon
}

// Salon DIInterface
type Salon interface {
	Find(w http.ResponseWriter, r *http.Request)
}

// NewSalon DI初期化関数
func NewSalon(
	salonUsecase usecase.Salon,
) Salon {
	return salon{
		salonUsecase: salonUsecase,
	}
}

// Find
// @Summary 美容院検索
// @Accept  json
// @Produce  json
// @Param data body requestmodel.SalonFind true "Request body"
// @Success 200 {object} responsemodel.SalonFind
// @Failure 500 {object} resource.Error "Something went wrong"
// @Router /api/v1/salon/find [get]
func (s salon) Find(w http.ResponseWriter, r *http.Request) {
	req, err := request.NewSalonFind(r)
	if err != nil {
		log.Warningf(r.Context(), "SalonFind.Request %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := s.salonUsecase.Find(r.Context(), req)
	if err != nil {
		log.Errorf(r.Context(), "SalonFind: %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	factory.JSON(w, res)
	return
}
