package request

import (
	"encoding/json"
	"net/http"

	"github.com/myapp/noname/api/application/usecase/requestmodel"
	"github.com/myapp/noname/api/pkg/context"
)

// NewBeauticianCreate 美容師登録リクエスト関数
func NewBeauticianCreate(req *http.Request) (*requestmodel.BeauticianCreate, error) {
	r := &requestmodel.BeauticianCreate{}
	r.AuthID = context.AuthID(req.Context())
	err := json.NewDecoder(req.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// NewBeauticianGet 美容師情報取得リクエスト関数
func NewBeauticianGet(req *http.Request) (*requestmodel.BeauticianGet, error) {
	r := &requestmodel.BeauticianGet{}
	r.ID = 1
	// err := json.NewDecoder(req.Body).Decode(r)
	// if err != nil {
	// 	return nil, err
	// }
	return r, nil
}
