package repository

import (
	"context"

	"github.com/myapp/noname/api/domain/entity"
	"github.com/myapp/noname/api/domain/repository"
	"github.com/myapp/noname/api/infrastructure/db"
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
