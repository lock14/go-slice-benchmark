package benchmark

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func inline_slice[T Small | Medium | Large | ExtraLarge](sliceSize int) func(*testing.B) {
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a := make([]T, 0, sliceSize)
			for i := 0; i < sliceSize; i++ {
				var t T
				a = append(a, t)
			}
		}
	}
}
func pointer_slice[T Small | Medium | Large | ExtraLarge](sliceSize int) func(b *testing.B) {
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a := make([]*T, 0, sliceSize)
			for i := 0; i < sliceSize; i++ {
				a = append(a, new(T))
			}
		}
	}
}

type BenchmarkCase struct {
	name          string
	benchMarkFunc func(*testing.B)
}

func sliceBenchmarkCases[T Small | Medium | Large | ExtraLarge]() []BenchmarkCase {
	var t T
	typename := reflect.TypeOf(t).Name()
	cases := make([]BenchmarkCase, 0, 12)
	for i := 10; i <= 1000000; i *= 10 {
		cases = append(cases, BenchmarkCase{
			name:          fmt.Sprintf("%s_inline_slice_%d_elements", typename, i),
			benchMarkFunc: inline_slice[T](i),
		})
		cases = append(cases, BenchmarkCase{
			name:          fmt.Sprintf("%s_pointer_slice_%d_elements", typename, i),
			benchMarkFunc: pointer_slice[T](i),
		})
	}
	return cases
}

func BenchmarkSlice(b *testing.B) {
	b.Logf("size of Small: %d", unsafe.Sizeof(Small{}))
	b.Logf("size of Medium: %d", unsafe.Sizeof(Medium{}))
	b.Logf("size of Large: %d", unsafe.Sizeof(Large{}))
	b.Logf("size of ExtraLarge: %d", unsafe.Sizeof(ExtraLarge{}))

	cases := make([]BenchmarkCase, 0, 48)
	cases = append(cases, sliceBenchmarkCases[Small]()...)
	cases = append(cases, sliceBenchmarkCases[Medium]()...)
	cases = append(cases, sliceBenchmarkCases[Large]()...)
	cases = append(cases, sliceBenchmarkCases[ExtraLarge]()...)

	for _, bc := range cases {
		b.Run(bc.name, bc.benchMarkFunc)
	}
}

type Small struct {
	field int
}
type Medium struct {
	field [10]int
}

type Large struct {
	field [100]int
}

type ExtraLarge struct {
	field1 [1000]int
}
