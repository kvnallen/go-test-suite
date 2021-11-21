package mathutils_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	mathutils "meu.pack/math"
)

type MathSuite struct {
	suite.Suite
	numbers []int
}

// The SetupSuite method will be run by testify once, at the very
// start of the testing suite, before any tests are run.
func (s *MathSuite) SetupSuite() {
	fmt.Println("The start of everything")
}

// The TearDownSuite method will be run by testify once, at the very
// end of the testing suite, after all tests have been run.
func (s *MathSuite) TearDownSuite() {
	fmt.Println("The end of everything")
}

func (s *MathSuite) SetupTest() {
	s.numbers = []int{1, 2}
	fmt.Println("setuping test")
}

func (s *MathSuite) TestAllPositive() {
	require := s.Assert()
	require.Equal(3, mathutils.Sum(s.numbers))
}

func (s *MathSuite) TestAllNegative() {
	require := s.Assert()
	require.Equal(3, mathutils.Sum([]int{2, 1}))
}

func (s *MathSuite) TestAny() {

	tests := []struct {
		name    string
		numbers []int
	}{
		{
			name:    "01",
			numbers: []int{1, 2},
		},
		{
			name:    "02",
			numbers: []int{2, 1},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			s.T().Parallel()
			s.SetupTest()
			s.Assert().Equal(3, mathutils.Sum(test.numbers))
		})
	}
}

func TestSuite(t *testing.T) {
	suite.Run(t, &MathSuite{})
}
