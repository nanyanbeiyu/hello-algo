package double_ended_queue

import "fmt"

/* 基于环形数组实现的双向队列 */
type arrayDeque struct {
	nums        []int // 用于存储双向队列元素的数组
	front       int   // 队首指针，指向队首元素
	queSize     int   // 双向队列长度
	queCapacity int   // 队列容量（即最大容纳元素数量）
}

/* 初始化队列 */
func newArrayDeque(queCapacity int) *arrayDeque {
	return &arrayDeque{
		nums:        make([]int, queCapacity),
		queCapacity: queCapacity,
		front:       0,
		queSize:     0,
	}
}

/*获取双向队列长度*/
func (s *arrayDeque) size() int {
	return s.queSize
}

/*判断双向队列是否为空*/
func (s *arrayDeque) isEmpty() bool {
	return s.queSize == 0
}

/*计算环形数组的索引*/
func (s *arrayDeque) index(index int) int {
	// 通过取余操作实现数组首尾相连
	// 当 i 越过数组尾部后，回到头部
	// 当 i 越过数组头部后，回到尾部
	return (index + s.queCapacity) % s.queCapacity
}

/*队首入队*/
func (s *arrayDeque) pushFirst(val int) {
	if s.queSize == s.queCapacity {
		fmt.Println("双向队列已满")
		return
	}
	// 队首指针减1
	// 通过取余操作实现 front 越过数组头部后回到尾部
	s.front = s.index(s.front - 1)
	// 将val添加至队首
	s.nums[s.front] = val
	s.queSize++
}

/*队尾入队*/
func (s *arrayDeque) pushLast(val int) {
	if s.queSize == s.queCapacity {
		fmt.Println("双向队列已满")
		return
	}
	// 队尾指针加1
	rear := s.index(s.front + s.queSize)
	// 将val添加至队尾
	s.nums[rear] = val
	s.queSize++
}

/*队首出队*/
func (s *arrayDeque) popFirst() int {
	if s.queSize == 0 {
		fmt.Println("双向队列为空")
		return -1
	}
	// 队首元素出队
	val := s.nums[s.front]
	// 队首指针加1
	// 通过取余操作实现 front 越过数组头部后回到尾部
	s.front = s.index(s.front + 1)
	s.queSize--
	return val
}

/*队尾出队*/
func (s *arrayDeque) popLast() int {
	if s.queSize == 0 {
		fmt.Println("双向队列为空")
		return -1
	}
	// 队尾元素出队
	num := s.peekLast()
	s.queSize--
	return num
}

/*访问队首元素*/
func (s *arrayDeque) peekFirst() int {
	if s.queSize == 0 {
		fmt.Println("双向队列为空")
		return -1
	}
	// 返回队首元素
	return s.nums[s.front]
}

/*访问队尾元素*/
func (s *arrayDeque) peekLast() int {
	if s.queSize == 0 {
		fmt.Println("双向队列为空")
		return -1
	}
	// 返回队尾元素
	return s.nums[s.index(s.front+s.queSize-1)]
}

/* 获取 Slice 用于打印 */
func (s *arrayDeque) toSlice() []int {
	// 仅转换有效长度范围内的列表元素
	res := make([]int, s.queSize)
	for i, j := 0, s.front; i < s.queSize; i++ {
		res[i] = s.nums[s.index(j)]
		j++
	}
	return res
}
