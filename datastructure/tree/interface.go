package tree

type BalancedBinarySearchTree interface {
	Search (value int) bool
	Insert (value int)
	Delete (value int)
	Print ()
}
