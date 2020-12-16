package usecase

import (
	"context"
	"time"

	"github.com/myapp/noname/api/application/usecase/requestmodel"
	"github.com/myapp/noname/api/application/usecase/responsemodel"
	"github.com/myapp/noname/api/infrastructure/response"

	"github.com/myapp/noname/api/domain/repository"
	"github.com/rs/xid"
)

// Beautician 美容師
type Beautician interface {
	Create(ctx context.Context, r *requestmodel.BeauticianCreate) (*responsemodel.BeauticianCreate, error)
	Get(ctx context.Context, r *requestmodel.BeauticianGet) (*responsemodel.BeauticianGet, error)
	Find(ctx context.Context, r *requestmodel.BeauticianFind) (*responsemodel.BeauticianFind, error)
}

type beautician struct {
	beauticianRepository repository.Beautician
	beauticianResponse   response.Beautician
	menuRepository       repository.Menu
	salonRepository      repository.Salon
}

// NewBeautician usecaseの初期化
func NewBeautician(
	beauticianRepository repository.Beautician,
	beauticianResponse response.Beautician,
	menuRepository repository.Menu,
	salonRepository repository.Salon,
) Beautician {
	return &beautician{
		beauticianRepository: beauticianRepository,
		beauticianResponse:   beauticianResponse,
		menuRepository:       menuRepository,
		salonRepository:      salonRepository,
	}
}

func (b *beautician) Create(ctx context.Context, r *requestmodel.BeauticianCreate) (*responsemodel.BeauticianCreate, error) {
	ent := r.NewBeautician(xid.New().String())
	if err := b.beauticianRepository.Create(ctx, ent); err != nil {
		return nil, err
	}
	bt, err := b.beauticianRepository.GetByAuthID(ctx, ent.AuthID)
	if err != nil {
		return nil, err
	}
	res := b.beauticianResponse.NewBeauticianCreate(bt)
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
	ent, err := b.beauticianRepository.GetByRandID(ctx, r.RandID)
	if err != nil {
		return nil, err
	}
	res := b.beauticianResponse.NewBeauticianGet(ent)
	return res, nil
}
