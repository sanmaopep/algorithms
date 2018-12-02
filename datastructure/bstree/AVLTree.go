package bstree

func balFac(entry *Entry) int{
	return height(entry.left)-height(entry.right)
}

func balance(entry *Entry) {


}

type AVLTree struct {
	head *Entry
}

func (avlTree *AVLTree) Search(key int) interface{} {
	if entry, _, ok := searchNode(&avlTree.head, key); ok {
		return entry.value
	} else {
		return nil
	}
}

func (avlTree *AVLTree) Insert(key int, value interface{}) bool {
	_, ok := insertNode(&avlTree.head, key, value);
	if !ok {
		return ok
	}
	// find unbalanced node

	return ok
}

func (avlTree *AVLTree) Delete(key int) bool {
	panic("implement me")
}

func (avlTree *AVLTree) Print() {
	panic("implement me")
}
