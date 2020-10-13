package request

import (
	"net/http"

	"github.com/gorilla/schema"
	"github.com/myapp/noname/api/application/usecase/requestmodel"
)

// NewSalonFind 美容院検索リクエスト関数
func NewSalonFind(req *http.Request) (*requestmodel.SalonFind, error) {
	r := &requestmodel.SalonFind{}
	if err := schema.NewDecoder().Decode(r, req.URL.Query()); err != nil {
		return nil, err
	}
	return r, nil
}
