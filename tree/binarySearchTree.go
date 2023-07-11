package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

// 按照左小右大的原则添加新叶子节点
func (t *Node) Add(value int) *Node {
	if t == nil {
		t = new(Node)
		t.Value = value
		return t
	}
	if t.Value > value {
		t.Left = t.Left.Add(value)
	} else {
		t.Right = t.Right.Add(value)
	}
	return t
}

func Sort(values []int) {
	var root *Node
	for _, v := range values {
		root = root.Add(v)
	}
	sortedValues := root.PreTraverse(values[:0])
	fmt.Println(sortedValues)
	fmt.Println(values)
	fmt.Printf("%p\n%p", values, sortedValues)
}
