package repository

import (
	"context"
	"time"

	"github.com/myapp/noname/api/domain/entity"
	"github.com/myapp/noname/api/domain/repository"
	"github.com/myapp/noname/api/infrastructure/db"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
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

func (b *beautician) GetByRandID(ctx context.Context, randID string) (*entity.Beautician, error) {
	return entity.Beauticians(
		entity.BeauticianWhere.RandID.EQ(randID),
		entity.BeauticianWhere.DeletedAt.IsNull(),
	).One(ctx, b.Conn)
}

func (b *beautician) GetAll(ctx context.Context) ([]*entity.Beautician, error) {
	return entity.Beauticians(
		entity.BeauticianWhere.DeletedAt.IsNull(),
	).All(ctx, b.Conn)
}

func (b *beautician) Find(ctx context.Context, date time.Time, menu, salon *int64) ([]*entity.Beautician, error) {
	qms := []qm.QueryMod{}
	if menu != nil {
		qms = append(qms, entity.MenuWhere.ID.EQ(*menu), qm.InnerJoin("beautician_menus ON beautician_menus.beautician_id = beauticians.id"),
			qm.InnerJoin("menus ON menus.id = beautician_menus.menu_id"))
	}
	if salon != nil {
		qms = append(qms, entity.SalonWhere.ID.EQ(*salon), qm.InnerJoin("beautician_salons ON beautician_salons.beautician_id = beauticians.id"),
			qm.InnerJoin("salons ON salons.id = beautician_salons.salon_id"))
	}
	if !date.IsZero() {
		qms = append(qms, qm.Load(entity.BeauticianRels.Reservations, entity.ReservationWhere.Date.EQ(date)))
	}
	if len(qms) == 0 {
		return b.GetAll(ctx)
	}
	qms = append(qms, entity.BeauticianWhere.DeletedAt.IsNull())
	ents, err := entity.Beauticians(
		qms...,
	).All(ctx, b.Conn)
	if err != nil {
		return nil, err
	}
	var beautician []*entity.Beautician
	for _, ent := range ents {
		check := true
		if ent.R != nil {
			for _, r := range ent.R.Reservations {
				if r.Date == date {
					check = false
				}
			}
		}
		if check {
			beautician = append(beautician, ent)
		}
	}
	return beautician, nil
}

func (b *beautician) FindPossibleSalon(ctx context.Context, beauciaonID int64) (entity.BeauticianSalonSlice, error) {
	return entity.BeauticianSalons(
		entity.BeauticianSalonWhere.BeauticianID.EQ(beauciaonID),
		entity.BeauticianSalonWhere.DeletedAt.IsNull(),
	).All(ctx, b.Conn)
}
