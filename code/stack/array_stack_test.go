package stack

import (
	"testing"
)

func TestNewArrayStack(t *testing.T) {
	stack := newArrayStack()
	if stack.size() != 0 {
		t.Errorf("Expected stack size to be 0, but got %d", stack.size())
	}
	if !stack.isEmpty() {
		t.Error("Expected stack to be empty, but it is not")
	}
}

func TestArrayStackPush(t *testing.T) {
	stack := newArrayStack()
	stack.push(1)
	stack.push(2)
	stack.push(3)

	if stack.size() != 3 {
		t.Errorf("Expected stack size to be 3, but got %d", stack.size())
	}
	if stack.isEmpty() {
		t.Error("Expected stack to not be empty, but it is")
	}
}

func TestArrayStackPop(t *testing.T) {
	stack := newArrayStack()
	stack.push(1)
	stack.push(2)
	stack.push(3)

	val := stack.pop()
	if val != 3 {
		t.Errorf("Expected popped value to be 3, but got %v", val)
	}

	if stack.size() != 2 {
		t.Errorf("Expected stack size to be 2 after pop, but got %d", stack.size())
	}

	val = stack.pop()
	if val != 2 {
		t.Errorf("Expected popped value to be 2, but got %v", val)
	}

	if stack.size() != 1 {
		t.Errorf("Expected stack size to be 1 after pop, but got %d", stack.size())
	}

	stack.pop()

	if stack.size() != 0 {
		t.Errorf("Expected stack size to be 0 after pop, but got %d", stack.size())
	}
	if !stack.isEmpty() {
		t.Error("Expected stack to be empty after pop, but it is not")
	}
}

func TestArrayStackPeek(t *testing.T) {
	stack := newArrayStack()
	stack.push(1)
	stack.push(2)
	stack.push(3)

	val := stack.peek()
	if val != 3 {
		t.Errorf("Expected peeked value to be 3, but got %v", val)
	}

	stack.pop()
	val = stack.peek()
	if val != 2 {
		t.Errorf("Expected peeked value to be 2, but got %v", val)
	}
}

func TestArrayStackToSlice(t *testing.T) {
	stack := newArrayStack()
	stack.push(1)
	stack.push(2)
	stack.push(3)

	slice := stack.toSlice()
	expected := []int{1, 2, 3}
	if len(slice) != len(expected) {
		t.Errorf("Expected slice length to be %d, but got %d", len(expected), len(slice))
	}
	for i := 0; i < len(slice); i++ {
		if slice[i] != expected[i] {
			t.Errorf("Expected slice element at index %d to be %d, but got %d", i, expected[i], slice[i])
		}
	}
}
