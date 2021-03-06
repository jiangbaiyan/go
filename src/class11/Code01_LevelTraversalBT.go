package class11

import (
	"fmt"
	"github.com/jiangbaiyan/go/src/class03"
)

type Node struct {
	value interface{}
	left  *Node
	right *Node
}

// LevelTraversalBT ε±εΊιε
func LevelTraversalBT(head *Node) {
	if head == nil {
		return
	}
	queue := &class03.MyRingQueue{}
	queue.Push(head)
	for !queue.IsEmpty() {
		cur := queue.Pop().(*Node)
		fmt.Println(cur.value)
		if cur.left != nil {
			queue.Push(cur.left)
		}
		if cur.right != nil {
			queue.Push(cur.right)
		}
	}
}
