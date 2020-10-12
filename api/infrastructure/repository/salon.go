package repository

import (
	"context"

	"github.com/myapp/noname/api/domain/entity"
	"github.com/myapp/noname/api/domain/repository"
	"github.com/myapp/noname/api/infrastructure/db"
)

type salon struct {
	Conn *db.Conn
}

// NewSalon DI初期化関数
func NewSalon(conn *db.Conn) repository.Salon {
	return &salon{
		Conn: conn,
	}
}

func (m *salon) GetByRandID(ctx context.Context, randID string) (*entity.Salon, error) {
	return entity.Salons(
		entity.SalonWhere.RandID.EQ(randID),
		entity.SalonWhere.DeletedAt.IsNull(),
	).One(ctx, m.Conn)
}
