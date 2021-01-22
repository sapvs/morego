package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type node struct {
	data  int
	left  *node
	right *node
}

func insert(root *node, value int) *node {
	if root == nil {
		return &node{data: value}
	}
	if value <= root.data {
		root.left = insert(root.left, value)
	} else {
		root.right = insert(root.right, value)
	}
	return root
}

func printPretty(root *node) {
	// to print
	h := height(root)

}

func inorder(root *node, data *[]int) {
	if root == nil {
		return
	}

	inorder(root.left, data)
	*data = append(*data, root.data)
	inorder(root.right, data)
}

func preorder(root *node, data *[]int) {
	if root == nil {
		return
	}

	*data = append(*data, root.data)
	preorder(root.left, data)
	preorder(root.right, data)

}

func postorder(root *node, data *[]int) {
	if root == nil {
		return
	}

	postorder(root.left, data)
	postorder(root.right, data)
	*data = append(*data, root.data)
}

func bfs(root *node, data *[]int) {
	if root == nil {
		return
	}
	var que []*node
	que = enque(que[:], root)

	nd, que := deque(que[:])
	for nd != nil {
		*data = append(*data, nd.data)
		que = enque(que[:], nd.left)
		que = enque(que[:], nd.right)
		nd, que = deque(que[:])
	}

}

func enque(que []*node, val *node) []*node {
	if val == nil {
		return que
	}
	return append(que, val)
}

func deque(que []*node) (*node, []*node) {
	if len(que) == 0 {
		return nil, nil
	}
	return que[0], que[1:]
}

func height(root *node) int {
	if root == nil {
		return -1
	}

	hl := height(root.left)
	hr := height(root.right)

	if hl > hr {
		return hl + 1
	}

	return hr + 1
}

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()

	numNodes, _ := strconv.ParseInt(scan.Text(), 10, 32)
	var root *node
	for i := numNodes; i > 0; i-- {
		scan.Scan()
		value, _ := strconv.ParseInt(scan.Text(), 10, 32)
		root = insert(root, int(value))
	}

	fmt.Printf("height %d\n", height(root))

	var data []int
	inorder(root, &data)
	fmt.Printf(" inorder %v \n", data)
	data = nil
	preorder(root, &data)
	fmt.Printf(" preorder %v \n", data)
	data = nil
	postorder(root, &data)
	fmt.Printf(" postorder %v \n", data)
	data = nil
	bfs(root, &data)
	fmt.Printf(" bfs %v \n", data)
	data = nil
}
