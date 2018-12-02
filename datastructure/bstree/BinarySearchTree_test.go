package bstree

import (
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	bst := BinarySearchTree{}

	bst.Insert(2,2)
	bst.Insert(1,1)
	bst.Insert(4,4)
	bst.Insert(3,3)
	bst.Insert(5,5)
	bst.Insert(6,6)
	bst.Insert(7,7)
	bst.Insert(8,8)
	bst.Print()

	if bst.Search(2) != 2 {
		t.Fail()
	}

	bst.Delete(2)
	bst.Print()

	if bst.Search(2) == 2 {
		t.Fail()
	}
}