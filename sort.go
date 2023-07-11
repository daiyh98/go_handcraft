package main

import (
	"fmt"
	"go_handcraft/tree"
)

type myTreeNode struct { //使用组合来扩展类型
	node *tree.Node
	//tree.Node
}

type newTreeNode struct { //使用内嵌来扩展类型
	*tree.Node
}

func (myNode *myTreeNode) postTraverse(values []int) []int {
	if myNode != nil {
		if myNode.node.Left != nil {
			node := myTreeNode{myNode.node.Left}
			values = node.postTraverse(values) //编译器会自动转成指针
		}
		if myNode.node.Right != nil {
			//treeNode := &myTreeNode{myNode.node.Right}
			values = (&myTreeNode{myNode.node.Right}).postTraverse(values) //此处必须
			// 添加取地址&符号并且加括号，因为结构体字面量不享受go方法调用值或指针自动转换的语法糖！！！
		}
		values = append(values, myNode.node.Value)
	}
	return values
}

func (newNode *newTreeNode) postTraverse(values []int) []int {
	if newNode != nil {
		if newNode.Left != nil {
			values = (&newTreeNode{newNode.Left}).postTraverse(values) //编译器会自动转成指针
		}
		if newNode.Right != nil {
			values = (&newTreeNode{newNode.Right}).postTraverse(values)
		}
		values = append(values, newNode.Value)
	}
	return values
}

func main() {
	values := []int{3, 7, 5, 9, 6, 4}
	//tree.Sort(values)

	//var root = new(myTreeNode)
	//for _, v := range values {
	//	root.node = root.node.Add(v)
	//}
	//sortedValues := root.postTraverse(values[:0])
	//fmt.Println(sortedValues)
	//fmt.Println(values)
	//fmt.Printf("%p\n%p", values, sortedValues)

	var root = new(newTreeNode)
	for _, v := range values {
		root = &newTreeNode{root.Add(v)}
	}
	sortedValues := root.postTraverse(values[:0])
	fmt.Println(sortedValues)
	fmt.Println(values)
	fmt.Printf("%p\n%p", values, sortedValues)
}
