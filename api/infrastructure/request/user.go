package request

import (
	"encoding/json"
	"net/http"

	"github.com/myapp/noname/api/application/usecase/requestmodel"
	"github.com/myapp/noname/api/pkg/context"
)

// NewUserGet ゲスト情報取得
func NewUserGet(req *http.Request) *requestmodel.UserGet {
	r := &requestmodel.UserGet{}
	r.AuthID = context.AuthID(req.Context())
	return r
}

// NewUserCreate ゲスト新規登録リクエスト関数
func NewUserCreate(req *http.Request) (*requestmodel.UserCreate, error) {
	r := &requestmodel.UserCreate{}
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
