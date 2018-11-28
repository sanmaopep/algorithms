package tree

import "algorithms/datastructure/interfaces"

type BinaryTreeNode struct {
	key interfaces.Comparable
	value interface{}
	left *BinaryTreeNode
	right *BinaryTreeNode
}

type BTreeNode struct {

}