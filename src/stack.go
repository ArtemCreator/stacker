package src

import "errors"

type Stack []interface{}

func (stack Stack) Len() int {
	return len(stack)
}

func (stack Stack) Cap() int {
	return cap(stack)
}

func (stack Stack) isEmpty() bool {
	return len(stack) == 0
}

func (stack *Stack) Push(x interface{})  {
	*stack = append(*stack, x)
}

func (stack *Stack) Pop() (interface{}, error)  {
	theStack := *stack
	if len(theStack) == 0{
		return nil, errors.New("can't Pop() is a empty stack")
	}
	x := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return x, nil
}

func (stack Stack) Top() (interface{}, error)  {
	if len(stack) == 0 {
		return nil, errors.New("can't Top() is a empty stack")
	}

	return stack[len(stack) - 1], nil
}