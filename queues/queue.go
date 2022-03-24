package queues

type Queue struct {
	data []interface{}
}

func NewQueue() *Queue {
	return &Queue{data: make([]interface{}, 0)}
}

func (queue *Queue) Enqueue(item interface{}) {
	queue.data = append(queue.data, item)
}

func (queue *Queue) Dequeue() interface{} {
	if queue.Empty() {
		return nil
	}

	item := queue.data[0]
	queue.data = queue.data[1:]

	return item
}

func (queue *Queue) Peek() interface{} {
	if queue.Empty() {
		return nil
	}

	return queue.data[0]
}

func (queue *Queue) Length() int {
	return len(queue.data)
}

func (queue *Queue) Empty() bool {
	return queue.Length() == 0
}
