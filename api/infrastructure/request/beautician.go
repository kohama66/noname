package request

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/schema"
	"github.com/myapp/noname/api/application/usecase/requestmodel"
	"github.com/myapp/noname/api/pkg/context"
)

// NewBeauticianCreate 美容師登録request関数
func NewBeauticianCreate(req *http.Request) (*requestmodel.BeauticianCreate, error) {
	r := &requestmodel.BeauticianCreate{}
	r.AuthID = context.AuthID(req.Context())
	err := json.NewDecoder(req.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// NewBeauticianGet 美容師情報取得request関数
// func NewBeauticianGet(req *http.Request) (*requestmodel.BeauticianGet, error) {
// 	r := &requestmodel.BeauticianGet{}
// 	r.AuthID = context.AuthID(req.Context())
// 	return r, nil
// }

// NewBeauticianFind 美容師検索request関数
func NewBeauticianFind(req *http.Request) (*requestmodel.BeauticianFind, error) {
	r := &requestmodel.BeauticianFind{}
	if err := schema.NewDecoder().Decode(r, req.URL.Query()); err != nil {
		return nil, err
	}
	return r, nil
}

// NewBeauticianGet 美容師情報取得request関数
func NewBeauticianGet(req *http.Request) (*requestmodel.BeauticianGet, error) {
	r := &requestmodel.BeauticianGet{}
	RandID := chi.URLParam(req, "randID")
	r.RandID = RandID
	return r, nil
}
