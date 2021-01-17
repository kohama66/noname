package usecase

import (
	"context"
	"fmt"

	"github.com/myapp/noname/api/application/usecase/requestmodel"
	"github.com/myapp/noname/api/application/usecase/responsemodel"
	"github.com/myapp/noname/api/domain/entity"
	"github.com/myapp/noname/api/domain/repository"
	"github.com/myapp/noname/api/infrastructure/response"
)

// Menu DIinterface
type Menu interface {
	Find(ctx context.Context, r *requestmodel.MenuFind) (*responsemodel.MenuFind, error)
	FindByBeauticianWithMenuRandIDs(ctx context.Context, r *requestmodel.MenuFindByBeauticianWithMenuRandIDs) (*responsemodel.MenuFindByBeauticianWithMenuRandIDs, error)
	CreateToBeautician(ctx context.Context, r *requestmodel.BeauticianMenuCreate) (*responsemodel.BeauticianMenuCreate, error)
	DeleteToBeautician(ctx context.Context, r *requestmodel.BeauticianMenuDelete) error
}

type menu struct {
	menuRepository repository.Menu
	menuResponse   response.Menu
	userRepository repository.User
}

// NewMenu Di初期化関数
func NewMenu(
	menuRepository repository.Menu,
	menuResponse response.Menu,
	userRepository repository.User,
) Menu {
	return &menu{
		menuRepository: menuRepository,
		menuResponse:   menuResponse,
		userRepository: userRepository,
	}
}

func (m *menu) Find(ctx context.Context, r *requestmodel.MenuFind) (*responsemodel.MenuFind, error) {
	var menus []*entity.Menu
	if r.BeauticianRandID != "" {
		bt, err := m.userRepository.GetByRandID(ctx, r.BeauticianRandID)
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

func (m *menu) FindByBeauticianWithMenuRandIDs(ctx context.Context, r *requestmodel.MenuFindByBeauticianWithMenuRandIDs) (*responsemodel.MenuFindByBeauticianWithMenuRandIDs, error) {
	bt, err := m.userRepository.GetByRandID(ctx, r.BeauticianRandID)
	if err != nil {
		return nil, err
	}
	mns, err := m.menuRepository.FindByBeauticianWithMenuRandIDs(ctx, bt.ID, r.MenuRandIDs)
	if err != nil {
		return nil, err
	}
	return m.menuResponse.NewFindByBeauticianWithMenuRandIDs(mns), nil
}

func (m *menu) CreateToBeautician(ctx context.Context, r *requestmodel.BeauticianMenuCreate) (*responsemodel.BeauticianMenuCreate, error) {
	bt, err := m.userRepository.GetByAuthID(ctx, r.AuthID)
	if err != nil {
		return nil, err
	}
	mn, err := m.menuRepository.GetByRandID(ctx, r.MenuCategory)
	if err != nil {
		return nil, err
	}
	ent := r.NewBeauticianMenu(bt.ID, mn.ID)
	if err := m.menuRepository.CreateBeauticianMenu(ctx, ent); err != nil {
		return nil, err
	}
	return m.menuResponse.NewBeauticianMenuCreate(ent), nil
}

func (m *menu) DeleteToBeautician(ctx context.Context, r *requestmodel.BeauticianMenuDelete) error {
	me, err := m.userRepository.GetByAuthID(ctx, r.AuthID)
	if err != nil {
		return err
	}
	mn, err := m.menuRepository.GetBeauticianMenuByRandID(ctx, r.RandID)
	if err != nil {
		return err
	}
	if me.ID != mn.BeauticianID {
		return fmt.Errorf("Not yours menu")
	}
	if _, err := m.menuRepository.DeleteBeauticianMenu(ctx, mn); err != nil {
		return err
	}
	return nil
}
