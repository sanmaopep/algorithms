package stack

import (
	"algorithms/datastructure/utils"
	"errors"
)

type Stacker interface {
	utils.Container
	utils.PushAble
	utils.PopAble
}

type Stack []interface{}

func (stack *Stack) Empty() bool {
	return stack.Size() == 0
}

func (stack *Stack) Size() int {
	return len(*stack)
}

func (stack *Stack) Pop() (interface{},error) {
	var ret interface{}
	var err error
	if stack.Empty() {
		ret = (*stack)[len(*stack)-1]
		*stack = (*stack)[:len(*stack)-1]
	}else{
		err = errors.New("Stack empty!")
	}
	return ret, err
}

func (stack *Stack) Top() (interface{},error) {
	if stack.Empty() {
		return (*stack)[len(*stack)-1], nil
	}else{
		return 0, errors.New("Stack empty!")
	}
}

func (stack *Stack) Push(value interface{}) {
	*stack = append(*stack,value)
}



