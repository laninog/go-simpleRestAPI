package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"log"
	"net/http"

	_controllers "github.com/laninog/go-simpleRestAPI/controllers"
	_repositories "github.com/laninog/go-simpleRestAPI/repositories"
)

func InitControllers() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	_repository := _repositories.NewUserRepository()
	_controllers.NewUsersController(_repository, router)

	return router
}

func main() {

	routes := InitControllers()

	walkFunc := func (method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(routes, walkFunc); err != nil {
		log.Panicf("Error!! %s", err.Error())
	}

	log.Fatal(http.ListenAndServe(":8000", routes))
}