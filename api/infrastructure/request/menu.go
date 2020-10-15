package request

import (
	"net/http"

	"github.com/gorilla/schema"
	"github.com/myapp/noname/api/application/usecase/requestmodel"
)

// NewMenuFind メニュー検索
func NewMenuFind(req *http.Request) (*requestmodel.MenuFind, error) {
	r := &requestmodel.MenuFind{}
	if err := schema.NewDecoder().Decode(r, req.URL.Query()); err != nil {
		return nil, err
	}
	return r, nil
}
