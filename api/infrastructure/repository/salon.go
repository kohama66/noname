package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/myapp/noname/api/domain/entity"
	"github.com/myapp/noname/api/domain/repository"
	"github.com/myapp/noname/api/infrastructure/db"
	"github.com/volatiletech/sqlboiler/boil"
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

func (s *salon) GetAll(ctx context.Context) (entity.SalonSlice, error) {
	return entity.Salons(
		entity.SalonWhere.DeletedAt.IsNull(),
	).All(ctx, s.Conn)
}

// func (s *salon) FindByBeautician(ctx context.Context, beauticianID int64) (entity.SalonSlice, error) {
// 	return entity.Salons(
// 		qm.InnerJoin("beautician_salons ON salons.id = beautician_salons.salon_id"),
// 		qm.InnerJoin("beauticians ON beauticians.id = beautician_salons.beautician_id"),
// 		entity.BeauticianWhere.ID.EQ(beauticianID),
// 		entity.SalonWhere.DeletedAt.IsNull(),
// 	).All(ctx, s.Conn)
// }

func (s *salon) Find(ctx context.Context, beauticianID *int64, date *time.Time) (entity.SalonSlice, error) {
	var sls entity.SalonSlice
	var count int
	if date != nil {
		sl, err := entity.Salons(
			qm.Distinct(entity.TableNames.Salons+".*"),
			qm.InnerJoin(fmt.Sprintf("%s ON %s = %s", entity.TableNames.Spaces, entity.SpaceColumns.SalonID, "salons.id")),
			qm.LeftOuterJoin(fmt.Sprintf("%s ON %s = %s", entity.TableNames.Reservations, entity.ReservationColumns.SpaceID, "spaces.id")),
			qm.GroupBy("spaces.id"),
			qm.Having("COUNT(reservations.date = ? OR NULL) = 0", &date),
			entity.SalonWhere.DeletedAt.IsNull(),
		).All(ctx, s.Conn)
		if err != nil {
			return nil, err
		}
		sls = append(sls, sl...)
		count++
	}
	if beauticianID != nil {
		sl, err := entity.Salons(
			qm.InnerJoin(fmt.Sprintf("beautician_salons ON beautician_salons.salon_id = salons.id")),
			qm.InnerJoin(fmt.Sprintf("users ON users.id = beautician_salons.beautician_id")),
			entity.UserWhere.ID.EQ(*beauticianID),
			entity.SalonWhere.DeletedAt.IsNull(),
		).All(ctx, s.Conn)
		if err != nil {
			return nil, err
		}
		sls = append(sls, sl...)
		count++
	}
	if count == 0 {
		return s.GetAll(ctx)
	} else if count == 1 {
		return sls, nil
	}
	var salons entity.SalonSlice
	m := make(map[int64]int64)
	for _, sl := range sls {
		if _, ok := m[sl.ID]; !ok {
			m[sl.ID] = sl.ID
		} else {
			salons = append(salons, sl)
		}
	}
	return salons, nil
}

func (s *salon) GetVacantSpace(ctx context.Context, date time.Time, salonID int64) (*entity.Space, error) {
	return entity.Spaces(
		qm.LeftOuterJoin("reservations ON reservations.space_id = spaces.id"),
		entity.SpaceWhere.SalonID.EQ(salonID),
		qm.GroupBy("spaces.id"),
		qm.Having("COUNT(reservations.date = ? OR NULL) = 0", date),
	).One(ctx, s.Conn)
}

func (s *salon) ExistsByBeauticianWithSalon(ctx context.Context, beauticianID, salonID int64) (bool, error) {
	return entity.BeauticianSalons(
		entity.BeauticianSalonWhere.BeauticianID.EQ(beauticianID),
		entity.BeauticianSalonWhere.SalonID.EQ(salonID),
		entity.BeauticianSalonWhere.DeletedAt.IsNull(),
	).Exists(ctx, s.Conn)
}

// func (s *salon) GetBeauticianSalons(ctx context.Context, beauticianID int64) (entity.SalonSlice, error) {
// 	bs, err := entity.BeauticianSalons(
// 		entity.BeauticianSalonWhere.BeauticianID.EQ(beauticianID),
// 		entity.BeauticianSalonWhere.DeletedAt.IsNull(),
// 	).All(ctx, s.Conn)
// 	if err != nil {
// 		return nil, err
// 	}
// 	ids := make([]int64, len(bs))
// 	for i, v := range bs {
// 		ids[i] = v.SalonID
// 	}
// 	return entity.Salons(
// 		entity.SalonWhere.ID.IN(ids),
// 		entity.SalonWhere.DeletedAt.IsNull(),
// 	).All(ctx, s.Conn)
// }

func (s *salon) FindByBeauticianID(ctx context.Context, beauticianID int64) (entity.SalonSlice, error) {
	return entity.Salons(
		qm.InnerJoin(fmt.Sprintf("%s ON %s.%v = %s.%v", entity.TableNames.BeauticianSalons, entity.TableNames.BeauticianSalons, entity.BeauticianSalonColumns.SalonID, entity.TableNames.Salons, entity.SalonColumns.ID)),
		entity.BeauticianSalonWhere.BeauticianID.EQ(beauticianID),
		entity.SalonWhere.DeletedAt.IsNull(),
		entity.BeauticianSalonWhere.DeletedAt.IsNull(),
	).All(ctx, s.Conn)
}

func (s *salon) GetBySpaceID(ctx context.Context, spaceID int64) (*entity.Salon, error) {
	return entity.Salons(
		qm.InnerJoin("spaces ON spaces.salon_id = salons.id"),
		entity.SpaceWhere.ID.EQ(spaceID),
		entity.SpaceWhere.DeletedAt.IsNull(),
		entity.SalonWhere.DeletedAt.IsNull(),
	).One(ctx, s.Conn)
}

func (s *salon) FindNotBelongs(ctx context.Context, beauticianID int64) (entity.SalonSlice, error) {
	return entity.Salons(
		qm.LeftOuterJoin(fmt.Sprintf(`%s ON %v = %v AND %v = %v`, entity.TableNames.BeauticianSalons, entity.BeauticianSalonColumns.SalonID, entity.SalonColumns.ID, entity.BeauticianSalonColumns.BeauticianID, beauticianID)),
		qm.Where(fmt.Sprintf("%v IS NULL", entity.BeauticianMenuColumns.BeauticianID)),
		entity.SalonWhere.DeletedAt.IsNull(),
	).All(ctx, s.Conn)
}

func (s *salon) CreateBeauticianSalon(ctx context.Context, ent *entity.BeauticianSalon) error {
	return ent.Insert(ctx, s.Conn, boil.Infer())
}

func (s *salon) DeleteBeauticianSalon(ctx context.Context, ent *entity.BeauticianSalon) (int64, error) {
	return ent.Delete(ctx, s.Conn)
}

func (s *salon) Create(ctx context.Context, ent *entity.Salon) error {
	return ent.Insert(ctx, s.Conn, boil.Infer())
}

func (s *salon) CreateUserSalon(ctx context.Context, ent *entity.UserSalon) error {
	return ent.Insert(ctx, s.Conn, boil.Infer())
}
