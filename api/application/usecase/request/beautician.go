package request

import "net/http"

type BeauticianCreate struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Age         int64  `json:"age"`
	PhoneNumber string `json:"phoneNumber"`
}

func NewBeauticianCreate(req *http.Request) (*BeauticianCreate, error) {
}
