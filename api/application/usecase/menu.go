package usecase

import (
	"context"

	"github.com/myapp/noname/api/application/usecase/requestmodel"
	"github.com/myapp/noname/api/application/usecase/responsemodel"
	"github.com/myapp/noname/api/domain/repository"
	"github.com/myapp/noname/api/infrastructure/response"
)

// Menu DIinterface
type Menu interface {
	Find(ctx context.Context, r *requestmodel.MenuFind) (*responsemodel.MenuFind, error)
}

type menu struct {
	menuRepository repository.Menu
	menuResponse   response.Menu
}

// NewMenu Di初期化関数
func NewMenu(
	menuRepository repository.Menu,
	menuResponse response.Menu,
) Menu {
	return &menu{
		menuRepository: menuRepository,
		menuResponse:   menuResponse,
	}
}

func (m *menu) Find(ctx context.Context, r *requestmodel.MenuFind) (*responsemodel.MenuFind, error) {
	ents, err := m.menuRepository.Find(ctx)
	if err != nil {
		return nil, err
	}
	return m.menuResponse.NewMenuFind(ents), nil
}
