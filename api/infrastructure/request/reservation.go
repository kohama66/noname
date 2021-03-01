package request

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
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

// NewReservationFindByUser ゲスト予約履歴取得関数
func NewReservationFindByUser(req *http.Request) *requestmodel.ReservationFindByUser {
	r := &requestmodel.ReservationFindByUser{}
	r.AuthID = context.AuthID(req.Context())
	return r
}

// NewReservationGetInfo 予約詳細取得
func NewReservationGetInfo(req *http.Request) (*requestmodel.ReservationGetInfo, error) {
	r := &requestmodel.ReservationGetInfo{}
	RandID := chi.URLParam(req, "randID")
	r.RandID = RandID
	return r, nil
}

// NewReservationSetHoliday 美容師休日設定
func NewReservationSetHoliday(req *http.Request) (*requestmodel.ReservationSetHoliday, error) {
	r := &requestmodel.ReservationSetHoliday{}
	r.AuthID = context.AuthID(req.Context())
	err := json.NewDecoder(req.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	fmt.Println(r.Holiday)
	if err := validate.Struct(r); err != nil {
		return nil, err
	}
	return r, err
}

// NewReservationCreateBySalonDayOff 美容師休日設定
func NewReservationCreateBySalonDayOff(req *http.Request) (*requestmodel.ReservationCreateBySalonDayOff, error) {
	r := &requestmodel.ReservationCreateBySalonDayOff{}
	r.AuthID = context.AuthID(req.Context())
	if err := json.NewDecoder(req.Body).Decode(r); err != nil {
		return nil, err
	}
	if err := validate.Struct(r); err != nil {
		return nil, err
	}
	return r, nil
}
