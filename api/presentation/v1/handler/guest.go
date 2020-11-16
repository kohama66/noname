package handler

import (
	"net/http"

	"github.com/myapp/noname/api/application/usecase"
	"github.com/myapp/noname/api/infrastructure/request"
	"github.com/myapp/noname/api/pkg/log"
	"github.com/myapp/noname/api/presentation/v1/resource/factory"
)

// Guest DIInterface
type Guest interface {
	Get(w http.ResponseWriter, r *http.Request)
}
type guest struct {
	guestUsecase usecase.Guest
}

// NewGuest DI初期化関数
func NewGuest(
	guestUsecase usecase.Guest,
) Guest {
	return &guest{
		guestUsecase: guestUsecase,
	}
}

// Get
// @Summary ゲスト情報取得
// @Accept  json
// @Produce  json
// @Param data body requestmodel.GuestGet true "Request body"
// @Success 200 {object} responsemodel.GuestGet
// @Failure 500 {object} resource.Error "Something went wrong"
// @Router /api/v1/guest [get]
func (g guest) Get(w http.ResponseWriter, r *http.Request) {
	req := request.NewGuestGet(r)
	res, err := g.guestUsecase.Get(r.Context(), req)
	if err != nil {
		log.Errorf(r.Context(), "GuestGet: %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	factory.JSON(w, res)
	return
}
