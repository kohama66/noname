package handler

import (
	"net/http"

	"github.com/myapp/noname/api/application/usecase"
	"github.com/myapp/noname/api/infrastructure/request"
	"github.com/myapp/noname/api/pkg/log"
	"github.com/myapp/noname/api/presentation/v1/resource/factory"
)

// User DIInterface
type User interface {
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}
type user struct {
	userUsecase usecase.User
}

// NewUser DI初期化関数
func NewUser(
	userUsecase usecase.User,
) User {
	return &user{
		userUsecase: userUsecase,
	}
}

// Get
// @Summary ゲスト情報取得
// @Accept  json
// @Produce  json
// @Param data body requestmodel.UserGet true "Request body"
// @Success 200 {object} responsemodel.UserGet
// @Failure 500 {object} resource.Error "Something went wrong"
// @Router /api/v1/user [get]
func (g user) Get(w http.ResponseWriter, r *http.Request) {
	req := request.NewUserGet(r)
	res, err := g.userUsecase.Get(r.Context(), req)
	if err != nil {
		log.Errorf(r.Context(), "UserGet: %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	factory.JSON(w, res)
	return
}

// Create
// @Summary ゲスト新規登録
// @Accept  json
// @Produce  json
// @Param data body requestmodel.UserCreate true "Request body"
// @Success 200 {object} responsemodel.UserCreate
// @Failure 500 {object} resource.Error "Something went wrong"
// @Router /api/v1/user [post]
func (g user) Create(w http.ResponseWriter, r *http.Request) {
	req, err := request.NewUserCreate(r)
	if err != nil {
		log.Warningf(r.Context(), "UserCreate.Request %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := g.userUsecase.Create(r.Context(), req)
	if err != nil {
		log.Errorf(r.Context(), "UserCreate: %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	factory.JSON(w, res)
	return
}
