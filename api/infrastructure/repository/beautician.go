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

func (b *beautician) GetByUserID(ctx context.Context, ID int64) (*entity.Beautician, error) {
	return entity.Beauticians(
		entity.BeauticianWhere.UserID.EQ(ID),
		entity.BeauticianWhere.DeletedAt.IsNull(),
		// qm.Load(entity.BeauticianRels.BeauticianMenus),
	).One(ctx, b.Conn)
}

// func (b *beautician) GetByRandID(ctx context.Context, randID string) (*entity.Beautician, error) {
// 	return entity.Beauticians(
// 		entity.BeauticianWhere.RandID.EQ(randID),
// 		entity.BeauticianWhere.DeletedAt.IsNull(),
// 		qm.Load(entity.BeauticianRels.BeauticianMenus),
// 	).One(ctx, b.Conn)
// }

func (b *beautician) GetAll(ctx context.Context) ([]*entity.User, error) {
	return entity.Users(
		entity.UserWhere.DeletedAt.IsNull(),
		qm.Load(entity.UserRels.BeauticianBeauticianMenus),
		qm.Load(entity.UserRels.Beautician),
	).All(ctx, b.Conn)
}

// func (b *beautician) Find(ctx context.Context, date time.Time, salon *int64, menus []int64) ([]*entity.Beautician, error) {
// 	qms := []qm.QueryMod{}
// 	// メニュー用SQL文
// 	if len(menus) != 0 {
// 		convertedIDs := make([]interface{}, len(menus))
// 		for i, v := range menus {
// 			convertedIDs[i] = v
// 		}
// 		qms = append(qms,
// 			qm.Select(entity.TableNames.Beauticians+".*"),
// 			qm.InnerJoin("beautician_menus ON beautician_menus.beautician_id = beauticians.id"),
// 			qm.WhereIn("beautician_menus.menu_id IN ?", convertedIDs...),
// 			qm.GroupBy("beauticians.id"),
// 			qm.Having("COUNT(*) = ?", len(menus)),
// 		)
// 	}
// 	//美容院用SQL文
// 	if salon != nil {
// 		fmt.Println("salon")
// 		qms = append(qms, entity.SalonWhere.ID.EQ(*salon), qm.InnerJoin("beautician_salons ON beautician_salons.beautician_id = beauticians.id"),
// 			qm.InnerJoin("salons ON salons.id = beautician_salons.salon_id"))
// 	}
// 	//日付用SQL文
// 	if !date.IsZero() {
// 		qms = append(qms,
// 			qm.Select(entity.TableNames.Beauticians+".*"),
// 			qm.LeftOuterJoin(entity.TableNames.Reservations+" ON reservations.beautician_id = beauticians.id"),
// 			qm.GroupBy("beauticians.id"),
// 			qm.Having("COUNT(date = ? OR NULL) = 0", date),
// 		)
// 	}
// 	//条件無しの場合はGetAllへ飛ばす
// 	if len(qms) == 0 {
// 		return b.GetAll(ctx)
// 	}
// 	// 共通SQL文
// 	qms = append(qms, entity.BeauticianWhere.DeletedAt.IsNull())

// 	return entity.Beauticians(
// 		qms...,
// 	).All(ctx, b.Conn)
// }

// func (b *beautician) Find(ctx context.Context, date time.Time, salon *int64, menus []int64) ([]*entity.Beautician, error) {
// 	var bts entity.BeauticianSlice
// 	var count int64
// 	if !date.IsZero() {
// 		bt, err := entity.Beauticians(
// 			qm.Select(entity.TableNames.Beauticians+".*"),
// 			qm.LeftOuterJoin(entity.TableNames.Reservations+" ON reservations.beautician_id = beauticians.id"),
// 			qm.GroupBy("beauticians.id"),
// 			qm.Having("COUNT(date = ? OR NULL) = 0", date),
// 			qm.Load(entity.BeauticianRels.BeauticianMenus),
// 		).All(ctx, b.Conn)
// 		if err != nil {
// 			return nil, err
// 		}
// 		bts = append(bts, bt...)
// 		count++
// 	}
// 	if salon != nil {
// 		bt, err := entity.Beauticians(
// 			entity.BeauticianSalonWhere.SalonID.EQ(*salon),
// 			qm.InnerJoin("beautician_salons ON beautician_salons.beautician_id = beauticians.id"),
// 			qm.Load(entity.BeauticianRels.BeauticianMenus),
// 		).All(ctx, b.Conn)
// 		if err != nil {
// 			return nil, err
// 		}
// 		bts = append(bts, bt...)
// 		count++
// 	}
// 	if len(menus) != 0 {
// 		convertedIDs := make([]interface{}, len(menus))
// 		for i, v := range menus {
// 			convertedIDs[i] = v
// 		}
// 		bt, err := entity.Beauticians(
// 			qm.Select(entity.TableNames.Beauticians+".*"),
// 			qm.InnerJoin("beautician_menus ON beautician_menus.beautician_id = beauticians.id"),
// 			qm.WhereIn("beautician_menus.menu_id IN ?", convertedIDs...),
// 			qm.GroupBy("beauticians.id"),
// 			qm.Having("COUNT(*) = ?", len(menus)),
// 			qm.Load(entity.BeauticianRels.BeauticianMenus),
// 		).All(ctx, b.Conn)
// 		if err != nil {
// 			return nil, err
// 		}
// 		bts = append(bts, bt...)
// 		count++
// 	}
// 	if count == 0 {
// 		return b.GetAll(ctx)
// 	}
// 	var beauticians entity.BeauticianSlice
// 	m := make(map[int64]int64)
// 	for _, v := range bts {
// 		if ct, ok := m[v.ID]; !ok {
// 			m[v.ID]++
// 			if count == 1 {
// 				beauticians = append(beauticians, v)
// 			}
// 		} else if ct < count {
// 			m[v.ID]++
// 			if m[v.ID] == count {
// 				beauticians = append(beauticians, v)
// 			}
// 		}
// 	}
// 	return beauticians, nil
// }

func (b *beautician) FindPossibleSalon(ctx context.Context, beauciaonID int64) (entity.BeauticianSalonSlice, error) {
	return entity.BeauticianSalons(
		entity.BeauticianSalonWhere.BeauticianID.EQ(beauciaonID),
		entity.BeauticianSalonWhere.DeletedAt.IsNull(),
	).All(ctx, b.Conn)
}

func (b *beautician) Find(ctx context.Context, date time.Time, salon *int64, menus []int64) ([]*entity.User, error) {
	var uss entity.UserSlice
	var count int64
	if !date.IsZero() {
		us, err := entity.Users(
			qm.Select(entity.TableNames.Users+".*"),
			qm.LeftOuterJoin(entity.TableNames.Reservations+" ON reservations.beautician_id = users.id"),
			qm.GroupBy("users.id"),
			qm.Having("COUNT(date = ? OR NULL) = 0", date),
			qm.Load(entity.UserRels.BeauticianBeauticianMenus),
			qm.Load(entity.UserRels.Beautician),
		).All(ctx, b.Conn)
		if err != nil {
			return nil, err
		}
		uss = append(uss, us...)
		count++
	}
	if salon != nil {
		us, err := entity.Users(
			entity.BeauticianSalonWhere.SalonID.EQ(*salon),
			qm.InnerJoin("beautician_salons ON beautician_salons.beautician_id = users.id"),
			qm.Load(entity.UserRels.BeauticianBeauticianMenus),
			qm.Load(entity.UserRels.Beautician),
		).All(ctx, b.Conn)
		if err != nil {
			return nil, err
		}
		uss = append(uss, us...)
		count++
	}
	if len(menus) != 0 {
		convertedIDs := make([]interface{}, len(menus))
		for i, v := range menus {
			convertedIDs[i] = v
		}
		us, err := entity.Users(
			qm.Select(entity.TableNames.Users+".*"),
			qm.InnerJoin("beautician_menus ON beautician_menus.beautician_id = users.id"),
			qm.WhereIn("beautician_menus.menu_id IN ?", convertedIDs...),
			qm.GroupBy("users.id"),
			qm.Having("COUNT(*) = ?", len(menus)),
			qm.Load(entity.UserRels.BeauticianBeauticianMenus),
			qm.Load(entity.UserRels.Beautician),
		).All(ctx, b.Conn)
		if err != nil {
			return nil, err
		}
		uss = append(uss, us...)
		count++
	}
	if count == 0 {
		return b.GetAll(ctx)
	}
	var beauticians entity.UserSlice
	m := make(map[int64]int64)
	for _, v := range uss {
		if ct, ok := m[v.ID]; !ok {
			m[v.ID]++
			if count == 1 {
				beauticians = append(beauticians, v)
			}
		} else if ct < count {
			m[v.ID]++
			if m[v.ID] == count {
				beauticians = append(beauticians, v)
			}
		}
	}
	return beauticians, nil
}
