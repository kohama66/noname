package request_test

import (
	"net/http"
	"testing"

	"github.com/myapp/noname/api/infrastructure/request"
	"github.com/myapp/noname/api/pkg/context"
)

func TestNewUserGet(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "http://example.com", nil)
	ctx := req.Context()
	req = req.WithContext(context.WithAuthID(ctx, "test_token"))
	r := request.NewUserGet(req)
	if r == nil {
		t.Errorf("TestNewUserGet is nil")
	}
	if r.AuthID == "" {
		t.Errorf("TestNewUserGet authID is nil")
	}
}
