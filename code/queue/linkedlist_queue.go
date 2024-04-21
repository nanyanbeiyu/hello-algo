package queue

import "container/list"

/*
基于链表实现队列
我们可以将链表的“头节点”和“尾节点”分别视为“队首”和“队尾”，规定队尾仅可添加节点，队首仅可删除节点。
*/
type linkedListQueue struct {
	data *list.List
}

/*初始化队列*/
func newLinkedListQueue() *linkedListQueue {
	return &linkedListQueue{
		data: list.New(),
	}
}

/*入队*/
func (s *linkedListQueue) push(val any) {
	s.data.PushBack(val)
}

/*出队*/
func (s *linkedListQueue) pop() any {
	if s.isEmpty() {
		return nil
	}
	e := s.data.Front()
	s.data.Remove(e)
	return e.Value
}

/*访问队首元素*/
func (s *linkedListQueue) peek() any {
	if s.isEmpty() {
		return nil
	}
	e := s.data.Front()
	return e.Value
}

/*获取队列长度*/
func (s *linkedListQueue) size() any {
	return s.data.Len()
}

/*判断队列是否为空*/
func (s *linkedListQueue) isEmpty() bool {
	return s.data.Len() == 0
}

/*获取list用于打印*/
func (s *linkedListQueue) toList() *list.List {
	return s.data
}
