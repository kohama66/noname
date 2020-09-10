package handler

import (
	"net/http"

	"github.com/myapp/noname/api/application/usecase"
	"github.com/myapp/noname/api/infrastructure/request"

	"github.com/myapp/noname/api/pkg/log"
	"github.com/myapp/noname/api/presentation/v1/resource/factory"
)

type Reservation interface {
	Create(w http.ResponseWriter, hr *http.Request)
	FindByBeautician(w http.ResponseWriter, hr *http.Request)
}

type reservation struct {
	reservationUsecase usecase.Reservation
}

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
