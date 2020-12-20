package request

import (
	"encoding/json"
	"net/http"

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
	if err := validate.Struct(r); err != nil {
		return nil, err
	}
	return r, nil
}

// NewBeauticianFind 美容師検索request関数
func NewBeauticianFind(req *http.Request) (*requestmodel.BeauticianFind, error) {
	r := &requestmodel.BeauticianFind{}
	if err := schema.NewDecoder().Decode(r, req.URL.Query()); err != nil {
		return nil, err
	}
	return r, nil
}

// NewBeauticianGet 美容師情報取得request関数
func NewBeauticianGet(req *http.Request) *requestmodel.BeauticianGet {
	r := &requestmodel.BeauticianGet{}
	r.AuthID = context.AuthID(req.Context())
	return r
}

// NewBeauticianUpdate 美容師情報更新request関数
func NewBeauticianUpdate(req *http.Request) (*requestmodel.BeauticianUpdate, error) {
	r := &requestmodel.BeauticianUpdate{}
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
