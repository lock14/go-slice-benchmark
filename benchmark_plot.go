package main

import (
	"fmt"
	"reflect"
	"testing"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func inline_slice[T Small | Medium | Large | ExtraLarge](sliceSize int) func(*testing.B) {
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a := make([]T, 0, sliceSize)
			for j := 0; j < sliceSize; j++ {
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
			for j := 0; j < sliceSize; j++ {
				a = append(a, new(T))
			}
		}
	}
}

func inlineSliceBenchmarkCases[T Small | Medium | Large | ExtraLarge]() []BenchmarkCase {
	var t T
	typename := reflect.TypeOf(t).Name()
	cases := make([]BenchmarkCase, 0, 12)
	for i := 10; i <= 1000000; i *= 10 {
		cases = append(cases, BenchmarkCase{
			name:          fmt.Sprintf("%s_inline_slice_%d_elements", typename, i),
			xAxis:         i,
			benchMarkFunc: inline_slice[T](i),
		})
	}
	return cases
}

func pointerSliceBenchmarkCases[T Small | Medium | Large | ExtraLarge]() []BenchmarkCase {
	var t T
	typename := reflect.TypeOf(t).Name()
	cases := make([]BenchmarkCase, 0, 12)
	for i := 10; i <= 1000000; i *= 10 {
		cases = append(cases, BenchmarkCase{
			name:          fmt.Sprintf("%s_pointer_slice_%d_elements", typename, i),
			xAxis:         i,
			benchMarkFunc: pointer_slice[T](i),
		})
	}
	return cases
}

type BenchmarkCase struct {
	name          string
	xAxis         int
	benchMarkFunc func(*testing.B)
}

type PlotBenchMark struct {
	plotName string
	plotXYs  plotter.XYs
	cases    []BenchmarkCase
}

func main() {
	p := plot.New()
	p.Title.Text = "Slice of Pointer vs Concrete Type"
	p.X.Label.Text = "Slice Size"
	p.X.Scale = plot.LogScale{}
	p.X.Tick.Marker = plot.LogTicks{}
	p.Y.Label.Text = "Ns/Op"
	p.Y.Scale = plot.LogScale{}
	p.Y.Tick.Marker = plot.LogTicks{}

	plotBenchMarks := []PlotBenchMark{
		{
			plotName: "small inline",
			plotXYs:  plotter.XYs{},
			cases:    inlineSliceBenchmarkCases[Small](),
		},
		{
			plotName: "small pointer",
			plotXYs:  plotter.XYs{},
			cases:    pointerSliceBenchmarkCases[Small](),
		},
		{
			plotName: "medium inline",
			plotXYs:  plotter.XYs{},
			cases:    inlineSliceBenchmarkCases[Medium](),
		},
		{
			plotName: "medium pointer",
			plotXYs:  plotter.XYs{},
			cases:    pointerSliceBenchmarkCases[Medium](),
		},
		{
			plotName: "large inline",
			plotXYs:  plotter.XYs{},
			cases:    inlineSliceBenchmarkCases[Large](),
		},
		{
			plotName: "large pointer",
			plotXYs:  plotter.XYs{},
			cases:    pointerSliceBenchmarkCases[Large](),
		},
		{
			plotName: "x-large inline",
			plotXYs:  plotter.XYs{},
			cases:    inlineSliceBenchmarkCases[ExtraLarge](),
		},
		{
			plotName: "x-large pointer",
			plotXYs:  plotter.XYs{},
			cases:    pointerSliceBenchmarkCases[ExtraLarge](),
		},
	}
	a := make([]interface{}, 0)
	for _, plotBenchMark := range plotBenchMarks {
		for _, benchMarkCase := range plotBenchMark.cases {
			result := testing.Benchmark(benchMarkCase.benchMarkFunc)
			plotBenchMark.plotXYs = append(plotBenchMark.plotXYs, plotter.XY{
				X: float64(benchMarkCase.xAxis),
				Y: float64(result.NsPerOp()),
			})
		}
		a = append(a, plotBenchMark.plotName)
		a = append(a, plotBenchMark.plotXYs)
	}
	err := plotutil.AddLinePoints(p, a...)
	if err != nil {
		panic(err)
	}
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
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
