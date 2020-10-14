package request

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
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

// NewReservationFindByBeautician 美容師予約情報取得request関数
func NewReservationFindByBeautician(req *http.Request) (*requestmodel.ReservationFindByBeautician, error) {
	r := &requestmodel.ReservationFindByBeautician{}
	r.AuthID = context.AuthID(req.Context())
	if err := schema.NewDecoder().Decode(r, req.URL.Query()); err != nil {
		return nil, err
	}
	return r, nil
}

// NewReservationFind 予約検索request関数
func NewReservationFind(req *http.Request) (*requestmodel.ReservationFind, error) {
	r := &requestmodel.ReservationFind{}
	if err := schema.NewDecoder().Decode(r, req.URL.Query()); err != nil {
		return nil, err
	}
	return r, nil
}
