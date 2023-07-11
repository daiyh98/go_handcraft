package queue

type Queue []interface{}

func (q *Queue) Push(value int) {
	*q = append(*q, value)
}

func (q *Queue) Pop() int {
	head := (*q)[0] //head目前是interface，返回时需要断言为int
	*q = (*q)[1:]
	return head.(int)
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
