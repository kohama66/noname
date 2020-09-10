package repository

import (
	"context"

	"github.com/myapp/noname/api/domain/entity"
	"github.com/myapp/noname/api/domain/repository"
	"github.com/myapp/noname/api/infrastructure/db"
	"github.com/volatiletech/sqlboiler/boil"
)

type beautician struct {
	Conn *db.Conn
}

// NewBeautician DI初期化関数
func NewBeautician(conn *db.Conn) repository.Beautician {
	return &beautician{Conn: conn}
}

func (b *beautician) Create(ctx context.Context, ent *entity.Beautician) error {
	return ent.Insert(ctx, b.Conn, boil.Infer())
}

func (b *beautician) GetByAuthID(ctx context.Context, authID string) (*entity.Beautician, error) {
	return entity.Beauticians(
		entity.BeauticianWhere.AuthID.EQ(authID),
		entity.BeauticianWhere.DeletedAt.IsNull(),
	).One(ctx, b.Conn)
}
