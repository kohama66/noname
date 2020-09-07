package repository

import (
	"context"

	"github.com/myapp/noname/api/domain/entity"
	"github.com/myapp/noname/api/domain/repository"
	"github.com/myapp/noname/api/infrastructure/db"
	"github.com/volatiletech/null"
)

type reservation struct {
	Conn *db.Conn
}

// NewReservation DI初期化関数
func NewReservation() repository.Reservation {
	return &reservation{}
}

// func (r *reservation) GetByDate(ctx context.Context, date null.Time, time null.String) (*entity.Reservation, error) {
// 	return entity.Reservations(
// 		entity.ReservationWhere.Date.EQ(date),
// 		entity.ReservationWhere.Time.EQ(time),
// 		entity.ReservationWhere.DeletedAt.IsNull(),
// 	).One(ctx, r.Conn)
// }

func (r *reservation) ExistsSpaceDoubleBooking(ctx context.Context, date null.Time, time null.String, spaceID int64) (bool, error) {
	return entity.Reservations(
		entity.ReservationWhere.Date.EQ(date),
		entity.ReservationWhere.Time.EQ(time),
		entity.ReservationWhere.SpaceID.EQ(spaceID),
		entity.ReservationWhere.DeletedAt.IsNull(),
	).Exists(ctx, r.Conn)
}

func (r *reservation) ExistsBeauticianDoubleBooking(ctx context.Context, date null.Time, time null.String, beauticianID int64) (bool, error) {
	return entity.Reservations(
		entity.ReservationWhere.Date.EQ(date),
		entity.ReservationWhere.Time.EQ(time),
		entity.ReservationWhere.BeauticianID.EQ(beauticianID),
		entity.ReservationWhere.DeletedAt.IsNull(),
	).Exists(ctx, r.Conn)
}
