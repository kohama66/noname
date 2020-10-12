package usecase

import (
	"context"

	"github.com/myapp/noname/api/application/usecase/requestmodel"
)

// Salon DIInterface
type Salon interface{}
type salon struct{}

// NewSalon DI初期化関数
func NewSalon() Salon {
	return &salon{}
}

func (s *salon) Find(ctx context.Context, r *requestmodel.SalonFind) {

}
