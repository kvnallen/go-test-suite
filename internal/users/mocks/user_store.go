package mocks

import (
	"github.com/stretchr/testify/mock"
	"meu.pack/internal/users"
)

type UserStoreMock struct {
	mock.Mock
}

func (_m *UserStoreMock) Store(user *users.User) error {
	args := _m.Called(user)
	return args.Error(0)
}
