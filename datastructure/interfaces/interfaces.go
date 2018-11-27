package interfaces

type Container interface {
	size() int
}

type Pushable interface {
	push(value int)
}

type Popable interface {
	pop() int
	top() int
}
