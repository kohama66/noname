package entityx

import (
	"github.com/myapp/noname/api/application/usecase/request"
	"github.com/myapp/noname/api/domain/entity"
)

func (b request.BeauticianCreate) NewBeautician() *entity.Beautician {
	return &entity.Beautician{
		AuthID: b.AuthID,
	}
}
