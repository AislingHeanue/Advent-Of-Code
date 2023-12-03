package util

import (
	"cmp"
	"fmt"
)

type UnorderedMatrix[V any] [][]V

func (m *UnorderedMatrix[V]) GetHeight() int {
	return len(*m)
}

func (m *UnorderedMatrix[V]) GetWidth() int {
	if m.GetHeight() == 0 {
		return 0
	}
	return len((*m)[0])
}

func (m *UnorderedMatrix[V]) Clear() {
	h, w := m.GetHeight(), m.GetWidth()
	*m = make(UnorderedMatrix[V], h)
	for i := range *m {
		(*m)[i] = make([]V, w)
	}
}

func (m *UnorderedMatrix[V]) Fill(v V) {
	m.SetByRule(func(y, x int) V {
		return v
	})
}

func (m *UnorderedMatrix[V]) MustSet(y, x int, v V) {
	(*m)[y][x] = v
}

func (m *UnorderedMatrix[V]) Set(y, x int, v V) bool {
	if y < 0 || x < 0 || y >= m.GetHeight() || x >= m.GetWidth() {
		return false
	}
	(*m)[y][x] = v

	return true
}

func (m *UnorderedMatrix[V]) MustGet(y, x int) V {
	return (*m)[y][x]
}

func (m *UnorderedMatrix[V]) Get(y, x int) (V, bool) {
	if y < 0 || x < 0 || y >= m.GetHeight() || x >= m.GetWidth() {
		var val V
		return val, false
	}
	return (*m)[y][x], true
}

func (m *UnorderedMatrix[V]) SetByRule(f func(y int, x int) V) {
	for x := 0; x < m.GetWidth(); x++ {
		for y := 0; y < m.GetHeight(); y++ {
			m.Set(y, x, f(y, x))
		}
	}
}

func (m *UnorderedMatrix[V]) Transpose() UnorderedMatrix[V] {
	n := UnorderedMap[V](*m, func(y, x int, v V) V {
		return (*m)[x][y]
	})
	return n
}

func (m *UnorderedMatrix[V]) Print(delimiter string) {
	for y := 0; y < m.GetHeight(); y++ {
		for x := 0; x < m.GetWidth(); x++ {
			fmt.Printf("%v%s", (*m)[y][x], delimiter)
		}
		fmt.Printf("\n")
	}
}

func (m *UnorderedMatrix[V]) PrintEvenlySpaced(delimiter string) {
	maxLength := 1
	for x := 0; x < m.GetWidth(); x++ {
		for y := 0; y < m.GetHeight(); y++ {
			maxLength = max(maxLength, len(fmt.Sprint(m.Get(y, x))))
		}
	}
	var leftSpacing int
	for y := 0; y < m.GetHeight(); y++ {
		for x := 0; x < m.GetWidth(); x++ {
			leftSpacing = maxLength - len(fmt.Sprint(m.Get(y, x)))
			fmt.Printf("%*s%v%s", leftSpacing, "", (*m)[y][x], delimiter)
		}
		fmt.Printf("\n")
	}
}

func UnorderedMap[K, V any](m UnorderedMatrix[K], f func(y int, x int, value K) V) UnorderedMatrix[V] {
	n := NewUnorderedMatrix[V](m.GetHeight(), m.GetWidth())
	n.SetByRule(func(y, x int) V {
		v, _ := m.Get(y, x)
		return f(y, x, v)
	})

	return n
}

func UnorderedMapToOrdered[K any, V cmp.Ordered](m UnorderedMatrix[K], f func(y int, x int, value K) V) Matrix[V] {
	n := NewMatrix[V](m.GetHeight(), m.GetWidth())
	n.SetByRule(func(y, x int) V {
		v, _ := m.Get(y, x)
		return f(y, x, v)
	})

	return n
}

func MapToUnordered[K cmp.Ordered, V any](m Matrix[K], f func(y int, x int, value K) V) UnorderedMatrix[V] {
	n := NewUnorderedMatrix[V](m.GetHeight(), m.GetWidth())
	n.SetByRule(func(y, x int) V {
		v, _ := m.Get(y, x)
		return f(y, x, v)
	})

	return n
}

func NewUnorderedMatrix[V any](height int, width int) UnorderedMatrix[V] {
	m := make(UnorderedMatrix[V], height)
	for i := range m {
		m[i] = make([]V, width)
	}

	return m
}
