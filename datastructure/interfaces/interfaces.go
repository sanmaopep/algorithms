package interfaces

type Container interface {
	Size() int
	Empty() bool
}

type Pushable interface {
	Push(value int)
}

type Popable interface {
	Pop() int
	Top() int
}

type Comparable interface {
	GreaterThan(Comparable) bool
	LessThan(Comparable) bool
}

type Set interface {
	Search (key Comparable) bool
	Insert (key Comparable)
	Delete (key Comparable)
	Print ()
}
