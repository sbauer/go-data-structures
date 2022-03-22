package stacks

type Stack struct {
	data []interface{}
}

func NewStack() *Stack {
	return &Stack{data: make([]interface{}, 0)}
}

func (stack *Stack) Push(item interface{}) {
	stack.data = append(stack.data, item)
}

func (stack *Stack) Pop() interface{} {
	length := stack.Length()

	if length < 1 {
		return nil
	}

	item := stack.data[length-1]
	stack.data = stack.data[:length-1]

	return item
}

func (stack *Stack) Length() int {
	return len(stack.data)
}

func (stack *Stack) Empty() bool {
	return stack.Length() == 0
}
