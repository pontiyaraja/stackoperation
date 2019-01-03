package stackoperation

import "fmt"

type myStackType []int

func (myStack *myStackType) Push(data int) {
	*myStack = append(*myStack, data)
	fmt.Println(myStack)
}

func (myStack *myStackType) Pull() *int {
	var popData *int
	if len(*myStack) > 0 {
		popData = &myStack.getStackArray()[myStack.getLength()-1]
		fmt.Println(*popData)
	}
	*myStack = myStack.RearrangemyStack()
	return popData
}

func (myStack *myStackType) RearrangemyStack() myStackType {
	if myStack.getLength() <= 1 {
		myStack = new(myStackType)
		return *myStack
	}
	*myStack = myStack.getStackArray()[0 : myStack.getLength()-1]
	return *myStack
}

func (myStack *myStackType) ShowMyStack() {
	fmt.Println(*myStack)
}

func (myStack *myStackType) getLength() int {
	return len(*myStack)
}

func (myStack *myStackType) getStackArray() myStackType {
	return *myStack
}

func (myStack *myStackType) GetPeek() *int {
	if myStack.getLength() == 0 {
		return nil
	}
	return &myStack.getStackArray()[myStack.getLength()-1]
}
