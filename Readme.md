```
$ go test -benchmem -bench ./...
goos: linux
goarch: amd64
pkg: github.com/lock14/go-slice-benchmark
cpu: Intel(R) Xeon(R) W-2135 CPU @ 3.70GHz
BenchmarkSlice/Small_inline_slice_10_elements-12         	17634830	        68.42 ns/op	      80 B/op	       1 allocs/op
BenchmarkSlice/Small_pointer_slice_10_elements-12        	 3117170	       366.1 ns/op	     160 B/op	      11 allocs/op
BenchmarkSlice/Small_inline_slice_100_elements-12        	 3189034	       368.2 ns/op	     896 B/op	       1 allocs/op
BenchmarkSlice/Small_pointer_slice_100_elements-12       	  382812	      3214 ns/op	    1696 B/op	     101 allocs/op
BenchmarkSlice/Small_inline_slice_1000_elements-12       	  276752	      3790 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSlice/Small_pointer_slice_1000_elements-12      	   36691	     31753 ns/op	   16192 B/op	    1001 allocs/op
BenchmarkSlice/Small_inline_slice_10000_elements-12      	   44324	     28266 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSlice/Small_pointer_slice_10000_elements-12     	    3559	    284725 ns/op	  161920 B/op	   10001 allocs/op
BenchmarkSlice/Small_inline_slice_100000_elements-12     	    4950	    262121 ns/op	  802825 B/op	       1 allocs/op
BenchmarkSlice/Small_pointer_slice_100000_elements-12    	     345	   3292657 ns/op	 1602826 B/op	  100001 allocs/op
BenchmarkSlice/Small_inline_slice_1000000_elements-12    	     524	   2225758 ns/op	 8003587 B/op	       1 allocs/op
BenchmarkSlice/Small_pointer_slice_1000000_elements-12   	      45	  24279759 ns/op	16003599 B/op	 1000001 allocs/op
BenchmarkSlice/Medium_inline_slice_10_elements-12        	 3085176	       396.1 ns/op	     896 B/op	       1 allocs/op
BenchmarkSlice/Medium_pointer_slice_10_elements-12       	 1474717	       782.8 ns/op	     880 B/op	      11 allocs/op
BenchmarkSlice/Medium_inline_slice_100_elements-12       	  285411	      3920 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSlice/Medium_pointer_slice_100_elements-12      	  163668	      6779 ns/op	    8896 B/op	     101 allocs/op
BenchmarkSlice/Medium_inline_slice_1000_elements-12      	   33260	     36157 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSlice/Medium_pointer_slice_1000_elements-12     	   18082	     67389 ns/op	   88192 B/op	    1001 allocs/op
BenchmarkSlice/Medium_inline_slice_10000_elements-12     	    3440	    317145 ns/op	  802824 B/op	       1 allocs/op
BenchmarkSlice/Medium_pointer_slice_10000_elements-12    	    1640	    694422 ns/op	  881922 B/op	   10001 allocs/op
BenchmarkSlice/Medium_inline_slice_100000_elements-12    	     414	   2611110 ns/op	 8003601 B/op	       1 allocs/op
BenchmarkSlice/Medium_pointer_slice_100000_elements-12   	     128	   9423877 ns/op	 8802827 B/op	  100001 allocs/op
BenchmarkSlice/Medium_inline_slice_1000000_elements-12   	      73	  16265265 ns/op	80003078 B/op	       1 allocs/op
BenchmarkSlice/Medium_pointer_slice_1000000_elements-12  	      24	  51886459 ns/op	88003673 B/op	 1000001 allocs/op
BenchmarkSlice/Large_inline_slice_10_elements-12         	  386858	      3372 ns/op	    8192 B/op	       1 allocs/op
BenchmarkSlice/Large_pointer_slice_10_elements-12        	  335685	      3175 ns/op	    9040 B/op	      11 allocs/op
BenchmarkSlice/Large_inline_slice_100_elements-12        	   45427	     25257 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSlice/Large_pointer_slice_100_elements-12       	   40512	     30624 ns/op	   90496 B/op	     101 allocs/op
BenchmarkSlice/Large_inline_slice_1000_elements-12       	    5011	    231234 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSlice/Large_pointer_slice_1000_elements-12      	    3183	    348893 ns/op	  904194 B/op	    1001 allocs/op
BenchmarkSlice/Large_inline_slice_10000_elements-12      	     590	   2153888 ns/op	 8003602 B/op	       1 allocs/op
BenchmarkSlice/Large_pointer_slice_10000_elements-12     	     222	   5340563 ns/op	 9041937 B/op	   10001 allocs/op
BenchmarkSlice/Large_inline_slice_100000_elements-12     	      90	  13447315 ns/op	80003090 B/op	       1 allocs/op
BenchmarkSlice/Large_pointer_slice_100000_elements-12    	      39	  31407895 ns/op	90402833 B/op	  100001 allocs/op
BenchmarkSlice/Large_inline_slice_1000000_elements-12    	      10	 101270908 ns/op	800006144 B/op	       1 allocs/op
BenchmarkSlice/Large_pointer_slice_1000000_elements-12   	       6	 207495868 ns/op	904003616 B/op	 1000001 allocs/op
BenchmarkSlice/ExtraLarge_inline_slice_10_elements-12    	   52473	     24751 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSlice/ExtraLarge_pointer_slice_10_elements-12   	   46599	     29918 ns/op	   82000 B/op	      11 allocs/op
BenchmarkSlice/ExtraLarge_inline_slice_100_elements-12   	    6196	    165310 ns/op	  802819 B/op	       1 allocs/op
BenchmarkSlice/ExtraLarge_pointer_slice_100_elements-12  	    4149	    245199 ns/op	  820100 B/op	     101 allocs/op
BenchmarkSlice/ExtraLarge_inline_slice_1000_elements-12  	     933	   1407656 ns/op	 8003589 B/op	       1 allocs/op
BenchmarkSlice/ExtraLarge_pointer_slice_1000_elements-12 	     271	   4285251 ns/op	 8200210 B/op	    1001 allocs/op
BenchmarkSlice/ExtraLarge_inline_slice_10000_elements-12 	     100	  12002116 ns/op	80003080 B/op	       1 allocs/op
BenchmarkSlice/ExtraLarge_pointer_slice_10000_elements-12         	      50	  24341043 ns/op	82001946 B/op	   10001 allocs/op
BenchmarkSlice/ExtraLarge_inline_slice_100000_elements-12         	      12	  85325388 ns/op	800006144 B/op	       1 allocs/op
BenchmarkSlice/ExtraLarge_pointer_slice_100000_elements-12        	       6	 186596768 ns/op	820002864 B/op	  100001 allocs/op
BenchmarkSlice/ExtraLarge_inline_slice_1000000_elements-12        	       1	1643117000 ns/op	8000004096 B/op	       1 allocs/op
BenchmarkSlice/ExtraLarge_pointer_slice_1000000_elements-12       	       1	1585078405 ns/op	8200003688 B/op	 1000003 allocs/op
PASS
ok  	github.com/lock14/go-slice-benchmark	76.041s
```
