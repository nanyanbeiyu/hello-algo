package double_ended_queue

import "container/list"

/*
对于双向队列而言，头部和尾部都可以执行入队和出队操作。换句话说，双向队列需要实现另一个对称方向的操作。为此，我们采用“双向链表”作为双向队列的底层数据结构。
*/
/* 基于双向链表实现的双向队列 */
type linkedListDeque struct {
	// 使用内置包 list
	data *list.List
}

/* 初始化双端队列 */
func newLinkedListDeque() *linkedListDeque {
	return &linkedListDeque{
		data: list.New(),
	}
}

/* 队首元素入队 */
func (s *linkedListDeque) pushFirst(val any) {
	s.data.PushFront(val)
}

/*队尾元素入队*/
func (s *linkedListDeque) pushLast(val any) {
	s.data.PushBack(val)
}

/*队首元素出队*/
func (s *linkedListDeque) popFirst() any {
	if s.isEmpty() {
		return nil
	}
	e := s.data.Front()
	s.data.Remove(e)
	return e
}

/*队尾元素出队*/
func (s *linkedListDeque) popLast() any {
	if s.isEmpty() {
		return nil
	}
	e := s.data.Back()
	s.data.Remove(e)
	return e
}

/*访问队首元素*/
func (s *linkedListDeque) peekFirst() any {
	if s.isEmpty() {
		return nil
	}
	e := s.data.Front()
	return e.Value
}

/*访问队尾元素*/
func (s *linkedListDeque) peekLast() any {
	if s.isEmpty() {
		return nil
	}
	e := s.data.Back()
	return e.Value
}

/*获取队列长度*/
func (s *linkedListDeque) size() int {
	return s.data.Len()
}

/*判断队列是否为空*/
func (s *linkedListDeque) isEmpty() bool {
	return s.data.Len() == 0
}

/* 获取 List 用于打印 */
func (s *linkedListDeque) toList() *list.List {
	return s.data
}
