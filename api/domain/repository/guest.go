package repository

import (
	"context"

	"github.com/myapp/noname/api/domain/entity"
)

// Guest DIInterface
type Guest interface {
	GetByAuthID(ctx context.Context, authID string) (*entity.Guest, error)
}
