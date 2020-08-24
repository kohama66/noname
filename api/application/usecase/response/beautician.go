package response

import (
	"time"

	"github.com/myapp/noname/api/domain/entity"
)

// Beautician レスポンス
type Beautician struct {
	ID          int64     `json:"id"`
	RandID      string    `json:"randId"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Age         int64     `json:"age"`
	PhoneNumber string    `json:"phoneNumber"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// NewBeautician エンティティーをレスポンスへ変換する
func NewBeautician(ent *entity.Beautician) *Beautician {
	return &Beautician{
		ID:          ent.ID,
		RandID:      ent.RandID,
		FirstName:   ent.FirstName,
		LastName:    ent.LastName,
		Age:         ent.Age,
		PhoneNumber: ent.PhoneNumber,
		CreatedAt:   ent.CreatedAt,
		UpdatedAt:   ent.UpdatedAt,
	}
}

// BeauticianCreate レスポンス
type BeauticianCreate struct {
	Beautician *Beautician `json:"beautician"`
}

// NewBeauticianCreate レスポンス変換
func NewBeauticianCreate(beautician *entity.Beautician) *BeauticianCreate {
	return &BeauticianCreate{
		Beautician: NewBeautician(beautician),
	}
}
