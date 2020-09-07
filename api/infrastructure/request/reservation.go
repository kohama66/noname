package request

import (
	"encoding/json"
	"net/http"

	"github.com/myapp/noname/api/application/usecase/requestmodel"
	"github.com/myapp/noname/api/pkg/context"
)

// NewReservationCreate 予約登録Request関数
func NewReservationCreate(req *http.Request) (*requestmodel.ReservationCreate, error) {
	r := &requestmodel.ReservationCreate{}
	r.AuthID = context.AuthID(req.Context())
	err := json.NewDecoder(req.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	return r, err
}
