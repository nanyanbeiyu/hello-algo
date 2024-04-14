package stack

import "container/list"

/*基于链表实现栈*/
type linkedListStack struct {
	// 使用内置的链表
	data *list.List
}

/*初始化栈*/
func newLinkedListStack() *linkedListStack {
	return &linkedListStack{
		data: list.New(),
	}
}

/*入栈*/
func (s *linkedListStack) push(val int) {
	s.data.PushBack(val)
}

/*出栈*/
func (s *linkedListStack) pop() any {
	if s.isEmpty() {
		return nil
	}
	e := s.data.Back()
	s.data.Remove(e)
	return e.Value
}

/*访问栈顶元素*/
func (s *linkedListStack) peek() any {
	if s.isEmpty() {
		return nil
	}
	return s.data.Back().Value
}

/*获取栈长度*/
func (s *linkedListStack) size() int {
	return s.data.Len()
}

/*判断是否为空*/
func (s *linkedListStack) isEmpty() bool {
	return s.data.Len() == 0
}

/*获取list打印*/
func (s *linkedListStack) toList() *list.List {
	return s.data
}
