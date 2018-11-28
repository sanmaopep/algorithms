package bstree

import (
	"algorithms/datastructure/interfaces"
	"errors"
)

type avlAttribute struct {
	balance int // left_height - right_height
}

func updateBalance(curr *Entry) {

	for curr != nil {
		switch {
		case curr.left == nil && curr.right == nil:
			curr.attribute.(avlAttribute).balance = 0
		case curr.left == nil:
			curr.attribute.(avlAttribute).balance =
				-curr.right.attribute.(avlAttribute).balance - 1
		case curr.right == nil:
			curr.attribute.(avlAttribute).balance =
				curr.left.attribute.(avlAttribute).balance + 1
		default:
			curr.attribute.(avlAttribute).balance =
				curr.left.attribute.(avlAttribute).balance - curr.right.
					attribute.(avlAttribute).balance
		}
		curr = curr.parent
	}

}

type AVLTree struct {
	head Entry
}

func (avlTree *AVLTree) Search(key interfaces.Comparable) interface{} {
	if entry, ok := avlTree.head.searchNode(key); ok {
		return entry.value
	} else {
		return nil
	}
}

func (avlTree *AVLTree) Insert(key interfaces.Comparable, value interface{}) error {
	if entry, ok := avlTree.head.searchNode(key); ok {
		return errors.New("Key Existed")
	} else {
		cmp := key.Compare(entry.key)
		insertEntry := &Entry{
			key:    key,
			value:  value,
			parent: entry,
		}

		switch {
		case cmp > 0:
			entry.right = insertEntry
		case cmp < 0:
			entry.left = insertEntry
		}

		// update balance

		return nil
	}
}

func (avlTree *AVLTree) Delete(key interfaces.Comparable) error {
	panic("implement me")
}

func (avlTree *AVLTree) Print() {
	panic("implement me")
}
