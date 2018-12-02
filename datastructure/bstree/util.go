package bstree

import (
	"fmt"
)

type Entry struct {
	key       int
	attribute interface{}
	value     interface{}
	left      *Entry
	right     *Entry
	parent    *Entry
	height    int
}

// _hot is a special pointer
func searchNode(entry **Entry, key int) (curr *Entry, _hot **Entry, result bool) {
	// wrong here, &entry is new!
	_hot = entry
	parent := (*entry).parent

	for *_hot != nil {
		cmp := key - (*_hot).key
		parent = *_hot
		switch {
		case cmp == 0:
			return *_hot, _hot, true
		case cmp < 0:
			_hot = &((*_hot).left)
		default:
			_hot = &((*_hot).right)
		}
	}
	return parent, _hot, false
}

func insertNode(entry **Entry, key int, value interface{}) (insert *Entry, result bool) {
	if *entry == nil {
		*entry = &Entry{
			key:    key,
			value:  value,
			height: 1,
		}
	}

	if parent, hot, ok := searchNode(entry, key); ok {
		return nil, false
	} else {
		insertEntry := &Entry{
			key:    key,
			value:  value,
			parent: parent,
			height: 1,
		}

		*hot = insertEntry
		updateHeight(insertEntry)

		return insertEntry, true
	}
}

func deleteNode(entry **Entry, key int) (succ *Entry, result bool) {
	if delEntry, hot, ok := searchNode(entry, key); ok {
		var succ, oldSuccParent *Entry // successor
		parent := delEntry.parent
		switch {
		case delEntry.left == nil:
			succ = delEntry.right
		case delEntry.right == nil:
			succ = delEntry.left
		default:
			//  neither left and right is nil
			succ = delEntry.right
			for succ.left != nil {
				succ = succ.left
			}
			oldSuccParent = succ.parent
			succ.parent = parent
			delEntry.left.parent = succ
			delEntry.right.parent = succ
		}

		*hot = succ

		updateHeight(oldSuccParent)
	}
	return nil, false
}

func updateHeight(entry *Entry){
	curr := entry
	for curr != nil {
		curr.height = height(tallerChild(curr)) + 1
		curr = curr.parent
	}
}

func height(entry *Entry) int{
	if entry == nil {
		return 0
	}else{
		return entry.height
	}
}

func tallerChild(entry *Entry) *Entry{
	if height(entry.left) > height(entry.right) {
		return entry.left
	}else{
		return entry.right
	}
}

func (entry *Entry) print() {
	if entry == nil {
		return
	}
	fmt.Printf("%d(h:%d,v:%d) ", entry.key, entry.height, entry.value)
	entry.left.print()
	entry.right.print()
	if entry.parent == nil {
		fmt.Print("\n")
	}
}
