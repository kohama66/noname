package usecase

import (
	"context"

	"github.com/myapp/noname/api/application/usecase/requestmodel"
	"github.com/myapp/noname/api/application/usecase/responsemodel"
	"github.com/myapp/noname/api/domain/entity"
	"github.com/myapp/noname/api/domain/repository"
	"github.com/myapp/noname/api/infrastructure/response"
)

// Menu DIinterface
type Menu interface {
	Find(ctx context.Context, r *requestmodel.MenuFind) (*responsemodel.MenuFind, error)
}

type menu struct {
	menuRepository       repository.Menu
	menuResponse         response.Menu
	beauticianRepository repository.Beautician
}

// NewMenu Di初期化関数
func NewMenu(
	menuRepository repository.Menu,
	menuResponse response.Menu,
	beauticianRepository repository.Beautician,
) Menu {
	return &menu{
		menuRepository:       menuRepository,
		menuResponse:         menuResponse,
		beauticianRepository: beauticianRepository,
	}
}

func (m *menu) Find(ctx context.Context, r *requestmodel.MenuFind) (*responsemodel.MenuFind, error) {
	var menus []*entity.Menu
	if r.BeauticianRandID != "" {
		bt, err := m.beauticianRepository.GetByRandID(ctx, r.BeauticianRandID)
		if err != nil {
			return nil, err
		}
		mn, err := m.menuRepository.FindByBeautician(ctx, bt.ID)
		if err != nil {
			return nil, err
		}
		menus = mn
	} else {
		mn, err := m.menuRepository.GetAll(ctx)
		if err != nil {
			return nil, err
		}
		menus = mn
	}
	return m.menuResponse.NewMenuFind(menus), nil
}
