package main

import (
	"fmt"
	"github.com/xlab/treeprint"
)

type node struct {
	left  *node
	right *node
	data  int64
}

func (n *node) insert(data int64) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &node{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &node{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}

type binaryTree struct {
	root *node
}

func (t *binaryTree) insert(data int64) {
	if t.root == nil {
		t.root = &node{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
}

func print(node *node, space int, ch string) {
	if node == nil {
		return
	}

	for i := 0; i < space; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%s:%v\n", ch, node.data)
	print(node.left, space+2, "L")
	print(node.right, space+2, "R")
}

func treePrint(node *node, tp treeprint.Tree) {
	if node == nil {
		return
	}
	one := tp.AddBranch(node.data)
	treePrint(node.left, one)
	treePrint(node.right, one)
}

func main() {
	tree := &binaryTree{}
	tree.insert(100)
	tree.insert(20)
	tree.insert(120)
	tree.insert(50)
	tree.insert(15)
	tree.insert(85)
	tp := treeprint.New()
	treePrint(tree.root, tp)
	fmt.Println(tp.String())
}
