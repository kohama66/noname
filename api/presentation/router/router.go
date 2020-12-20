package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/myapp/noname/api/di"
	"github.com/myapp/noname/api/env"

	_ "github.com/myapp/noname/api/docs"
	"github.com/myapp/noname/api/presentation/middleware"

	//swagger import
	httpSwagger "github.com/swaggo/http-swagger"
)

type Router struct {
	chi.Router
}

func Init() *Router {
	r := chi.NewRouter()
	return &Router{r}
}

func (r *Router) Run() {
	port := env.GetPort()
	fmt.Println("Listening on port " + port)
	http.ListenAndServe(":"+port, r)
}

func (r *Router) Routes() {
	beauticianController := di.InitBeautician()
	reservationController := di.InitReservation()
	salonController := di.InitSalon()
	menuController := di.InitMenu()
	userController := di.InitUser()

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Group(func(r chi.Router) {
				r.Route("/beautician", func(r chi.Router) {
					r.Get("/find", beauticianController.Find)
					r.Group(func(r chi.Router) {
						r.Use(middleware.AuthAPI)
						r.Post("/", beauticianController.Create)
						r.Get("/", beauticianController.GetByAuthID)
						r.Put("/", beauticianController.Update)
					})
				})
			})
			r.Group(func(r chi.Router) {
				r.Route("/reservation", func(r chi.Router) {
					r.Get("/find", reservationController.Find)
					r.Group(func(r chi.Router) {
						r.Use(middleware.AuthAPI)
						r.Post("/", reservationController.Create)
						r.Get("/beautician", reservationController.FindByBeautician)
						r.Get("/user", reservationController.FindByUser)
					})
				})
			})
			r.Group(func(r chi.Router) {
				r.Route("/salon", func(r chi.Router) {
					r.Get("/find", salonController.Find)
				})
			})
			r.Group(func(r chi.Router) {
				r.Route("/menu", func(r chi.Router) {
					r.Get("/find", menuController.Find)
					r.Get("/find/{beauticianRandID}", menuController.FindByBeauticianWithMenuRandID)
				})
			})
			r.Group(func(r chi.Router) {
				r.Use(middleware.AuthAPI)
				r.Route("/user", func(r chi.Router) {
					r.Get("/", userController.Get)
					r.Post("/", userController.Create)
				})
			})
		})
	})
}

func (r *Router) Swagger() {
	r.Use(middleware.CORS)
	r.Get("/api/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/api/swagger/doc.json"),
	))
}
