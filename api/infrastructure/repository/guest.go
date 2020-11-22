package repository

import (
	"context"

	"github.com/myapp/noname/api/domain/entity"
	"github.com/myapp/noname/api/domain/repository"
	"github.com/myapp/noname/api/infrastructure/db"
	"github.com/volatiletech/sqlboiler/boil"
)

type guest struct {
	Conn *db.Conn
}

// NewGuest DI初期化関数
func NewGuest(conn *db.Conn) repository.Guest {
	return &guest{
		Conn: conn,
	}
}

func (g *guest) GetByAuthID(ctx context.Context, authID string) (*entity.Guest, error) {
	return entity.Guests(
		entity.GuestWhere.AuthID.EQ(authID),
		entity.GuestWhere.DeletedAt.IsNull(),
	).One(ctx, g.Conn)
}

func (g *guest) Create(ctx context.Context, ent *entity.Guest) error {
	return ent.Insert(ctx, g.Conn, boil.Infer())
}
