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

func (list *LinkedList) Append(data interface{}) {
	newNode := &Node{data: data, next: nil}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
	list.length++
}

func (list *LinkedList) Exists(data interface{}) bool {
	if list.head == nil {
		return false
	}

	current := list.head

	for current != nil {
		if current.data == data {
			return true
		}

		current = current.next
	}

	return false
}
