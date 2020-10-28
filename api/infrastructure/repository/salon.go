package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/myapp/noname/api/domain/entity"
	"github.com/myapp/noname/api/domain/repository"
	"github.com/myapp/noname/api/infrastructure/db"
	"github.com/volatiletech/sqlboiler/queries/qm"
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

func (s *salon) GetByRandID(ctx context.Context, randID string) (*entity.Salon, error) {
	return entity.Salons(
		entity.SalonWhere.RandID.EQ(randID),
		entity.SalonWhere.DeletedAt.IsNull(),
	).One(ctx, s.Conn)
}

// func (s *salon) Find(ctx context.Context, beauticianID int64) (entity.SalonSlice, error) {
// 	qms := []qm.QueryMod{}
// 	if beauticianID != 0 {
// 		qms = append(qms, entity.BeauticianWhere.ID.EQ(beauticianID), qm.InnerJoin("beautician_salons ON beautician_salons.salon_id = salons.id"),
// 			qm.InnerJoin("beauticians ON beauticians.id = beautician_salons.beautician_id"))
// 	}
// 	qms = append(qms, entity.SalonWhere.DeletedAt.IsNull())
// 	return entity.Salons(
// 		qms...,
// 	).All(ctx, s.Conn)
// }

func (s *salon) GetAll(ctx context.Context) (entity.SalonSlice, error) {
	return entity.Salons(
		entity.SalonWhere.DeletedAt.IsNull(),
	).All(ctx, s.Conn)
}

func (s *salon) FindByBeautician(ctx context.Context, beauticianID int64) (entity.SalonSlice, error) {
	return entity.Salons(
		qm.InnerJoin("beautician_salons ON salons.id = beautician_salons.salon_id"),
		qm.InnerJoin("beauticians ON beauticians.id = beautician_salons.beautician_id"),
		entity.BeauticianWhere.ID.EQ(beauticianID),
		entity.SalonWhere.DeletedAt.IsNull(),
	).All(ctx, s.Conn)
}

func (s *salon) Find(ctx context.Context, beauticianID int64, date time.Time) (entity.SalonSlice, error) {
	var sls entity.SalonSlice
	fmt.Printf("%s")
	if !date.IsZero() {
		sl, err := entity.Salons(
			qm.InnerJoin(fmt.Sprintf("%s ON %s = %s", entity.TableNames.Spaces, entity.SpaceColumns.SalonID, "salons.id")),
			qm.LeftOuterJoin(fmt.Sprintf("%s ON %s = %s", entity.TableNames.Reservations, entity.ReservationColumns.SpaceID, "spaces.id")),
			entity.ReservationWhere.Date.NEQ(date),
			qm.Or("reservations.date IS NULL"),
			entity.SalonWhere.DeletedAt.IsNull(),
		).All(ctx, s.Conn)
		if err != nil {
			return nil, err
		}
		sls = append(sls, sl...)
	}
	return sls, nil
}
