package request

import (
	"encoding/json"
	"net/http"

	"github.com/myapp/noname/api/pkg/context"
)

type BeauticianCreate struct {
	AuthID      string `json:"-"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Age         int64  `json:"age"`
	PhoneNumber string `json:"phoneNumber"`
}

func NewBeauticianCreate(req *http.Request) (*BeauticianCreate, error) {
	r := &BeauticianCreate{}
	r.AuthID = context.AuthID(req.Context())
	err := json.NewDecoder(req.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// func (b BeauticianCreate) NewBeautician() *entity.Beautician {
// 	return &entity.Beautician{
// 		AuthID: b.AuthID,
// 	}
// }
