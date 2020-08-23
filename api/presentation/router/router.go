package router

import (
	"log"
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
	// fmt.Println("Listening on port 3000")
	// http.ListenAndServe(":3000", r)

	log.Printf("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}

func (r *Router) Routes() {
	testController := di.InitTest()

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Group(func(r chi.Router) {
				r.Use(middleware.AuthAPI)
				r.Route("/beautician", func(r chi.Router) {
					r.Post("/", testController.Get)
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
