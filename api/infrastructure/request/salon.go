package request

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/schema"
	"github.com/myapp/noname/api/application/usecase/requestmodel"
	"github.com/myapp/noname/api/pkg/context"
)

// NewSalonFind 美容院検索リクエスト関数
func NewSalonFind(req *http.Request) (*requestmodel.SalonFind, error) {
	r := &requestmodel.SalonFind{}
	if err := schema.NewDecoder().Decode(r, req.URL.Query()); err != nil {
		return nil, err
	}
	return r, nil
}

// NewSalonFindNotBelongs 美容師が所属してない美容院検索
func NewSalonFindNotBelongs(req *http.Request) *requestmodel.SalonFindNotBelongs {
	r := &requestmodel.SalonFindNotBelongs{}
	r.AuthID = context.AuthID(req.Context())
	return r
}

// NewBeauticianSalonCreata 美容師の行ける美容院追加
func NewBeauticianSalonCreata(req *http.Request) (*requestmodel.BeauticianSalonCreata, error) {
	r := &requestmodel.BeauticianSalonCreata{}
	r.AuthID = context.AuthID(req.Context())
	if err := json.NewDecoder(req.Body).Decode(r); err != nil {
		return nil, err
	}
	if err := validate.Struct(r); err != nil {
		return nil, err
	}
	return r, nil
}

// NewBeauticianSalonDelete 美容師の行ける美容院削除
func NewBeauticianSalonDelete(req *http.Request) *requestmodel.BeauticianSalonDelete {
	r := &requestmodel.BeauticianSalonDelete{}
	r.AuthID = context.AuthID(req.Context())
	r.SalonRandID = chi.URLParam(req, "randID")
	return r
}

// NewSalonCreate 美容院作成
func NewSalonCreate(req *http.Request) (*requestmodel.SalonCreate, error) {
	r := &requestmodel.SalonCreate{}
	r.AuthID = context.AuthID(req.Context())
	if err := json.NewDecoder(req.Body).Decode(r); err != nil {
		return nil, err
	}
	if err := validate.Struct(r); err != nil {
		return nil, err
	}
	return r, nil
}
