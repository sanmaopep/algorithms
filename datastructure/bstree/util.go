package bstree

import "algorithms/datastructure/interfaces"


type Entry struct {
	key   interfaces.Comparable
	attribute interface{}
	value interface{}
	left  *Entry
	right *Entry
	parent *Entry
}

func (entry *Entry) searchNode(key interfaces.Comparable) (curr *Entry,result bool) {
	_hot := entry

	for {
		cmp := key.Compare(_hot.key)
		switch {
		case cmp == 0:
			return _hot, true
		case  cmp < 0:
			if _hot.left == nil {
				break
			}
			_hot = _hot.left
		default:
			if _hot.right == nil {
				break
			}
			_hot = _hot.right
		}
	}
	return _hot,false
}

func (entry *Entry) print(){

}

// rotate to right
func zig(entry *Entry){
	if entry.left == nil {
		panic("left not nil required in zig operation")
	}
	left := entry.left
	// left -> parent
	if entry.parent.left == entry {
		entry.parent.left = left
	} else {
		entry.parent.right = left
	}

	// entry -> left.right
	left.right = entry
	entry.parent = left

	// left.right -> entry.left
	entry.left = left.right
}

// rotate to left
func zag(entry *Entry) {
	if entry.right == nil {
		panic("right not nil required in zag operation")
	}

	right := entry.right
	// right -> parent
	if entry.parent.left == entry {
		entry.parent.left = right
	} else {
		entry.parent.right = right
	}

	// entry -> right.left
	right.left = entry
	entry.parent = right

	// right.left -> entry.right
	entry.right = right.left
}