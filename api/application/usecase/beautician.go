package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/myapp/noname/api/application/usecase/requestmodel"
	"github.com/myapp/noname/api/application/usecase/responsemodel"
	"github.com/myapp/noname/api/infrastructure/response"

	"github.com/myapp/noname/api/domain/repository"
)

// Beautician 美容師
type Beautician interface {
	Create(ctx context.Context, r *requestmodel.BeauticianCreate) (*responsemodel.BeauticianCreate, error)
	Get(ctx context.Context, r *requestmodel.BeauticianGet) (*responsemodel.BeauticianGet, error)
	Find(ctx context.Context, r *requestmodel.BeauticianFind) (*responsemodel.BeauticianFind, error)
	Update(ctx context.Context, r *requestmodel.BeauticianUpdate) error
	GetMyPage(ctx context.Context, r *requestmodel.BeauticianMyPageGet) (*responsemodel.BeauticianMyPageGet, error)
}

type beautician struct {
	beauticianRepository repository.Beautician
	beauticianResponse   response.Beautician
	menuRepository       repository.Menu
	salonRepository      repository.Salon
	userRepository       repository.User
}

// NewBeautician usecaseの初期化
func NewBeautician(
	beauticianRepository repository.Beautician,
	beauticianResponse response.Beautician,
	menuRepository repository.Menu,
	salonRepository repository.Salon,
	userRepository repository.User,
) Beautician {
	return &beautician{
		beauticianRepository: beauticianRepository,
		beauticianResponse:   beauticianResponse,
		menuRepository:       menuRepository,
		salonRepository:      salonRepository,
		userRepository:       userRepository,
	}
}

func (b *beautician) Create(ctx context.Context, r *requestmodel.BeauticianCreate) (*responsemodel.BeauticianCreate, error) {
	me, err := b.userRepository.GetByAuthID(ctx, r.AuthID)
	if err != nil {
		return nil, err
	}
	ent := r.NewBeautician(me.ID)
	if err := b.beauticianRepository.Create(ctx, ent); err != nil {
		return nil, err
	}
	res := b.beauticianResponse.NewBeauticianCreate(ent)
	return res, nil
}

// func (b *beautician) Get(ctx context.Context, r *requestmodel.BeauticianGet) (*responsemodel.BeauticianGet, error) {
// 	ent, err := b.beauticianRepository.GetByAuthID(ctx, r.AuthID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	res := b.beauticianResponse.NewBeauticianGet(ent)
// 	return res, nil
// }

func (b *beautician) Find(ctx context.Context, r *requestmodel.BeauticianFind) (*responsemodel.BeauticianFind, error) {
	var date time.Time
	var menu []int64
	var salon *int64
	if r.Date != "" {
		dt, err := time.Parse("2006-01-02 15:04:05", r.Date)
		if err != nil {
			return nil, err
		}
		date = dt
	}
	if len(r.MenuRandIDs) != 0 {
		mns, err := b.menuRepository.FindByRandID(ctx, r.MenuRandIDs)
		if err != nil {
			return nil, err
		}
		for _, v := range mns {
			menu = append(menu, v.ID)
		}
	}
	if r.SalonRandID != "" {
		sl, err := b.salonRepository.GetByRandID(ctx, r.SalonRandID)
		if err != nil {
			return nil, err
		}
		salon = &sl.ID
	}
	ents, err := b.beauticianRepository.Find(ctx, date, salon, menu)
	if err != nil {
		return nil, err
	}
	return b.beauticianResponse.NewBeauticianFind(ents), nil
}

func (b *beautician) Get(ctx context.Context, r *requestmodel.BeauticianGet) (*responsemodel.BeauticianGet, error) {
	us, err := b.userRepository.GetByAuthID(ctx, r.AuthID)
	if err != nil {
		return nil, err
	}
	ent, err := b.beauticianRepository.GetByUserID(ctx, us.ID)
	if err != nil {
		return nil, err
	}
	return b.beauticianResponse.NewBeauticianGet(ent), nil
}

func (b *beautician) Update(ctx context.Context, r *requestmodel.BeauticianUpdate) error {
	us, err := b.userRepository.GetByAuthID(ctx, r.AuthID)
	if err != nil {
		return err
	}
	if !us.IsBeautician {
		return errors.New("not beautician")
	}
	bt, err := b.beauticianRepository.GetByUserID(ctx, us.ID)
	if err != nil {
		return err
	}
	us, bt = r.NewBeautician(us, bt)
	if err := b.beauticianRepository.Update(ctx, us, bt); err != nil {
		return err
	}
	return nil
}

func (b *beautician) GetMyPage(ctx context.Context, r *requestmodel.BeauticianMyPageGet) (*responsemodel.BeauticianMyPageGet, error) {
	u, err := b.userRepository.GetByAuthID(ctx, r.AuthID)
	if err != nil {
		return nil, err
	}
	be, err := b.beauticianRepository.GetByUserID(ctx, u.ID)
	if err != nil {
		return nil, err
	}
	s, err := b.salonRepository.FindByBeauticianID(ctx, u.ID)
	if err != nil {
		return nil, err
	}
	return b.beauticianResponse.NewBeauticianMyPageGet(u, be, s), nil
}
