package repository

import (
	"context"

	"github.com/myapp/noname/api/domain/entity"
	"github.com/myapp/noname/api/domain/repository"
	"github.com/myapp/noname/api/infrastructure/db"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type menu struct {
	Conn *db.Conn
}

// NewMenu DI初期化関数
func NewMenu(conn *db.Conn) repository.Menu {
	return &menu{
		Conn: conn,
	}
}

func (m *menu) GetByRandID(ctx context.Context, randID string) (*entity.Menu, error) {
	return entity.Menus(
		entity.MenuWhere.RandID.EQ(randID),
		entity.MenuWhere.DeletedAt.IsNull(),
	).One(ctx, m.Conn)
}

func (m *menu) GetAll(ctx context.Context) (entity.MenuSlice, error) {
	return entity.Menus(
		entity.MenuWhere.DeletedAt.IsNull(),
	).All(ctx, m.Conn)
}

func (m *menu) FindByBeautician(ctx context.Context, beauticianID int64) (entity.MenuSlice, error) {
	return entity.Menus(
		qm.InnerJoin("beautician_menus ON menus.id = beautician_menus.menu_id"),
		qm.InnerJoin("beauticians ON beauticians.id = beautician_menus.beautician_id"),
		entity.BeauticianWhere.ID.EQ(beauticianID),
		entity.MenuWhere.DeletedAt.IsNull(),
	).All(ctx, m.Conn)
}

func (m *menu) FindByRandID(ctx context.Context, randIDs []string) (entity.MenuSlice, error) {
	convertedRandIDs := make([]interface{}, len(randIDs))
	for i, v := range randIDs {
		convertedRandIDs[i] = v
	}
	return entity.Menus(
		qm.WhereIn("rand_id IN ?", convertedRandIDs...),
		entity.MenuWhere.DeletedAt.IsNull(),
	).All(ctx, m.Conn)
}

func (m *menu) FindByBeauticianWithMenuRandIDs(ctx context.Context, beauticianID int64, menuIDs []string) (entity.BeauticianMenuSlice, error) {
	return entity.BeauticianMenus(
		qm.InnerJoin("menus ON menus.id = beautician_menus.menu_id"),
		entity.MenuWhere.RandID.IN(menuIDs),
		entity.BeauticianMenuWhere.BeauticianID.EQ(beauticianID),
		entity.BeauticianMenuWhere.DeletedAt.IsNull(),
	).All(ctx, m.Conn)
}
