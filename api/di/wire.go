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
		repository.NewSalon,
		repository.NewMenu,
		usecase.NewBeautician,
		handler.NewBeautician,
	)
	return nil
}

func InitReservation() handler.Reservation {
	wire.Build(
		db.New,
		repository.NewGuest,
		repository.NewReservation,
		repository.NewBeautician,
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
		repository.NewBeautician,
		usecase.NewSalon,
		handler.NewSalon,
	)
	return nil
}

func InitMenu() handler.Menu {
	wire.Build(
		db.New,
		response.NewMenu,
		repository.NewBeautician,
		repository.NewMenu,
		usecase.NewMenu,
		handler.NewMenu,
	)
	return nil
}
