package tree

// MiddleTraverse 通过中序遍历二叉树，将二叉树中的数据加入切片中
// 先往左边回溯，一直到没有左节点，再放中间节点，再放右边节点
func (t *Node) MiddleTraverse(values []int) []int {
	if t != nil {
		values = t.Left.MiddleTraverse(values)
		values = append(values, t.Value)
		values = t.Right.MiddleTraverse(values)
	}
	return values
}

func (t *Node) PreTraverse(values []int) []int {
	if t != nil {
		values = append(values, t.Value)
		values = t.Left.PreTraverse(values)
		values = t.Right.PreTraverse(values)
	}
	return values
}
