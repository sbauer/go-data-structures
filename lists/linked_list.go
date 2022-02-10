package lists

type Node struct {
	next *Node
	data interface{}
}

type LinkedList struct {
	head   *Node
	length int
}

func (list *LinkedList) Insert(data interface{}) {
	newNode := &Node{data: data, next: list.head}

	list.head = newNode
	list.length++

}
