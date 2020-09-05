// +build wireinject

package di

import (
	"github.com/myapp/noname/api/application/usecase"
	"github.com/myapp/noname/api/infrastructure/db"
	"github.com/myapp/noname/api/infrastructure/repository"
	"github.com/myapp/noname/api/infrastructure/response"
	"github.com/myapp/noname/api/presentation/v1/handler"

	"github.com/google/wire"
)

func InitBeautician() handler.Beautician {
	wire.Build(
		db.New,
		response.NewBeautician,
		repository.NewBeautician,
		usecase.NewBeautician,
		handler.NewBeautician,
	)
	return nil
}
