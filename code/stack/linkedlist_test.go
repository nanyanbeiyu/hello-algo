package stack

import (
	"testing"
)

func TestNewLinkedListStack(t *testing.T) {
	stack := newLinkedListStack()
	if stack.data == nil {
		t.Errorf("expected non-nil data, got nil")
	}
}

func TestPush(t *testing.T) {
	stack := newLinkedListStack()
	stack.push(1)
	if stack.size() != 1 {
		t.Errorf("expected size 1, got %d", stack.size())
	}
}

func TestPop(t *testing.T) {
	stack := newLinkedListStack()
	stack.push(1)
	value := stack.pop()
	if value != 1 {
		t.Errorf("expected value 1, got %v", value)
	}
	if stack.size() != 0 {
		t.Errorf("expected size 0, got %d", stack.size())
	}
}

func TestPeek(t *testing.T) {
	stack := newLinkedListStack()
	stack.push(1)
	value := stack.peek()
	if value != 1 {
		t.Errorf("expected value 1, got %v", value)
	}
	if stack.size() != 1 {
		t.Errorf("expected size 1, got %d", stack.size())
	}
}

func TestSize(t *testing.T) {
	stack := newLinkedListStack()
	if stack.size() != 0 {
		t.Errorf("expected size 0, got %d", stack.size())
	}
	stack.push(1)
	stack.push(2)
	if stack.size() != 2 {
		t.Errorf("expected size 2, got %d", stack.size())
	}
}

func TestIsEmpty(t *testing.T) {
	stack := newLinkedListStack()
	if !stack.isEmpty() {
		t.Errorf("expected empty stack, got non-empty stack")
	}
	stack.push(1)
	if stack.isEmpty() {
		t.Errorf("expected non-empty stack, got empty stack")
	}
}

func TestToList(t *testing.T) {
	stack := newLinkedListStack()
	stack.push(1)
	stack.push(2)
	stack.push(3)
	list := stack.toList()
	if list.Len() != 3 {
		t.Errorf("expected list length 3, got %d", list.Len())
	}
}
