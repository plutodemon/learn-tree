package designPattern

import "sort"

// `~` 加不加的区别：加了可以包括自定义类型(底层是基本类型)，不加只能包括基本类型

// 类型集合

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func Min[E Unsigned | Signed](x, y E) E {
	if x < y {
		return x
	}
	return y
}

func IndexFunc[S ~[]E, E ~int](s S, f func(E) bool) int {
	for i := range s {
		if f(s[i]) {
			return i
		}
	}
	return -1
}

// 通用化的排序实现
type sortable[E any] struct {
	data []E
	cmp  CMP[E]
}

type CMP[E any] func(a, b E) bool

func (s sortable[E]) Len() int {
	return len(s.data)
}
func (s sortable[E]) Swap(i, j int) {
	s.data[i], s.data[j] = s.data[j], s.data[i]
}
func (s sortable[E]) Less(i, j int) bool {
	return s.cmp(s.data[i], s.data[j])
}

// Sort cmp(a, b) if a >= b is reverse
func Sort[E any](data []E, cmp CMP[E]) {
	sort.Sort(sortable[E]{data: data, cmp: cmp})
}
