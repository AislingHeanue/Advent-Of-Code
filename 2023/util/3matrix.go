package util

import (
	"cmp"
	"fmt"
	"image"
	"sort"
)

// ordering convention z,y,x (depth, height, width when viewed from the front)
type ThreeMatrix[V cmp.Ordered] [][][]V

func (m *ThreeMatrix[V]) GetDepth() int {
	return len(*m)
}

func (m *ThreeMatrix[V]) GetHeight() int {
	if m.GetDepth() == 0 {
		return 0
	}
	return len((*m)[0])
}

func (m *ThreeMatrix[V]) GetWidth() int {
	if m.GetHeight() == 0 {
		return 0
	}
	return len((*m)[0][0])
}

func (m *ThreeMatrix[V]) Clear() {
	d, h, w := m.GetDepth(), m.GetHeight(), m.GetWidth()
	*m = make(ThreeMatrix[V], d)
	for i := range *m {
		(*m)[i] = make([][]V, h)
		for j := range (*m)[i] {
			(*m)[i][j] = make([]V, w)
		}
	}
}

func (m *ThreeMatrix[V]) Fill(v V) {
	m.SetByRule(func(z, y, x int) V {
		return v
	})
}

func (m *ThreeMatrix[V]) Set(z, y, x int, v V) {
	(*m)[z][y][x] = v
}

func (m *ThreeMatrix[V]) Get(z, y, x int) V {
	return (*m)[z][y][x]
}

func (m *ThreeMatrix[V]) SetByRule(f func(z int, y int, x int) V) {
	for x := 0; x < m.GetWidth(); x++ {
		for y := 0; y < m.GetHeight(); y++ {
			for z := 0; z < m.GetDepth(); z++ {
				m.Set(z, y, x, f(z, y, x))
			}
		}
	}
}

func (m *ThreeMatrix[V]) Print(delimiter string) {
	for z := 0; z < m.GetDepth(); z++ {
		for y := 0; y < m.GetHeight(); y++ {
			for x := 0; x < m.GetWidth(); x++ {
				fmt.Printf("%v%s", (*m)[z][y][x], delimiter)
			}
			fmt.Print("\n")
		}
		fmt.Print("\n")
	}
}

func (m *ThreeMatrix[V]) PrintEvenlySpaced(delimiter string) {
	maxLength := 1
	for x := 0; x < m.GetWidth(); x++ {
		for y := 0; y < m.GetHeight(); y++ {
			for z := 0; z < m.GetDepth(); z++ {
				maxLength = max(maxLength, len(fmt.Sprint(m.Get(z, y, x))))
			}
		}
	}
	var leftSpacing int
	for z := 0; z < m.GetDepth(); z++ {
		for y := 0; y < m.GetHeight(); y++ {
			for x := 0; x < m.GetWidth(); x++ {
				leftSpacing = maxLength - len(fmt.Sprint(m.Get(z, y, x)))
				fmt.Printf("%*s%v%s", leftSpacing, "", (*m)[z][y][x], delimiter)
			}
			fmt.Printf("\n")
		}
		fmt.Print("\n")
	}
}

// due to the limits of my 2 dimensional screen, only one slice can be drawn at a time
func (m *ThreeMatrix[V]) Draw(z int) {
	Image = m.ToImage(z)
}

func (m *ThreeMatrix[V]) ToImage(z int) image.Image {
	tMap := make(map[V]float64)
	vs := m.Unique()
	for i, v := range vs {
		tMap[v] = float64(i) / (float64(len(vs)))
	}
	img := image.NewRGBA(image.Rect(0, 0, m.GetWidth(), m.GetHeight()))
	for x := 0; x < m.GetWidth(); x++ {
		for y := 0; y < m.GetHeight(); y++ {
			img.Set(x, y, ColourFunction(tMap[m.Get(z, y, x)]))
		}
	}

	return img
}

func (m *ThreeMatrix[V]) Unique() []V {
	set := make(map[V]bool)
	for x := 0; x < m.GetWidth(); x++ {
		for y := 0; y < m.GetHeight(); y++ {
			for z := 0; z < m.GetDepth(); z++ {
				set[m.Get(z, y, x)] = true
			}
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

func ThreeMap[K, V cmp.Ordered](m ThreeMatrix[K], f func(z int, y int, x int, value K) V) ThreeMatrix[V] {
	n := NewThreeMatrix[V](m.GetDepth(), m.GetHeight(), m.GetWidth())
	n.SetByRule(func(z, y, x int) V {
		return f(z, y, x, m.Get(z, y, x))
	})

	return n
}

func NewThreeMatrix[V cmp.Ordered](depth, height, width int) ThreeMatrix[V] {
	m := make(ThreeMatrix[V], depth)
	for i := range m {
		m[i] = make([][]V, height)
		for j := range m[i] {
			m[i][j] = make([]V, width)
		}
	}

	return m
}
