package main

import (
	"fmt"
)

type SayAble[T int | string] interface {
	Say() T
}
type Person[T int | string] struct {
	msg T
}

func (p Person[T]) Say() T {
	return p.msg
}

// 队列
type Queue[T any] []T

func (q *Queue[T]) Push(e T) {
	*q = append(*q, e)
}
func (q *Queue[T]) Pop(e T) (_ T) {
	if q.Size() > 0 {
		res := q.Peek()
		*q = (*q)[1:]
		return res
	}
	return
}
func (q *Queue[T]) Peek() (_ T) {
	if q.Size() > 0 {
		return (*q)[0]
	}
	return
}
func (q *Queue[T]) Size() int {
	return len(*q)
}

// 堆

type Comparator[T any] func(a, b T) int

type BinaryHeap[T any] struct {
	s []T
	c Comparator[T]
}

func (heap *BinaryHeap[T]) Peek() (_ T) {
	if heap.Size() > 0 {
		return heap.s[0]
	}
	return
}

//func (heap *BinaryHeap[T]) Pop() (_ T) {
//	size := heap.Size()
//	if size > 0 {
//		res := heap.s[0]
//		heap.s[0], heap.s[size-1] = heap.s[size-1], heap.s[0]
//		heap.s = heap.s[:size-1]
//		heap.down(0)
//		return res
//	}
//	return
//}
//func (heap *BinaryHeap[T]) Push(e T) {
//	heap.s = append(heap.s, e)
//	heap.up(heap.Size() - 1)
//}
//func (heap *BinaryHeap[T]) up(i int) {
//	if heap.Size() == 0 || i < 0 || i >= heap.Size() {
//		return
//	}
//	for parentIndex := i>>1 - 1; parentIndex >= 0; parentIndex = i>>1 - 1 {
//		if heap.compare(heap.s[i], heap.s[parentIndex]) >= 0 {
//			break
//		}
//		heap.s[i], heap.s[parentIndex] = heap.s[parentIndex], heap.s[i]
//		i = parentIndex
//	}
//}

//	func (heap *BinaryHeap[T]) down(i int) {
//		if heap.Size() == 0 || i < 0 || i >= heap.Size() {
//			return
//		}
//		size := heap.Size()
//		for lsonIndex := i<<1 + 1; lsonIndex < size; lsonIndex = i<<1 + 1 {
//			rsonIndex := lsonIndex + 1
//			if rsonIndex < size && heap.compare(heap.s[rsonIndex], heap.s[lsonIndex]) < 0 {
//				lsonIndex = rsonIndex
//			}
//			if heap.compare(heap.s[i], heap.s[lsonIndex]) <= 0 {
//				break
//			}
//			heap.s[i], heap.s[lsonIndex] = heap.s[lsonIndex], heap.s[i]
//			i = lsonIndex
//		}
//	}
func (heap *BinaryHeap[T]) Size() int {
	return len(heap.s)
}

type Person2 struct {
	Age  int
	Name string
}

func main() {
	var s SayAble[string]
	s = Person[string]{"hello"}
	fmt.Println(s.Say())

	//heap := NewHeap[Person2](10, func(a, b Person2) int {
	//	return cmp.Compare(a.Age, b.Age)
	//})
	//heap.Push(Person2{Age: 10, Name: "John"})
	//heap.Push(Person2{Age: 18, Name: "tom"})
	//heap.Push(Person2{Age: 9, Name: "lili"})
	//heap.Push(Person2{Age: 39, Name: "miki"})
	//fmt.Println(heap.Peek())
	//fmt.Println(heap.Pop())
	//fmt.Println(heap.Peek())

}
