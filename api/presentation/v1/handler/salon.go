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
	FindNotBelongs(w http.ResponseWriter, r *http.Request)
	CreateBeauticianSalon(w http.ResponseWriter, r *http.Request)
	DeleteBeauticianSalon(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	GetMyPage(w http.ResponseWriter, r *http.Request)
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

// FindNotBelongs
// @Summary 美容院検索
// @Accept  json
// @Produce  json
// @Param data body requestmodel.SalonFindNotBelongs true "Request body"
// @Success 200 {object} responsemodel.SalonFindNotBelongs
// @Failure 500 {object} resource.Error "Something went wrong"
// @Router /api/v1/salon/belongs [get]
func (s salon) FindNotBelongs(w http.ResponseWriter, r *http.Request) {
	req := request.NewSalonFindNotBelongs(r)
	res, err := s.salonUsecase.FindNotBelongs(r.Context(), req)
	if err != nil {
		log.Errorf(r.Context(), "SalonFindNotBelongs: %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	factory.JSON(w, res)
	return
}

// CreateBeauticianSalon
// @Summary 美容師美容院追加
// @Description 美容師が仕事可能な美容院を一つ追加します
// @Accept  json
// @Param data body requestmodel.BeauticianSalonCreata true "Request body"
// @Success 200
// @Failure 500 {object} resource.Error "Something went wrong"
// @Router /api/v1/salon/beautician [post]
func (s salon) CreateBeauticianSalon(w http.ResponseWriter, r *http.Request) {
	req, err := request.NewBeauticianSalonCreata(r)
	if err != nil {
		log.Warningf(r.Context(), "CreateBeauticianSalon.Request %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.salonUsecase.CreateToBeautician(r.Context(), req)
	if err != nil {
		log.Errorf(r.Context(), "CreateBeauticianSalon: %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

// DeleteBeauticianSalon
// @Summary 美容師美容院削除
// @Description 美容師が仕事可能な美容院を一つ削除します
// @Accept  json
// @Param data body requestmodel.BeauticianSalonDelete true "Request body"
// @Success 200
// @Failure 500 {object} resource.Error "Something went wrong"
// @Router /api/v1/salon/beautician/{randID} [delete]
func (s salon) DeleteBeauticianSalon(w http.ResponseWriter, r *http.Request) {
	req := request.NewBeauticianSalonDelete(r)
	err := s.salonUsecase.DeleteToBeautician(r.Context(), req)
	if err != nil {
		log.Errorf(r.Context(), "DeleteBeauticianSalon: %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

// Create
// @Summary 美容院作成
// @Description 美容院を作成します
// @Accept  json
// @Param data body requestmodel.SalonCreate true "Request body"
// @Success 200
// @Failure 500 {object} resource.Error "Something went wrong"
// @Router /api/v1/salon [post]
func (s salon) Create(w http.ResponseWriter, r *http.Request) {
	req, err := request.NewSalonCreate(r)
	if err != nil {
		log.Warningf(r.Context(), "CreateSalon.Request %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := s.salonUsecase.Create(r.Context(), req)
	if err != nil {
		log.Errorf(r.Context(), "CreateSalon: %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	factory.JSON(w, res)
	return
}

// GetMyPage
// @Summary 美容院マイページ取得
// @Accept  json
// @Produce  json
// @Param data body requestmodel.SalonMyPageGet true "Request body"
// @Success 200 {object} responsemodel.SalonMyPageGet
// @Failure 500 {object} resource.Error "Something went wrong"
// @Router /api/v1/salon/mypage/{randID} [get]
func (s salon) GetMyPage(w http.ResponseWriter, r *http.Request) {
	req := request.NewSalonMyPageGet(r)
	res, err := s.salonUsecase.GetMyPage(r.Context(), req)
	if err != nil {
		log.Errorf(r.Context(), "SalonMyPageGet: %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	factory.JSON(w, res)
	return
}
