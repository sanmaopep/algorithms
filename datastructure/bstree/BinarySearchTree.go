package bstree

type BinarySearchTree struct {
	head *Entry
}

func (bst *BinarySearchTree) Search(key int) interface{} {
	if entry, _, ok := searchNode(&bst.head, key); ok {
		return entry.value
	} else {
		return nil
	}
}

func (bst *BinarySearchTree) Insert(key int, value interface{}) bool {
	_, ok := insertNode(&bst.head, key, value);
	return ok
}

func (bst *BinarySearchTree) Delete(key int) bool {
	_, ok := deleteNode(&bst.head, key);
	return ok
}

func (bst *BinarySearchTree) Print() {
	bst.head.print()
}
