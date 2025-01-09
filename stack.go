package gscXml

type Stack[T any] struct {
	items []T
}

// Push 操作：将元素推入栈中
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop 操作：移除并返回栈顶元素
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zeroValue T
		return zeroValue, false // 栈为空时返回零值
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, true
}

// Peek 操作：查看栈顶元素但不移除它
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zeroValue T
		return zeroValue, false // 栈为空时返回零值
	}
	return s.items[len(s.items)-1], true
}

// IsEmpty 操作：检查栈是否为空
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}
