package tree

import (
	"algorithms/datastructure/interfaces"
	"algorithms/datastructure/stack"
	"errors"
)

type BinarySearchTree struct {
	head BinaryTreeNode
}

func (bst *BinarySearchTree) SearchNode(key interfaces.Comparable) (BinaryTreeNode,error) {
	var stack stack.Stack
	var _hot BinaryTreeNode
	stack.Push(bst.head)

	for !stack.Empty() {
		curr, _ := stack.Pop()
		_hot = curr.(BinaryTreeNode)

		if _hot.key == key {
			// check if curr key is matched
			return _hot, nil
		}else if key.LessThan(_hot.key ) {
			// search left subtree
			if _hot.left == nil {
				break
			}
			stack.Push(_hot.left)
		}else {
			// search right subtree
			if _hot.right == nil {
				break
			}
			stack.Push(_hot.right)
		}
	}
	return _hot, errors.New("Not Found")
}

func (bst *BinarySearchTree) Insert(key interface{}) {

	panic("implement me")
}

func (bst *BinarySearchTree) Delete(key interface{}) {
	panic("implement me")
}

func (bst *BinarySearchTree) Print() {
	panic("implement me")
}


