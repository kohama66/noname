package request

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/schema"
	"github.com/myapp/noname/api/application/usecase/requestmodel"
)

// NewMenuFind メニュー検索
func NewMenuFind(req *http.Request) (*requestmodel.MenuFind, error) {
	r := &requestmodel.MenuFind{}
	if err := schema.NewDecoder().Decode(r, req.URL.Query()); err != nil {
		return nil, err
	}
	return r, nil
}

// NewMenuFindByBeauticianWithMenuRandIDs 美容師の詳細メニュー取得
func NewMenuFindByBeauticianWithMenuRandIDs(req *http.Request) (*requestmodel.MenuFindByBeauticianWithMenuRandIDs, error) {
	r := &requestmodel.MenuFindByBeauticianWithMenuRandIDs{}
	if err := schema.NewDecoder().Decode(r, req.URL.Query()); err != nil {
		return nil, err
	}
	BeauticianRandID := chi.URLParam(req, "beauticianRandID")
	r.BeauticianRandID = BeauticianRandID
	return r, nil
}
