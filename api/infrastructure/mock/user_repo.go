package mock

import (
	"context"

	"github.com/myapp/noname/api/domain/entity"
	"github.com/myapp/noname/api/domain/repository"
)

// NewUserRepo DI初期化
func NewUserRepo() repository.User {
	return &userrepo{}
}

type userrepo struct{}

func (g *userrepo) GetByAuthID(ctx context.Context, authID string) (*entity.User, error) {
	switch authID {
	case "1":
		return &entity.User{
			ID:            1,
			AuthID:        "auth",
			RandID:        "rand",
			FirstName:     "太郎",
			FirstNameKana: "タロウ",
			LastName:      "山田",
			LastNameKana:  "ヤマダ",
			Email:         "test@test.com",
			IsBeautician:  false,
			PhoneNumber:   "09012345678",
		}, nil
	case "2":
		return &entity.User{
			ID:            2,
			AuthID:        "auth",
			RandID:        "rand",
			FirstName:     "太郎",
			FirstNameKana: "タロウ",
			LastName:      "山田",
			LastNameKana:  "ヤマダ",
			Email:         "test@test.com",
			IsBeautician:  true,
			PhoneNumber:   "09012345678",
		}, nil
	}
	return nil, nil
}

func (g *userrepo) Create(ctx context.Context, ent *entity.User) error {
	return nil
}

func (g *userrepo) GetByRandID(ctx context.Context, randID string) (*entity.User, error) {
	return nil, nil
}
