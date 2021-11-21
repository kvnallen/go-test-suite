package infra

import (
	"github.com/go-chi/chi/v5"
	"meu.pack/internal/users"
)

// NewUserController Simple factory that creates the dependency hierarchy
func NewUserController(r chi.Router) *users.UserController {
	userStore := users.NewUserStore()
	userService := users.NewUserService(userStore)
	return users.NewUserController(r, userService)
}
