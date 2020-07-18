// +build wireinject

package di

import (
	"github.com/myapp/noname/api/application/usecase"
	"github.com/myapp/noname/api/presentation/v1/handler"

	"github.com/google/wire"
)

func InitTest() handler.Test {
	wire.Build(
		usecase.NewTest,
		handler.NewTest,
	)
	return nil
}
