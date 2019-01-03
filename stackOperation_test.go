package stackoperation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	var maStack *myStackType
	maStack = new(myStackType)
	maStack.Push(10)
	maStack.Push(20)
	maStack.Push(5)
	maStack.ShowMyStack()
	maStack.Pull()
	maStack.ShowMyStack()
	maStack.Pull()
	maStack.ShowMyStack()
	assert.Equal(t, 10, *maStack.Pull())
}

func TestGetPeek(t *testing.T) {
	var maStack *myStackType
	maStack = new(myStackType)
	maStack.Push(10)
	fmt.Println(" 1111 ", maStack)
	maStack.ShowMyStack()
	assert.Equal(t, 10, *maStack.GetPeek())
}

func TestGetPeekEmpty(t *testing.T) {
	var maStack myStackType
	var exp *int
	assert.Equal(t, exp, maStack.GetPeek())
}

func TestEmptyPull(t *testing.T) {
	var maStack *myStackType
	var exp *int
	maStack = new(myStackType)
	maStack.ShowMyStack()
	assert.Equal(t, exp, maStack.Pull())
}
