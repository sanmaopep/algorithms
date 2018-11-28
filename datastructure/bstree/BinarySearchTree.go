package bstree

import (
	"algorithms/datastructure/interfaces"
	"errors"
)

type BinarySearchTree struct {
	head Entry
}

func (bst *BinarySearchTree) Search(key interfaces.Comparable) interface{} {
	if entry, ok := bst.head.searchNode(key); ok {
		return entry.value
	} else {
		return nil
	}
}

func (bst *BinarySearchTree) Insert(key interfaces.Comparable, value interface{}) error {
	if entry, ok := bst.head.searchNode(key); ok {
		return errors.New("Key Existed")
	} else {
		cmp := key.Compare(entry.key)
		insertEntry := &Entry{
			key: key,
			value: value,
			parent: entry,
		}

		switch {
		case cmp > 0:
			entry.right = insertEntry
		case cmp < 0:
			entry.left = insertEntry
		}
		return nil
	}
}

func (bst *BinarySearchTree) Delete(key interfaces.Comparable) error {
	if entry, ok := bst.head.searchNode(key); ok {
		switch  {
		case entry.left == nil && entry.right == nil:
			// left and right is nil, just replace it

			// only left is nil, replace left and entry

			// only right is nil, replace right and entry

			// neither left and right is nil, find a top and replace
		}


	} else {
		return errors.New("Key Not Existed")
	}
}

func (bst *BinarySearchTree) Print() {
	panic("implement me")
}
