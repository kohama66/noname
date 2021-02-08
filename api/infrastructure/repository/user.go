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

func (u *user) GetByAuthID(ctx context.Context, authID string) (*entity.User, error) {
	return entity.Users(
		entity.UserWhere.AuthID.EQ(authID),
		entity.UserWhere.DeletedAt.IsNull(),
		qm.Load(entity.UserRels.Beautician),
		qm.Load(entity.UserRels.BeauticianBeauticianMenus),
		qm.Load(entity.UserRels.UserSalons),
	).One(ctx, u.Conn)
}

func (u *user) Create(ctx context.Context, ent *entity.User) error {
	return ent.Insert(ctx, u.Conn, boil.Infer())
}

func (u *user) GetByRandID(ctx context.Context, randID string) (*entity.User, error) {
	return entity.Users(
		entity.UserWhere.RandID.EQ(randID),
		entity.UserWhere.DeletedAt.IsNull(),
		qm.Load(entity.UserRels.Beautician),
		qm.Load(entity.UserRels.BeauticianBeauticianMenus),
	).One(ctx, u.Conn)
}

func (u *user) FindBySalonID(ctx context.Context, salonID int64) (entity.UserSlice, error) {
	return entity.Users(
		qm.InnerJoin("%s ON %s.%s = %s.%s", entity.TableNames.Beauticians, entity.TableNames.Beauticians, entity.BeauticianColumns.UserID, entity.TableNames.Users, entity.UserColumns.ID),
		qm.InnerJoin("%s ON %s.%s = %s.%s", entity.TableNames.BeauticianSalons, entity.TableNames.BeauticianSalons, entity.BeauticianSalonColumns.BeauticianID, entity.TableNames.Beauticians, entity.BeauticianColumns.UserID),
		entity.BeauticianSalonWhere.SalonID.EQ(salonID),
		entity.UserWhere.DeletedAt.IsNull(),
		entity.BeauticianWhere.DeletedAt.IsNull(),
		entity.BeauticianSalonWhere.DeletedAt.IsNull(),
	).All(ctx, u.Conn)
}
