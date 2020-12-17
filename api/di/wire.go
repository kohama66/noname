// +build wireinject

package di

import (
	"github.com/myapp/noname/api/application/usecase"
	"github.com/myapp/noname/api/domain/entityx"
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
		repository.NewSalon,
		repository.NewUser,
		repository.NewMenu,
		usecase.NewBeautician,
		handler.NewBeautician,
	)
	return nil
}

func InitReservation() handler.Reservation {
	wire.Build(
		db.New,
		entityx.NewReservation,
		repository.NewUser,
		repository.NewMenu,
		repository.NewReservation,
		repository.NewBeautician,
		repository.NewSalon,
		response.NewReservation,
		usecase.NewReservation,
		handler.NewReservation,
	)
	return nil
}

func InitSalon() handler.Salon {
	wire.Build(
		db.New,
		response.NewSalon,
		repository.NewSalon,
		repository.NewUser,
		usecase.NewSalon,
		handler.NewSalon,
	)
	return nil
}

func InitMenu() handler.Menu {
	wire.Build(
		db.New,
		response.NewMenu,
		repository.NewUser,
		repository.NewMenu,
		usecase.NewMenu,
		handler.NewMenu,
	)
	return nil
}

func InitUser() handler.User {
	wire.Build(
		db.New,
		response.NewUser,
		entityx.NewReservation,
		repository.NewUser,
		repository.NewReservation,
		usecase.NewUser,
		handler.NewUser,
	)
	return nil
}
