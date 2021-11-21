package users

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

type UserController struct {
	router  chi.Router
	service UserService
}

func NewUserController(r chi.Router, service UserService) *UserController {
	return &UserController{
		router:  r,
		service: service,
	}
}

func (u *UserController) RegisterRoutes() {
	u.router.Get("/users", u.getUsers)
	u.router.Get("/users/{id}", u.getUser)
}

// Get Users
// @Summary Get Users without penis
// @Description Get Users
// @ID get-users
// @Accept  json
// @Produce  json
// @Param name query string true "Name"
// @Param status query string true "Status"
// @Success 200 {object} User
// @Router /users [get]
func (u *UserController) getUsers(w http.ResponseWriter, _ *http.Request) {
	users := []User{
		{Name: "User 01"},
		{Name: "User 02"},
		{Name: "User 03"},
	}
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "generic error", http.StatusInternalServerError)
	}
}

// Get a specific user
// @Summary Get a specific user
// @Description Get a specific user
// @ID get-user
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} User
// @Router /users/{id} [get]
func (u *UserController) getUser(w http.ResponseWriter, r *http.Request) {
	user := &User{}
	render.JSON(w, r, user)
}
