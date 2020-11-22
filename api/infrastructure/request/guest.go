package request

import (
	"encoding/json"
	"net/http"

	"github.com/myapp/noname/api/application/usecase/requestmodel"
	"github.com/myapp/noname/api/pkg/context"
)

// NewGuestGet ゲスト情報取得
func NewGuestGet(req *http.Request) *requestmodel.GuestGet {
	r := &requestmodel.GuestGet{}
	r.AuthID = context.AuthID(req.Context())
	return r
}

// NewGuestCreate ゲスト新規登録リクエスト関数
func NewGuestCreate(req *http.Request) (*requestmodel.GuestCreate, error) {
	r := &requestmodel.GuestCreate{}
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
