package users_test

import (
	"errors"
	"meu.pack/internal/users"
	"meu.pack/internal/users/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type SuiteStore struct {
	suite.Suite
	userStore *mocks.UserStoreMock
	users.UserService
}

func (s *SuiteStore) OnStore(user *users.User, err error) {
	s.userStore.On("Store", user).Return(err)
}

func (s *SuiteStore) SetupTest() {
	s.userStore = new(mocks.UserStoreMock)
	s.UserService = users.NewUserService(s.userStore)
}

func (s *SuiteStore) TestStore() {

	s.Run("when there's a error, should return it ", func() {
		user := &users.User{Name: "Kevin"}
		s.OnStore(user, errors.New("some error"))

		// act
		err := s.UserService.Store(user)

		// assert
		s.Require().Error(err)
	})

	s.Run("when success", func() {
		// arrange
		user := &users.User{Name: "Kevin"}

		// act
		_ = s.UserService.Store(user)

		// assert
		s.userStore.AssertCalled(s.T(), "Store",
			mock.MatchedBy(func(_u *users.User) bool {
				return _u.Name == "Kevin"
			}),
		)
	})
}

func TestUserServiceStore(t *testing.T) {

	suite.Run(t, new(SuiteStore))
}
