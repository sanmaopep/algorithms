package interfaces

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
	Compare(Comparable) int
}

type Set interface {
	Search (key Comparable) interface{}
	Insert (key Comparable, value interface{}) error
	Delete (key Comparable) error
	Print ()
}
