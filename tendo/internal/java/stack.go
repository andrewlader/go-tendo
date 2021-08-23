package java

type currentState uint8

const (
	withinNothing currentState = iota
	withinClass
	withinMethodOrFunction
)

type stackWalkState struct {
	stack []currentState
}

func (stack *stackWalkState) push(value currentState) {
	stack.stack = append(stack.stack, value)
}

func (stack *stackWalkState) pop() currentState {
	// save the top most element
	topElement := stack.peek()

	lenStack := len(stack.stack)
	if lenStack > 0 {
		// pop off the top most element of the stack
		stack.stack = stack.stack[:(lenStack - 1)]
	}

	return topElement
}

func (stack *stackWalkState) peek() currentState {
	lenStack := len(stack.stack)
	if lenStack == 0 {
		return withinNothing
	}

	return stack.stack[(lenStack - 1)]
}
