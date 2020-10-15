package response

import (
	"github.com/myapp/noname/api/application/usecase/responsemodel"
	"github.com/myapp/noname/api/domain/entity"
)

// Menu DIInterface
type Menu interface {
	NewMenuFind(ents []*entity.Menu) *responsemodel.MenuFind
}
type menu struct{}

// NewMenu DI初期化関数
func NewMenu() Menu {
	return &menu{}
}

// NewMenuResponsemodel エンティティ変換関数
func NewMenuResponsemodel(ent *entity.Menu) *responsemodel.Menu {
	return &responsemodel.Menu{
		ID:        ent.ID,
		RandID:    ent.RandID,
		Name:      ent.Name,
		CreatedAt: ent.CreatedAt,
		UpdatedAt: ent.UpdatedAt,
	}
}

func (m *menu) NewMenuFind(ents []*entity.Menu) *responsemodel.MenuFind {
	mn := make([]*responsemodel.Menu, len(ents))
	for i, v := range ents {
		mn[i] = NewMenuResponsemodel(v)
	}
	return &responsemodel.MenuFind{
		Menus: mn,
	}
}
