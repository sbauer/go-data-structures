package lists

type DoublyNode struct {
	next     *DoublyNode
	previous *DoublyNode
	data     interface{}
}

type DoublyLinkedList struct {
	head   *DoublyNode
	tail   *DoublyNode
	length int
}

func (list *DoublyLinkedList) Insert(data interface{}) {
	node := &DoublyNode{data: data}

	if list.head == nil {
		list.head = node
		list.tail = node
		list.incrementLength()
		return
	}

	node.next = list.head
	list.head = node
	list.incrementLength()
}

func (list *DoublyLinkedList) Append(data interface{}) {
	node := &DoublyNode{data: data}

	if list.head == nil {
		list.head = node
		list.tail = node
		list.incrementLength()
		return
	}

	list.tail.next = node
	node.previous = list.tail
	list.tail = node
	list.incrementLength()
}

func (list *DoublyLinkedList) Empty() bool {
	return list.head == nil
}

func (list *DoublyLinkedList) incrementLength() {
	list.length++
}

func (list *DoublyLinkedList) decrementLength() {
	list.length++
}
