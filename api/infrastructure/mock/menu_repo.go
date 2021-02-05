package mock

import (
	"context"

	"github.com/myapp/noname/api/domain/entity"
	"github.com/myapp/noname/api/domain/repository"
)

type menurepo struct{}

// NewMenuRepo DI初期化関数
func NewMenuRepo() repository.Menu {
	return &menurepo{}
}

func (m *menurepo) GetByRandID(ctx context.Context, randID string) (*entity.Menu, error) {
	return nil, nil
}

func (m *menurepo) GetAll(ctx context.Context) (entity.MenuSlice, error) {
	return nil, nil
}

func (m *menurepo) FindByBeautician(ctx context.Context, beauticianID int64) (entity.MenuSlice, error) {
	return nil, nil
}

func (m *menurepo) FindByRandID(ctx context.Context, randIDs []string) (entity.MenuSlice, error) {
	return nil, nil
}

func (m *menurepo) FindByBeauticianWithMenuRandIDs(ctx context.Context, beauticianID int64, menuIDs []string) (entity.BeauticianMenuSlice, error) {
	return nil, nil
}

func (m *menurepo) ExistsByBeauticianIDWithMenuIDs(ctx context.Context, beauticianID int64, menuIDs []int64) (bool, error) {
	return true, nil
}

func (m *menurepo) GetBeauticianMenusByReservationID(ctx context.Context, reservationID int64) (entity.BeauticianMenuSlice, error) {
	return nil, nil
}

func (m *menurepo) CreateBeauticianMenu(ctx context.Context, ent *entity.BeauticianMenu) error {
	return nil
}

func (m *menurepo) DeleteBeauticianMenu(ctx context.Context, ent *entity.BeauticianMenu) (int64, error) {
	return 0, nil
}

func (m *menurepo) GetBeauticianMenuByRandID(ctx context.Context, randID string) (*entity.BeauticianMenu, error) {
	return &entity.BeauticianMenu{
		ID:           1,
		RandID:       "rand",
		Name:         "カット",
		Price:        1000,
		BeauticianID: 1,
		MenuID:       1,
	}, nil
}
