package request

import (
	"encoding/json"
	"net/http"

	"github.com/myapp/noname/api/application/usecase/requestmodel"
)

// NewSalonFind 美容院検索リクエスト関数
func NewSalonFind(req *http.Request) (*requestmodel.SalonFind, error) {
	r := &requestmodel.SalonFind{}
	err := json.NewDecoder(req.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
