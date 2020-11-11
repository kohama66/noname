package request

import (
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
