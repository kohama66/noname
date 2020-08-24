package repository

import (
	"context"

	"github.com/myapp/noname/api/domain/entity"
)

type Beautician interface {
	Create(ctx context.Context, ent *entity.Beautician) error
}
