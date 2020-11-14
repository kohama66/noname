package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/myapp/noname/api/domain/entity"
	"github.com/myapp/noname/api/domain/entityx"
	"github.com/myapp/noname/api/domain/repository"
	"github.com/myapp/noname/api/infrastructure/db"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type reservation struct {
	Conn               *db.Conn
	reservationEntityX entityx.Reservation
}

// NewReservation DI初期化関数
func NewReservation(
	conn *db.Conn,
	reservationEntityX entityx.Reservation,
) repository.Reservation {
	return &reservation{
		Conn:               conn,
		reservationEntityX: reservationEntityX,
	}
}

func (r *reservation) ExistsSpaceDoubleBooking(ctx context.Context, date time.Time, spaceID int64) (bool, error) {
	return entity.Reservations(
		entity.ReservationWhere.Date.EQ(date),
		entity.ReservationWhere.SpaceID.EQ(spaceID),
		entity.ReservationWhere.DeletedAt.IsNull(),
	).Exists(ctx, r.Conn)
}

func (r *reservation) ExistsBeauticianDoubleBooking(ctx context.Context, date time.Time, beauticianID int64) (bool, error) {
	return entity.Reservations(
		entity.ReservationWhere.Date.EQ(date),
		entity.ReservationWhere.BeauticianID.EQ(beauticianID),
		entity.ReservationWhere.DeletedAt.IsNull(),
	).Exists(ctx, r.Conn)
}

func (r *reservation) Create(ctx context.Context, re *entity.Reservation, menuIDs []int64) error {
	return r.Conn.RunInTx(ctx, func(tx *db.Tx) error {
		if err := re.Insert(ctx, tx, boil.Infer()); err != nil {
			return err
		}
		for _, menuID := range menuIDs {
			rem := r.reservationEntityX.NewReservationMenu(re.ID, menuID)
			if err := rem.Insert(ctx, tx, boil.Infer()); err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *reservation) FindByBeautician(ctx context.Context, beauticianID int64) (entity.ReservationSlice, error) {
	today, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	return entity.Reservations(
		qm.Where("date BETWEEN ? AND ?", fmt.Sprint(today), fmt.Sprint(today.AddDate(0, 0, 13))),
		entity.ReservationWhere.BeauticianID.EQ(beauticianID),
	).All(ctx, r.Conn)
}

func (r *reservation) FindByGuest(ctx context.Context, guestID int64) ([]*entityx.ReservationGetByGuest, error) {
	var results []*entityx.ReservationGetByGuest
	if err := entity.NewQuery(
		qm.Select("*"),
		qm.Select("salons.name AS salon_name"),
		qm.From(entity.TableNames.Reservations),
		qm.InnerJoin("spaces ON spaces.id = reservations.space_id"),
		qm.InnerJoin("salons ON salons.id = spaces.salon_id"),
		qm.InnerJoin("beauticians ON beauticians.id = reservations.beautician_id"),
		entity.ReservationWhere.DeletedAt.IsNull(),
		entity.ReservationWhere.GuestID.EQ(guestID),
	).Bind(ctx, r.Conn, &results); err != nil {
		return nil, err
	}
	return results, nil
}
