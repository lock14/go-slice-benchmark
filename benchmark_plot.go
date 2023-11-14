package main

import (
	"testing"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func runInline[T any](size int) func(*testing.B) {
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a := make([]T, 0, size)
			for i := 0; i < size; i++ {
				var t T
				a = append(a, t)
			}
		}
	}
}

func runPointer[T any](size int) func(b *testing.B) {
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a := make([]*T, 0, size)
			for i := 0; i < size; i++ {
				var t T
				a = append(a, &t)
			}
		}
	}
}

func main() {
	p := plot.New()
	p.Title.Text = "Array of Pointer vs Concrete Type"
	p.X.Label.Text = "Size Of Array"
	p.X.Scale = plot.LogScale{}
	p.X.Tick.Marker = plot.LogTicks{}
	p.Y.Label.Text = "Ns/Op"
	p.Y.Scale = plot.LogScale{}
	p.Y.Tick.Marker = plot.LogTicks{}

	medConcrete := plotter.XYs{}
	medPointer := plotter.XYs{}
	lrgConcrete := plotter.XYs{}
	lrgPointer := plotter.XYs{}
	xlrgConcrete := plotter.XYs{}
	xlrgPointer := plotter.XYs{}

	for i := 1; i <= 1_000_000; i = i * 10 {
		result := testing.Benchmark(runInline[Medium](i))
		medConcrete = append(medConcrete, plotter.XY{float64(i), float64(result.NsPerOp())})
		result = testing.Benchmark(runPointer[Medium](i))
		medPointer = append(medPointer, plotter.XY{float64(i), float64(result.NsPerOp())})
		result = testing.Benchmark(runInline[Large](i))
		lrgConcrete = append(lrgConcrete, plotter.XY{float64(i), float64(result.NsPerOp())})
		result = testing.Benchmark(runPointer[Large](i))
		lrgPointer = append(lrgPointer, plotter.XY{float64(i), float64(result.NsPerOp())})
		result = testing.Benchmark(runInline[ExtraLarge](i))
		xlrgConcrete = append(xlrgConcrete, plotter.XY{float64(i), float64(result.NsPerOp())})
		result = testing.Benchmark(runPointer[ExtraLarge](i))
		xlrgPointer = append(xlrgPointer, plotter.XY{float64(i), float64(result.NsPerOp())})
	}
	err := plotutil.AddLinePoints(p,
		"medium inline", medConcrete,
		"medium pointer", medPointer,
		"large inline", lrgConcrete,
		"large pointer", lrgPointer,
		"x-large inline", xlrgConcrete,
		"x-large pointer", xlrgPointer)
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
