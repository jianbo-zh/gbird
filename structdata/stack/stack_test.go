package stack

import "testing"

var stack ItemStack

func initStack() *ItemStack {
	if stack.items == nil {
		stack = ItemStack{}
		stack.New()
	}

	return &stack
}

func TestPush(t *testing.T) {
	stack := initStack()
	stack.Push(1)

	top := stack.Top()
	if top != 1 {
		t.Errorf("Wrong stack top item, expected 1 and go %d", top)
	}

	stack.Push(2)
	stack.Push(3)

	if size := len(stack.items); size != 3 {
		t.Errorf("Wrong stack size, expected 3 and got %d", size)
	}
}

func TestPull(t *testing.T) {
	stack.Pop()
	if size := len(stack.items); size != 2 {
		t.Errorf("Wrong stack size, expected 2 and got %d", size)
	}

	stack.Pop()
	stack.Pop()

	if !stack.IsEmpty() {
		t.Errorf("Wrong stack size, expected empty")
	}
}
