package repository

import (
	"context"
	"time"

	"github.com/myapp/noname/api/domain/entity"
)

// Salon DIInterface
type Salon interface {
	GetByRandID(ctx context.Context, randID string) (*entity.Salon, error)
	GetAll(ctx context.Context) (entity.SalonSlice, error)
	// FindByBeautician(ctx context.Context, beauticianID int64) (entity.SalonSlice, error)
	Find(ctx context.Context, beauticianID *int64, date *time.Time) (entity.SalonSlice, error)
	GetVacantSpace(ctx context.Context, date time.Time, salonID int64) (*entity.Space, error)
	ExistsByBeauticianWithSalon(ctx context.Context, beauticianID, salonID int64) (bool, error)
	GetBeauticianSalons(ctx context.Context, beauticianID int64) (entity.SalonSlice, error)
	GetBySpaceID(ctx context.Context, spaceID int64) (*entity.Salon, error)
	FindNotBelongs(ctx context.Context, beauticianID int64) (entity.SalonSlice, error)
	CreateBeauticianSalon(ctx context.Context, ent *entity.BeauticianSalon) error
	DeleteBeauticianSalon(ctx context.Context, ent *entity.BeauticianSalon) (int64, error)
}
