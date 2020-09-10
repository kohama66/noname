package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/myapp/noname/api/di"

	//swagger import
	_ "github.com/myapp/noname/api/docs"
	"github.com/myapp/noname/api/presentation/middleware"

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

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Group(func(r chi.Router) {
				r.Use(middleware.AuthAPI)
				r.Route("/beautician", func(r chi.Router) {
					r.Post("/", beauticianController.Create)
					r.Get("/", beauticianController.Get)
				})
				r.Route("/reservation", func(r chi.Router) {
					r.Post("/", reservationController.Create)
					r.Get("/beautician", reservationController.FindByBeautician)
				})
			})
		})
	})
}

func (r *Router) Swagger() {
	// r.Use(middleware.CORS)
	r.Get("/api/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/api/swagger/doc.json"),
	))
}
