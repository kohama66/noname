package repository

import (
	"context"
	"fmt"

	"github.com/myapp/noname/api/domain/entity"
	"github.com/myapp/noname/api/domain/repository"
	"github.com/myapp/noname/api/infrastructure/db"
	"github.com/volatiletech/sqlboiler/boil"
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
		qm.InnerJoin("users ON users.id = beautician_menus.beautician_id"),
		entity.UserWhere.ID.EQ(beauticianID),
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
		qm.Load(entity.BeauticianMenuRels.Menu),
	).All(ctx, m.Conn)
}

func (m *menu) ExistsByBeauticianIDWithMenuIDs(ctx context.Context, beauticianID int64, menuIDs []int64) (bool, error) {
	convertedMenuIDs := make([]interface{}, len(menuIDs))
	for i, v := range menuIDs {
		convertedMenuIDs[i] = v
	}
	return entity.BeauticianMenus(
		entity.BeauticianMenuWhere.BeauticianID.EQ(beauticianID),
		qm.WhereIn("menu_id IN ?", convertedMenuIDs...),
		entity.BeauticianMenuWhere.DeletedAt.IsNull(),
	).Exists(ctx, m.Conn)
}

func (m *menu) GetBeauticianMenusByReservationID(ctx context.Context, reservationID int64) (entity.BeauticianMenuSlice, error) {
	return entity.BeauticianMenus(
		qm.InnerJoin(fmt.Sprintf("%v ON %v = %v", entity.TableNames.ReservationMenus, entity.ReservationMenuColumns.BeauticianMenuID, entity.BeauticianMenuColumns.ID)),
		entity.ReservationMenuWhere.ReservationID.EQ(reservationID),
		entity.BeauticianMenuWhere.DeletedAt.IsNull(),
	).All(ctx, m.Conn)
}

func (m *menu) CreateBeauticianMenu(ctx context.Context, ent *entity.BeauticianMenu) error {
	return ent.Insert(ctx, m.Conn, boil.Infer())
}
