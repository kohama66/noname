package repository

import (
	"context"

	"github.com/myapp/noname/api/domain/entity"
	"github.com/myapp/noname/api/domain/repository"
	"github.com/myapp/noname/api/infrastructure/db"
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
