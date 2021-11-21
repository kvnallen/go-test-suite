package users

import (
	"log"
)

type UserStore interface {
	Store(user *User) error
}

type userStore struct {
}

func NewUserStore() UserStore {
	return &userStore{}
}

func (u *userStore) Store(user *User) error {
	log.Println("saving user into DB")
	log.Printf("user %s saved successfully!\n", user)
	return nil
}
