package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/myapp/noname/api/domain/entity"
	"github.com/myapp/noname/api/domain/repository"
	"github.com/myapp/noname/api/infrastructure/db"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type reservation struct {
	Conn *db.Conn
}

// NewReservation DI初期化関数
func NewReservation(conn *db.Conn) repository.Reservation {
	return &reservation{Conn: conn}
}

func (r *reservation) ExistsSpaceDoubleBooking(ctx context.Context, date null.Time, time null.String, spaceID int64) (bool, error) {
	return entity.Reservations(
		qm.Where("date = ?", date),
		entity.ReservationWhere.Time.EQ(time),
		entity.ReservationWhere.SpaceID.EQ(spaceID),
		entity.ReservationWhere.DeletedAt.IsNull(),
	).Exists(ctx, r.Conn)
}

func (r *reservation) ExistsBeauticianDoubleBooking(ctx context.Context, date null.Time, time null.String, beauticianID int64) (bool, error) {
	return entity.Reservations(
		qm.Where("date = ?", date),
		entity.ReservationWhere.Time.EQ(time),
		entity.ReservationWhere.BeauticianID.EQ(beauticianID),
		entity.ReservationWhere.DeletedAt.IsNull(),
	).Exists(ctx, r.Conn)
}

func (r *reservation) Create(ctx context.Context, ent *entity.Reservation) error {
	return ent.Insert(ctx, r.Conn, boil.Infer())
}

func (r *reservation) FindByBeautician(ctx context.Context, beauticianID int64, date time.Time) (entity.ReservationSlice, error) {
	return entity.Reservations(
		qm.Where("date BETWEEN ? AND ?", fmt.Sprint(date), fmt.Sprint(date.AddDate(0, 0, 13))),
		entity.ReservationWhere.BeauticianID.EQ(beauticianID),
	).All(ctx, r.Conn)
}
