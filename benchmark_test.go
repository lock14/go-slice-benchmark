package benchmark

import (
    "testing"
    "unsafe"
)

func BenchmarkSlice(b *testing.B) {
    b.Logf("size of Small: %d", unsafe.Sizeof(Small{}))
    b.Logf("size of Medium: %d", unsafe.Sizeof(Medium{}))
    b.Logf("size of Large: %d", unsafe.Sizeof(Large{}))
    b.Logf("size of ExtraLarge: %d", unsafe.Sizeof(ExtraLarge{}))
    cases := []struct {
        name      string
        sliceSize int
    }{
        {
            name:      "10_elements",
            sliceSize: 10,
        },
        {
            name:      "100_elements",
            sliceSize: 100,
        },
        {
            name:      "1000_elements",
            sliceSize: 1000,
        },
        {
            name:      "10000_elements",
            sliceSize: 10000,
        },
        {
            name:      "100000_elements",
            sliceSize: 100000,
        },
        {
            name:      "1000000_elements",
            sliceSize: 1000000,
        },
    }
    for _, bc := range cases {
        bc := bc
        b.Run("small_inline_slice_"+bc.name, func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                a := make([]Small, 0, bc.sliceSize)
                for i := 0; i < bc.sliceSize; i++ {
                    a = append(a, Small{})
                }
            }
        })
        b.Run("small_pointer_slice_"+bc.name, func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                a := make([]*Small, 0, bc.sliceSize)
                for i := 0; i < bc.sliceSize; i++ {
                    a = append(a, &Small{})
                }
            }
        })
        b.Run("medium_inline_slice_"+bc.name, func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                a := make([]Medium, 0, bc.sliceSize)
                for i := 0; i < bc.sliceSize; i++ {
                    a = append(a, Medium{})
                }
            }
        })
        b.Run("medium_pointer_slice_"+bc.name, func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                a := make([]*Medium, 0, bc.sliceSize)
                for i := 0; i < bc.sliceSize; i++ {
                    a = append(a, &Medium{})
                }
            }
        })
        b.Run("large_inline_slice_"+bc.name, func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                a := make([]Large, 0, bc.sliceSize)
                for i := 0; i < bc.sliceSize; i++ {
                    a = append(a, Large{})
                }
            }
        })
        b.Run("large_pointer_slice_"+bc.name, func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                a := make([]*Large, 0, bc.sliceSize)
                for i := 0; i < bc.sliceSize; i++ {
                    a = append(a, &Large{})
                }
            }
        })
        b.Run("extra_large_inline_slice_"+bc.name, func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                a := make([]ExtraLarge, 0, bc.sliceSize)
                for i := 0; i < bc.sliceSize; i++ {
                    a = append(a, ExtraLarge{})
                }
            }
        })
        b.Run("extra_large_pointer_slice_"+bc.name, func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                a := make([]*ExtraLarge, 0, bc.sliceSize)
                for i := 0; i < bc.sliceSize; i++ {
                    a = append(a, &ExtraLarge{})
                }
            }
        })
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
