package testify_demo

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"testing"
)

// Test object
type MyMockedObject struct {
	mock.Mock
}

// DoSomething is a method on MyMockedObject that implements some interface
// and just records the activity, and returns what the Mock object tells it to.
//
// In the real object, this method would do something useful, but since this
// is a mocked object - we're just going to stub it out.
//
// NOTE: This method is not being tested here, code that uses this object is.
func (m *MyMockedObject) DoSomething(number int) (bool, error) {
	args := m.Called(number)
	return args.Bool(0), args.Error(1)
}

// Actual test functions
func TestMock(t *testing.T) {
	// Arrange
	testObj := &MyMockedObject{}

	// Mock
	// testObj.On("DoSomething", 123).Return(true, nil)
	testObj.On("DoSomething", mock.Anything).Return(true, nil)

	// Act
	res := targetFuncThatDoesSomethingWithObj(testObj, 123)

	// Assert
	testObj.AssertExpectations(t)
	assert.True(t, res)
}

func targetFuncThatDoesSomethingWithObj(obj *MyMockedObject, number int) bool {
	res, err := obj.DoSomething(number)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return false
	}
	return res
}
