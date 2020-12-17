package repository

import (
	"context"

	"github.com/myapp/noname/api/domain/entity"
	"github.com/myapp/noname/api/domain/repository"
	"github.com/myapp/noname/api/infrastructure/db"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type user struct {
	Conn *db.Conn
}

// NewUser DI初期化関数
func NewUser(conn *db.Conn) repository.User {
	return &user{
		Conn: conn,
	}
}

func (g *user) GetByAuthID(ctx context.Context, authID string) (*entity.User, error) {
	return entity.Users(
		entity.UserWhere.AuthID.EQ(authID),
		entity.UserWhere.DeletedAt.IsNull(),
		qm.Load(entity.UserRels.Beautician),
	).One(ctx, g.Conn)
}

func (g *user) Create(ctx context.Context, ent *entity.User) error {
	return ent.Insert(ctx, g.Conn, boil.Infer())
}

func (g *user) GetByRandID(ctx context.Context, randID string) (*entity.User, error) {
	return entity.Users(
		entity.UserWhere.RandID.EQ(randID),
		entity.UserWhere.DeletedAt.IsNull(),
		qm.Load(entity.UserRels.Beautician),
	).One(ctx, g.Conn)
}
