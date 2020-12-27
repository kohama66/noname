package handler

import (
	"net/http"

	"github.com/myapp/noname/api/application/usecase"
	"github.com/myapp/noname/api/infrastructure/request"
	"github.com/myapp/noname/api/pkg/log"
	"github.com/myapp/noname/api/presentation/v1/resource/factory"
)

// Beautician DIInterface
type Beautician interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetByAuthID(w http.ResponseWriter, r *http.Request)
	Find(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type beautician struct {
	beauticianUsecase usecase.Beautician
}

// NewBeautician DI初期化
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

// GetByAuthID
// @Summary 美容師情報取得
// @Accept  json
// @Produce  json
// @Param data body requestmodel.BeauticianGet true "Request body"
// @Success 200 {object} responsemodel.BeauticianGet
// @Failure 500 {object} resource.Error "Something went wrong"
// @Router /api/v1/beautician [get]
func (b beautician) GetByAuthID(w http.ResponseWriter, r *http.Request) {
	req := request.NewBeauticianGet(r)
	res, err := b.beauticianUsecase.Get(r.Context(), req)
	if err != nil {
		log.Errorf(r.Context(), "BeauticianGet: %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	factory.JSON(w, res)
	return
}

// Find
// @Summary 美容師検索
// @Accept  json
// @Produce  json
// @Param data body requestmodel.BeauticianFind true "Request body"
// @Success 200 {object} responsemodel.BeauticianFind
// @Failure 500 {object} resource.Error "Something went wrong"
// @Router /api/v1/beautician/find [get]
func (b beautician) Find(w http.ResponseWriter, r *http.Request) {
	req, err := request.NewBeauticianFind(r)
	if err != nil {
		log.Warningf(r.Context(), "BeauticianFind.Request %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := b.beauticianUsecase.Find(r.Context(), req)
	if err != nil {
		log.Errorf(r.Context(), "BeauticianFind: %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	factory.JSON(w, res)
	return
}

// Update
// @Summary 美容師情報更新
// @Accept  json
// @Produce  json
// @Param data body requestmodel.BeauticianUpdate true "Request body"
// @Success 200
// @Failure 500 {object} resource.Error "Something went wrong"
// @Router /api/v1/beautician [put]
func (b beautician) Update(w http.ResponseWriter, r *http.Request) {
	req, err := request.NewBeauticianUpdate(r)
	if err != nil {
		log.Warningf(r.Context(), "BeauticianUpdate.Request %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = b.beauticianUsecase.Update(r.Context(), req)
	if err != nil {
		log.Errorf(r.Context(), "BeauticianUpdate: %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}
