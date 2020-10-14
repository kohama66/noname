package handler

import (
	"net/http"

	"github.com/myapp/noname/api/application/usecase"
	"github.com/myapp/noname/api/infrastructure/request"

	"github.com/myapp/noname/api/pkg/log"
	"github.com/myapp/noname/api/presentation/v1/resource/factory"
)

// Reservation DIInterface
type Reservation interface {
	Create(w http.ResponseWriter, hr *http.Request)
	FindByBeautician(w http.ResponseWriter, hr *http.Request)
	Find(w http.ResponseWriter, hr *http.Request)
}

type reservation struct {
	reservationUsecase usecase.Reservation
}

// NewReservation DI初期化関数
func NewReservation(
	reservationUsecase usecase.Reservation,
) Reservation {
	return &reservation{
		reservationUsecase: reservationUsecase,
	}
}

// Create
// @Summary 予約登録
// @Accept  json
// @Produce  json
// @Param data body requestmodel.ReservationCreate true "Request body"
// @Success 200 {object} responsemodel.ReservationCreate
// @Failure 500 {object} resource.Error "Something went wrong"
// @Router /api/v1/reservation [post]
func (r reservation) Create(w http.ResponseWriter, hr *http.Request) {
	req, err := request.NewReservationCreate(hr)
	if err != nil {
		log.Warningf(hr.Context(), "ReservationCreate.Request %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := r.reservationUsecase.Create(hr.Context(), req)
	if err != nil {
		log.Errorf(hr.Context(), "ReservationCreate: %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	factory.JSON(w, res)
	return
}

// FindByBeautician
// @Summary 美容師予約情報取得
// @Accept  json
// @Produce  json
// @Param data body requestmodel.ReservationFindByBeautician true "Request body"
// @Success 200 {object} responsemodel.ReservationFindByBeautician
// @Failure 500 {object} resource.Error "Something went wrong"
// @Router /api/v1/reservation/beautician [Get]
func (r reservation) FindByBeautician(w http.ResponseWriter, hr *http.Request) {
	req, err := request.NewReservationFindByBeautician(hr)
	if err != nil {
		log.Warningf(hr.Context(), "ReservationFindByBeautician.Request %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := r.reservationUsecase.FindByBeautician(hr.Context(), req)
	if err != nil {
		log.Errorf(hr.Context(), "ReservationFindByBeautician: %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	factory.JSON(w, res)
	return
}

// Find
// @Summary 予約検索
// @Accept  json
// @Produce  json
// @Param data body requestmodel.ReservationFind true "Request body"
// @Success 200 {object} responsemodel.ReservationFind
// @Failure 500 {object} resource.Error "Something went wrong"
// @Router /api/v1/reservation/find [Get]
func (r reservation) Find(w http.ResponseWriter, hr *http.Request) {
	req, err := request.NewReservationFind(hr)
	if err != nil {
		log.Warningf(hr.Context(), "ReservationFind.Request %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := r.reservationUsecase.Find(hr.Context(), req)
	if err != nil {
		log.Errorf(hr.Context(), "ReservationFind: %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	factory.JSON(w, res)
	return
}
