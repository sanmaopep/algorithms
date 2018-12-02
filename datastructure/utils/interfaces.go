package utils

type Container interface {
	Size() int
	Empty() bool
}

type PushAble interface {
	Push(value int)
}

type PopAble interface {
	Pop() int
	Top() int
}

type Comparable interface {
	Compare(comparable interface{}) int
}

type Set interface {
	Search (key int) interface{}
	Insert (key int, value interface{}) error
	Delete (key int) error
	Print ()
}
