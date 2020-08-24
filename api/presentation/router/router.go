package router

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/myapp/noname/api/di"
	"net/http"

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
	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", r)
}

func (r *Router) Routes() {
	beauticianController := di.InitBeautician()
	testController := di.InitTest()

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Group(func(r chi.Router) {
				r.Use(middleware.AuthAPI)
				r.Route("/beautician", func(r chi.Router) {
					r.Post("/", beauticianController.Create)
				})
			})
			r.Route("/hello", func(r chi.Router) {
				r.Get("/", testController.Get)
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
