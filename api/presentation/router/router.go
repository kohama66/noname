package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/myapp/noname/api/di"

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
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", r)
}

func (r *Router) Routes() {
	beauticianController := di.InitBeautician()
	reservationController := di.InitReservation()
	salonController := di.InitSalon()

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Group(func(r chi.Router) {
				r.Route("/beautician", func(r chi.Router) {
					r.Get("/find", beauticianController.Find)
					r.Group(func(r chi.Router) {
						r.Use(middleware.AuthAPI)
						r.Post("/", beauticianController.Create)
						r.Get("/", beauticianController.Get)
					})
				})
			})
			r.Group(func(r chi.Router) {
				r.Use(middleware.AuthAPI)
				r.Route("/reservation", func(r chi.Router) {
					r.Post("/", reservationController.Create)
					r.Get("/find", reservationController.Find)
					r.Get("/beautician", reservationController.FindByBeautician)
				})
			})
			r.Group(func(r chi.Router) {
				r.Route("/salon", func(r chi.Router) {
					r.Get("/find", salonController.Find)
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
