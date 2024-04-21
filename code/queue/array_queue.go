package queue

/* 基于环形数组实现的队列 */

type arrayQueue struct {
	nums        []int // 用于存储队列元素的数组
	front       int   // 队首指针，指向队首元素
	queSize     int   // 队列长度
	queCapacity int   //队列容量（即最大容纳元素数量）
}

func newArrayQueue(queCapacity int) *arrayQueue {
	return &arrayQueue{
		nums:        make([]int, queCapacity),
		front:       0,
		queSize:     0,
		queCapacity: queCapacity,
	}
}

/* 获取队列长度*/
func (s *arrayQueue) size() int {
	return s.queSize
}

/* 判断队列是否为空*/
func (s *arrayQueue) isEmpty() bool {
	return s.queSize == 0
}

/*入队*/
func (s *arrayQueue) push(val int) {
	// 当 rear == queCapacity 表示队列已满
	if s.queSize == s.queCapacity {
		return
	}
	/// 计算队尾指针，指向队尾索引 + 1
	// 通过取余操作实现 rear 越过数组尾部后回到头部
	rear := (s.front + s.queSize) % s.queCapacity
	// 将val添加在队尾
	s.nums[rear] = val
	s.queSize++
}

/*出队*/
func (s *arrayQueue) pop() int {
	num := s.peek()
	// 队首指针向后移动一位，若越过尾部，则返回到数组头部
	s.front = (s.front + 1) % s.queCapacity
	s.queSize--
	return num
}

/*查询队首元素*/
func (s *arrayQueue) peek() int {
	if s.isEmpty() {
		return -1
	}
	return s.nums[s.front]
}

/* 获取 Slice 用于打印 */
func (q *arrayQueue) toSlice() []int {
	rear := (q.front + q.queSize)
	if rear >= q.queCapacity {
		rear %= q.queCapacity
		return append(q.nums[q.front:], q.nums[:rear]...)
	}
	return q.nums[q.front:rear]
}
