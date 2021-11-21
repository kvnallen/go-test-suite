package users

import (
	"log"
)

type UserService interface {
	Store(user *User) error
}

type userService struct {
	UserStore
}

func NewUserService(store UserStore) UserService {
	return &userService{
		store,
	}
}

func (s *userService) Store(user *User) error {
	log.Println("service: saving user into store")
	return s.UserStore.Store(user)
}
