package stack

/*基于切片实现栈*/
type arrayStack struct {
	data []int
}

/*初始化栈*/
func newArrayStack() *arrayStack {
	return &arrayStack{
		// 设置栈长度为0，容量为16
		data: make([]int, 0, 16),
	}
}

/*栈的长度*/
func (s *arrayStack) size() int {
	return len(s.data)
}

/*栈是否为空*/
func (s *arrayStack) isEmpty() bool {
	return len(s.data) == 0
}

/*入栈*/
func (s *arrayStack) push(val int) {
	// 切片会自动扩容
	s.data = append(s.data, val)
}

/*出栈*/
func (s *arrayStack) pop() any {
	if s.isEmpty() {
		return nil
	}
	val := s.peek()
	s.data = s.data[:len(s.data)-1]
	return val
}

/*获取栈顶元素*/
func (s *arrayStack) peek() any {
	if s.isEmpty() {
		return nil
	}
	val := s.data[len(s.data)-1]
	return val
}

/*获取slice用于打印*/
func (s *arrayStack) toSlice() []int {
	return s.data
}
