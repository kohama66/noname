package response

import (
	"github.com/myapp/noname/api/application/usecase/responsemodel"
	"github.com/myapp/noname/api/domain/entity"
)

// Guest DIInterface
type Guest interface {
	NewGuestGet(ent *entity.Guest) *responsemodel.GuestGet
}

type guest struct{}

// NewGuest DI初期化関数
func NewGuest() Guest {
	return &guest{}
}

// NewGuestResponsemodel レスポンスモデル変換関数
func NewGuestResponsemodel(ent *entity.Guest) *responsemodel.Guest {
	return &responsemodel.Guest{
		RandID:    ent.RandID,
		FirstName: ent.FirstName,
		LastName:  ent.LastName,
		CreatedAt: ent.CreatedAt,
		UpdatedAt: ent.CreatedAt,
	}
}

func (g *guest) NewGuestGet(ent *entity.Guest) *responsemodel.GuestGet {
	return &responsemodel.GuestGet{
		Guest: NewGuestResponsemodel(ent),
	}
}
