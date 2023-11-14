```
$ go test -benchmem -bench ./...
goos: linux 
goarch: amd64
pkg: main/benchmark
cpu: 11th Gen Intel(R) Core(TM) i7-1185G7 @ 3.00GHz
BenchmarkSlice/small_inline_slice_10_elements-8                   47252882             25.30 ns/op              80 B/op          1 allocs/op
BenchmarkSlice/small_pointer_slice_10_elements-8                   8427344             135.3 ns/op             160 B/op         11 allocs/op
BenchmarkSlice/medium_inline_slice_10_elements-8                  10051024             119.1 ns/op             896 B/op          1 allocs/op
BenchmarkSlice/medium_pointer_slice_10_elements-8                  4317172             325.6 ns/op             880 B/op         11 allocs/op
BenchmarkSlice/large_inline_slice_10_elements-8                     764784              1816 ns/op            8192 B/op          1 allocs/op
BenchmarkSlice/large_pointer_slice_10_elements-8                    457407              2423 ns/op            9040 B/op         11 allocs/op
BenchmarkSlice/extra_large_inline_slice_10_elements-8                77670             16228 ns/op           81920 B/op          1 allocs/op
BenchmarkSlice/extra_large_pointer_slice_10_elements-8               59167             19744 ns/op           82000 B/op         11 allocs/op
BenchmarkSlice/small_inline_slice_100_elements-8                   3914026             302.6 ns/op             896 B/op          1 allocs/op
BenchmarkSlice/small_pointer_slice_100_elements-8                   444705              2260 ns/op            1696 B/op        101 allocs/op
BenchmarkSlice/medium_inline_slice_100_elements-8                   564949              2069 ns/op            8192 B/op          1 allocs/op
BenchmarkSlice/medium_pointer_slice_100_elements-8                  257494              4488 ns/op            8896 B/op        101 allocs/op
BenchmarkSlice/large_inline_slice_100_elements-8                     95446             13682 ns/op           81920 B/op          1 allocs/op
BenchmarkSlice/large_pointer_slice_100_elements-8                    52207             21655 ns/op           90496 B/op        101 allocs/op
BenchmarkSlice/extra_large_inline_slice_100_elements-8                8936            140822 ns/op          802819 B/op          1 allocs/op
BenchmarkSlice/extra_large_pointer_slice_100_elements-8               5331            193365 ns/op          820100 B/op        101 allocs/op
BenchmarkSlice/small_inline_slice_1000_elements-8                   430786              2397 ns/op            8192 B/op          1 allocs/op
BenchmarkSlice/small_pointer_slice_1000_elements-8                   54160             20524 ns/op           16192 B/op       1001 allocs/op
BenchmarkSlice/medium_inline_slice_1000_elements-8                   71592             18534 ns/op           81920 B/op          1 allocs/op
BenchmarkSlice/medium_pointer_slice_1000_elements-8                  25383             42482 ns/op           88192 B/op       1001 allocs/op
BenchmarkSlice/large_inline_slice_1000_elements-8                     9700            124089 ns/op          802819 B/op          1 allocs/op
BenchmarkSlice/large_pointer_slice_1000_elements-8                    4358            249309 ns/op          904194 B/op       1001 allocs/op
BenchmarkSlice/extra_large_inline_slice_1000_elements-8               1081           1126305 ns/op         8003587 B/op          1 allocs/op
BenchmarkSlice/extra_large_pointer_slice_1000_elements-8               387           3293187 ns/op         8200201 B/op       1001 allocs/op
BenchmarkSlice/small_inline_slice_10000_elements-8                   51033             22304 ns/op           81920 B/op          1 allocs/op
BenchmarkSlice/small_pointer_slice_10000_elements-8                   4742            217997 ns/op          161920 B/op      10001 allocs/op
BenchmarkSlice/medium_inline_slice_10000_elements-8                   6373            283674 ns/op          802820 B/op          1 allocs/op
BenchmarkSlice/medium_pointer_slice_10000_elements-8                  2152            506319 ns/op          881921 B/op      10001 allocs/op
BenchmarkSlice/large_inline_slice_10000_elements-8                     710           1555243 ns/op         8003601 B/op          1 allocs/op
BenchmarkSlice/large_pointer_slice_10000_elements-8                    298           4110125 ns/op         9041927 B/op      10001 allocs/op
BenchmarkSlice/extra_large_inline_slice_10000_elements-8               116           9806305 ns/op        80003086 B/op          1 allocs/op
BenchmarkSlice/extra_large_pointer_slice_10000_elements-8               56          24319194 ns/op        82001920 B/op      10001 allocs/op
BenchmarkSlice/small_inline_slice_100000_elements-8                   6555            185899 ns/op          802823 B/op          1 allocs/op
BenchmarkSlice/small_pointer_slice_100000_elements-8                   416           2593094 ns/op         1602825 B/op     100001 allocs/op
BenchmarkSlice/medium_inline_slice_100000_elements-8                   625           2070135 ns/op         8003599 B/op          1 allocs/op
BenchmarkSlice/medium_pointer_slice_100000_elements-8                  176           6411874 ns/op         8802817 B/op     100001 allocs/op
BenchmarkSlice/large_inline_slice_100000_elements-8                     93          13473197 ns/op        80003088 B/op          1 allocs/op
BenchmarkSlice/large_pointer_slice_100000_elements-8                    38          27358065 ns/op        90402828 B/op     100001 allocs/op
BenchmarkSlice/extra_large_inline_slice_100000_elements-8               13          85183124 ns/op       800006151 B/op          1 allocs/op
BenchmarkSlice/extra_large_pointer_slice_100000_elements-8               6         194797194 ns/op       820002816 B/op     100001 allocs/op
BenchmarkSlice/small_inline_slice_1000000_elements-8                   859           1471497 ns/op         8003588 B/op          1 allocs/op
BenchmarkSlice/small_pointer_slice_1000000_elements-8                   56          25958045 ns/op        16003606 B/op    1000001 allocs/op
BenchmarkSlice/medium_inline_slice_1000000_elements-8                   58          18492885 ns/op        80003078 B/op          1 allocs/op
BenchmarkSlice/medium_pointer_slice_1000000_elements-8                  19          56913227 ns/op        88003584 B/op    1000001 allocs/op
BenchmarkSlice/large_inline_slice_1000000_elements-8                     9         127597741 ns/op       800006154 B/op          1 allocs/op
BenchmarkSlice/large_pointer_slice_1000000_elements-8                    4         272081173 ns/op       904003608 B/op    1000001 allocs/op
BenchmarkSlice/extra_large_inline_slice_1000000_elements-8               1        2646295183 ns/op      8000004096 B/op          1 allocs/op
BenchmarkSlice/extra_large_pointer_slice_1000000_elements-8              1        1621876718 ns/op      8200003680 B/op    1000002 allocs/op
PASS
ok      main/benchmark  74.534s
```
