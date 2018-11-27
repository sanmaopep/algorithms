package stack

type Stack []int

func (stack *Stack) size() int {
	return len(*stack)
}

func (stack *Stack) pop() int {
	var ret int
	if stack.size() != 0 {
		ret = (*stack)[len(*stack)-1]
		*stack = (*stack)[:len(*stack)-1]
	}
	return ret
}

func (stack *Stack) top() int {
	return (*stack)[len(*stack)-1]
}

func (stack *Stack) push(value int) {
	*stack = append(*stack,value)
}



