package handler

import (
	"net/http"

	"github.com/myapp/noname/api/application/usecase"
	"github.com/myapp/noname/api/infrastructure/request"
	"github.com/myapp/noname/api/pkg/log"
	"github.com/myapp/noname/api/presentation/v1/resource/factory"
)

// Menu DIInterface
type Menu interface {
	Find(w http.ResponseWriter, r *http.Request)
	FindByBeauticianWithMenuRandID(w http.ResponseWriter, r *http.Request)
}

type menu struct {
	menuUsecase usecase.Menu
}

// NewMenu DI初期化関数
func NewMenu(
	menuUsecase usecase.Menu,
) Menu {
	return &menu{
		menuUsecase: menuUsecase,
	}
}

// Find
// @Summary メニュー検索
// @Accept  json
// @Produce  json
// @Param data body requestmodel.MenuFind true "Request body"
// @Success 200 {object} responsemodel.MenuFind
// @Failure 500 {object} resource.Error "Something went wrong"
// @Router /api/v1/menu/find [get]
func (m menu) Find(w http.ResponseWriter, r *http.Request) {
	req, err := request.NewMenuFind(r)
	if err != nil {
		log.Warningf(r.Context(), "MenuFind.Request %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := m.menuUsecase.Find(r.Context(), req)
	if err != nil {
		log.Errorf(r.Context(), "MenuFind: %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	factory.JSON(w, res)
	return
}

// FindByBeauticianWithMenuRandID
// @Summary 美容師の詳細メニュー取得
// @Accept  json
// @Produce  json
// @Param data body requestmodel.MenuFindByBeauticianWithMenuRandIDs true "Request body"
// @Success 200 {object} responsemodel.MenuFindByBeauticianWithMenuRandIDs
// @Failure 500 {object} resource.Error "Something went wrong"
// @Router /api/v1/menu/find/{beauticianRandID} [get]
func (m menu) FindByBeauticianWithMenuRandID(w http.ResponseWriter, r *http.Request) {
	req, err := request.NewMenuFindByBeauticianWithMenuRandIDs(r)
	if err != nil {
		log.Warningf(r.Context(), "MenuFindByBeauticianWithMenuRandID.Request %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := m.menuUsecase.FindByBeauticianWithMenuRandIDs(r.Context(), req)
	if err != nil {
		log.Errorf(r.Context(), "MenuFindByBeauticianWithMenuRandID: %v", err)
		factory.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	factory.JSON(w, res)
	return
}
