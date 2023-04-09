package main

import (
	"fmt"
	"math"
	"strings"
)

type Tree struct {
	root *Node
}

type Node struct {
	parent *Node
	left   *Node
	right  *Node
	color  string //0 - черный, 1 - красный
	value  int
}

type LeveledNode struct {
	node  *Node
	level int
}

func (node *Node) getUncle() *Node {
	scan := node.parent
	if scan.parent == nil {
		return nil
	}
	if scan.parent.right == scan {
		return scan.parent.left
	} else {
		return scan.parent.right
	}
}

func (node *Node) getColor() string {
	if node == nil {
		return "BLACK"
	} else {
		return node.color
	}
}

func (node *Node) leftTurn() {
	pivot := node.right
	pivot.parent = node.parent
	if node.parent != nil {
		if node.parent.left == node {
			node.parent.left = pivot
		} else {
			node.parent.right = pivot
		}
	}
	node.right = pivot.left
	if pivot.left != nil {
		pivot.left.parent = node
	}
	node.parent = pivot
	pivot.left = node
}

func (node *Node) rightTurn() {
	pivot := node.left
	pivot.parent = node.parent
	if node.parent != nil {
		if node.parent.left == node {
			node.parent.left = pivot
		} else {
			node.parent.right = pivot
		}
	}
	node.left = pivot.right
	if pivot.right != nil {
		pivot.right.parent = node
	}
	node.parent = pivot
	pivot.right = node
}

func getHeight(node *Node) int {
	if node == nil {
		return 0
	}
	leftHeight := getHeight(node.left) + 1
	rightHeight := getHeight(node.right) + 1
	if leftHeight > rightHeight {
		return leftHeight
	} else {
		return rightHeight
	}
}

func (tree *Tree) balanceInsertion(node *Node) {
	if tree.root == node {
		node.color = "BLACK"
	}

	if node.parent.getColor() == "RED" {
		uncle := node.getUncle()
		if uncle != nil && uncle.color == "RED" {
			grandfather := node.parent.parent
			node.parent.color = "BLACK"
			uncle.color = "BLACK"
			grandfather.color = "RED"
			tree.balanceInsertion(grandfather)
		} else {
			grandfather := node.parent.parent
			if node == node.parent.right && node.parent == grandfather.left {
				node.parent.leftTurn()
				node = node.left
			} else if node == node.parent.left && node.parent == grandfather.right {
				node.parent.rightTurn()
				node = node.right
			}

			node.parent.color = "BLACK"
			grandfather.color = "RED"
			if node == node.parent.right && node.parent == grandfather.right {
				grandfather.leftTurn()
			} else if node == node.parent.left && node.parent == grandfather.left {
				grandfather.rightTurn()
			}
		}
	}
}

func (tree *Tree) insert(value int) {
	if tree.root == nil {
		tree.root = &Node{
			parent: nil,
			left:   nil,
			right:  nil,
			color:  "BLACK",
			value:  value,
		}
	} else {
		scan := tree.root
		var scan_parent *Node = nil
		for scan != nil {
			scan_parent = scan
			if value <= scan.value {
				scan = scan.left
			} else {
				scan = scan.right
			}
		}
		var newNode *Node
		if value <= scan_parent.value {
			newNode = &Node{
				parent: scan_parent,
				left:   nil,
				right:  nil,
				color:  "RED",
				value:  value,
			}
			scan_parent.left = newNode
		} else {
			newNode = &Node{
				parent: scan_parent,
				left:   nil,
				right:  nil,
				color:  "RED",
				value:  value,
			}
			scan_parent.right = newNode
		}
		tree.balanceInsertion(newNode)
	}
}

func getTreeSlice(node *Node, nodeSlice *[]*LeveledNode, level int) {
	*nodeSlice = append(*nodeSlice, &LeveledNode{node, level})
	if node != nil {
		getTreeSlice(node.left, nodeSlice, level+1)
		getTreeSlice(node.right, nodeSlice, level+1)
	}
}

func (tree *Tree) printTree() {
	colorReset := "\033[0m"
	colorRed := "\033[31m"

	var nodeSlice []*LeveledNode
	treeHeight := getHeight(tree.root)
	getTreeSlice(tree.root, &nodeSlice, 0)
	for i := 0; i < treeHeight; i++ {
		spaceNum := math.Pow(2, float64(treeHeight-i-1)) - 1
		spacing := strings.Repeat("  ", int(spaceNum))
		for j := 0; j < len(nodeSlice); j++ {
			if nodeSlice[j].node != nil {
				if nodeSlice[j].level == i {
					if nodeSlice[j].node.color == "RED" {
						fmt.Printf("%s%s%02d%s", colorRed, spacing, nodeSlice[j].node.value, spacing)
					} else {
						fmt.Printf("%s%s%02d%s", colorReset, spacing, nodeSlice[j].node.value, spacing)
					}

					fmt.Print("  ")
				}
			} else if nodeSlice[j].level == i {
				fmt.Print("  ")
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
