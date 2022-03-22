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

func (list *DoublyLinkedList) Contains(data interface{}) bool {
	return list.find(data) != nil
}

func (list *DoublyLinkedList) Find(data interface{}) *DoublyNode {
	return list.find(data)
}

func (list *DoublyLinkedList) find(data interface{}) *DoublyNode {
	current := list.head

	for current != nil {
		if current.data == data {
			return current
		}

		current = current.next
	}

	return nil
}

func (list *DoublyLinkedList) incrementLength() {
	list.length++
}
