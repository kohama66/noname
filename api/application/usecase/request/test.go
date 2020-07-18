package request

import (
	"encoding/json"
	"net/http"
)

type TestGet struct {
	Test string `json:"test"`
}

func NewTestGet(req *http.Request) (*TestGet, error) {
	r := &TestGet{}
	err := json.NewDecoder(req.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
