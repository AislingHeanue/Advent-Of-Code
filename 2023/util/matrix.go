package util

import (
	"cmp"
	"fmt"
	"image"
	"sort"
)

type Matrix[V cmp.Ordered] [][]V

func (m *Matrix[V]) GetHeight() int {
	return len(*m)
}

func (m *Matrix[V]) GetWidth() int {
	if m.GetHeight() == 0 {
		return 0
	}
	return len((*m)[0])
}

func (m *Matrix[V]) Clear() {
	h, w := m.GetHeight(), m.GetWidth()
	*m = make(Matrix[V], h)
	for i := range *m {
		(*m)[i] = make([]V, w)
	}
}

func (m *Matrix[V]) Fill(v V) {
	m.SetByRule(func(y, x int) V {
		return v
	})
}

func (m *Matrix[V]) MustSet(y, x int, v V) {
	(*m)[y][x] = v
}

func (m *Matrix[V]) Set(y, x int, v V) bool {
	if y < 0 || x < 0 || y >= m.GetHeight() || x >= m.GetWidth() {
		return false
	}
	(*m)[y][x] = v

	return true
}

func (m *Matrix[V]) MustGet(y, x int) V {
	return (*m)[y][x]
}

func (m *Matrix[V]) Get(y, x int) (V, bool) {
	if y < 0 || x < 0 || y >= m.GetHeight() || x >= m.GetWidth() {
		var val V
		return val, false
	}
	return (*m)[y][x], true
}

func (m *Matrix[V]) SetByRule(f func(y int, x int) V) {
	for x := 0; x < m.GetWidth(); x++ {
		for y := 0; y < m.GetHeight(); y++ {
			m.Set(y, x, f(y, x))
		}
	}
}

func (m *Matrix[V]) Transpose() Matrix[V] {
	n := Map[V](*m, func(y, x int, v V) V {
		return (*m)[x][y]
	})
	return n
}

func (m *Matrix[V]) Print(delimiter string) {
	for y := 0; y < m.GetHeight(); y++ {
		for x := 0; x < m.GetWidth(); x++ {
			fmt.Printf("%v%s", (*m)[y][x], delimiter)
		}
		fmt.Printf("\n")
	}
}

func (m *Matrix[V]) PrintEvenlySpaced(delimiter string) {
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

func (m *Matrix[V]) Draw() {
	Image = m.ToImage()
}

func (m *Matrix[V]) ToImage() image.Image {
	tMap := make(map[V]float64)
	vs := m.Unique()
	for i, v := range vs {
		tMap[v] = float64(i) / (float64(len(vs)))
	}
	img := image.NewRGBA(image.Rect(0, 0, m.GetWidth(), m.GetHeight()))
	for x := 0; x < m.GetWidth(); x++ {
		for y := 0; y < m.GetHeight(); y++ {
			v, _ := m.Get(y, x)
			img.Set(x, y, ColourFunction(tMap[v]))
		}
	}

	return img
}

func (m *Matrix[V]) Unique() []V {
	set := make(map[V]bool)
	for x := 0; x < m.GetWidth(); x++ {
		for y := 0; y < m.GetHeight(); y++ {
			v, _ := m.Get(y, x)
			set[v] = true
		}
	}
	keys := make([]V, 0, len(set))
	for k := range set {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	return keys
}

func Map[K, V cmp.Ordered](m Matrix[K], f func(y int, x int, value K) V) Matrix[V] {
	n := NewMatrix[V](m.GetHeight(), m.GetWidth())
	n.SetByRule(func(y, x int) V {
		v, _ := m.Get(y, x)
		return f(y, x, v)
	})

	return n
}

func NewMatrix[V cmp.Ordered](height int, width int) Matrix[V] {
	m := make(Matrix[V], height)
	for i := range m {
		m[i] = make([]V, width)
	}

	return m
}
