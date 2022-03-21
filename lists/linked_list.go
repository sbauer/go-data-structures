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
	list.incrementLength()
}

func (list *LinkedList) Append(data interface{}) {
	newNode := &Node{data: data, next: nil}

	if list.head == nil {
		list.head = newNode
		list.incrementLength()
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
	list.incrementLength()
}

func (list *LinkedList) Contains(data interface{}) bool {
	result := list.find(data)

	return result != nil
}

func (list *LinkedList) Find(data interface{}) *Node {
	return list.find(data)
}

func (list *LinkedList) find(data interface{}) *Node {
	current := list.head

	for current != nil {
		if current.data == data {
			return current
		}

		current = current.next
	}

	return nil
}

func (list *LinkedList) Remove(node *Node) bool {
	if node == nil {
		return false
	}

	if list.head == nil {
		return false
	}

	if list.head == node {
		list.head = list.head.next
		node.next = nil
		list.decrementLength()
		return true
	}

	current := list.head

	for current.next != nil {
		if current.next == node {
			current.next = node.next
			list.decrementLength()
			return true
		}

		current = current.next
	}

	return false
}

func (list *LinkedList) RemoveByValue(data interface{}) bool {
	if list.head == nil {
		return false
	}

	if list.head.data == data {
		list.head = list.head.next
		list.decrementLength()
		return true
	}

	current := list.head

	for current.next != nil {
		if current.next.data == data {
			current.next = current.next.next
			list.decrementLength()
			return true
		}

		current = current.next
	}

	return false
}

func (list *LinkedList) Length() int {
	return list.length
}

func (list *LinkedList) decrementLength() {
	list.length--
}

func (list *LinkedList) incrementLength() {
	list.length++
}
