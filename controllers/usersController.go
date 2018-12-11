package controllers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"

	"github.com/laninog/go-simpleRestAPI/models"
	"github.com/laninog/go-simpleRestAPI/repositories"
)

type usersController struct {
	repository repositories.Repository
}

func NewUsersController(r repositories.Repository, gr *chi.Mux) {
	handler := &usersController{
		repository: r,
	}

	gr.Get("/users/{id}", handler.GetUser)
	gr.Post("/users", handler.CreateUser)
	gr.Put("/users", handler.UpdateUser)
	gr.Delete("/users/{id}", handler.DeleteUser)
	gr.Get("/users", handler.GetUsers)

}

func (c *usersController) GetUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	user, err := c.repository.FindByID(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		render.JSON(w, r, user)
	}
}

func (c *usersController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user *models.User
	render.DecodeJSON(r.Body, &user)
	user = c.repository.Add(user)
	w.Header().Add("location", r.RequestURI + "/" + user.ID)
	w.WriteHeader(http.StatusCreated)
}

func (c *usersController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	res := make(map[string]string)
	res["res"] = "Updated"
	render.JSON(w, r, res)
}

func (c *usersController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	_, err := c.repository.Remove(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (c *usersController) GetUsers(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, c.repository.FindAll())
}