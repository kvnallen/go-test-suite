package mathutils_test

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	mathutils "meu.pack/math"
	"testing"
)

type SuiteMultiply struct {
	suite.Suite
	isValid bool
}

func TestMultiplySuite(t *testing.T) {
	suite.Run(t, new(SuiteMultiply))
}

func (s SuiteMultiply) SetupSuite() {
	fmt.Println("Starting suite...")
}

func (s SuiteMultiply) SetupTest() {
	fmt.Println("Setuping!")
	s.isValid = true
}

func (s SuiteMultiply) TestMultiply() {
	require := s.Require()
	require.Equal(1, mathutils.Multiply(1, 1))
}

func ExampleSum() {
	fmt.Println(mathutils.Sum([]int{4, 6}))
	// Output: 10
}

func (s SuiteMultiply) TearDownTest() {
	fmt.Println("The end of the tests")
}
