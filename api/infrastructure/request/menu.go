package request

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/schema"
	"github.com/myapp/noname/api/application/usecase/requestmodel"
	"github.com/myapp/noname/api/pkg/context"
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

// NewBeauticianMenuCreate 美容師のメニュー作成
func NewBeauticianMenuCreate(req *http.Request) (*requestmodel.BeauticianMenuCreate, error) {
	r := &requestmodel.BeauticianMenuCreate{}
	r.AuthID = context.AuthID(req.Context())
	err := json.NewDecoder(req.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	if err := validate.Struct(r); err != nil {
		return nil, err
	}
	return r, nil
}
