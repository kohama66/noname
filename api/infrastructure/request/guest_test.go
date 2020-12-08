package request_test

import (
	"net/http"
	"testing"

	"github.com/myapp/noname/api/infrastructure/request"
	"github.com/myapp/noname/api/pkg/context"
)

func TestNewGuestGet(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "http://example.com", nil)
	ctx := req.Context()
	req = req.WithContext(context.WithAuthID(ctx, "test_token"))
	r := request.NewGuestGet(req)
	if r == nil {
		t.Errorf("TestNewGuestGet is nil")
	}
	if r.AuthID == "" {
		t.Errorf("TestNewGuestGet authID is nil")
	}
}
