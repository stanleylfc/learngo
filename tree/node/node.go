package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node *Node) Print() {
	fmt.Println(node.Value)
}

func (node *Node) SetValue(value int) {
	node.Value = value
}

func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse() //node.Left = nil 也可以调用
	node.Print()
	node.Right.Traverse()
}

func CreateNode(value int) *Node {
	return &Node{Value: value} //局部变量
}
