# radix
[![Test](https://github.com/orisano/radix/actions/workflows/go.yml/badge.svg)](https://github.com/orisano/radix/actions/workflows/go.yml)
## Benchmarks
```
goos: darwin
goarch: amd64
pkg: github.com/orisano/radix
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkSortInt32/parallel-radix/1000-8         	   9175	   126633 ns/op
BenchmarkSortInt32/radix/1000-8                  	 146168	     7416 ns/op
BenchmarkSortInt32/std/1000-8                    	  24579	    49791 ns/op
BenchmarkSortInt32/parallel-radix/10000-8        	   5852	   196073 ns/op
BenchmarkSortInt32/radix/10000-8                 	  19602	    61577 ns/op
BenchmarkSortInt32/std/10000-8                   	   1104	  1092107 ns/op
BenchmarkSortInt32/parallel-radix/100000-8       	   2031	   564766 ns/op
BenchmarkSortInt32/radix/100000-8                	   1848	   642259 ns/op
BenchmarkSortInt32/std/100000-8                  	     85	 13660845 ns/op
BenchmarkSortInt32/parallel-radix/1000000-8      	    294	  3998838 ns/op
BenchmarkSortInt32/radix/1000000-8               	    177	  6596822 ns/op
BenchmarkSortInt32/std/1000000-8                 	      6	167633926 ns/op
BenchmarkSortInt32/parallel-radix/10000000-8     	     33	 33278574 ns/op
BenchmarkSortInt32/radix/10000000-8              	     18	 65536943 ns/op
BenchmarkSortInt32/std/10000000-8                	      1	2016386593 ns/op
BenchmarkSortInt64/radix/1000-8                  	  81774	    14584 ns/op
BenchmarkSortInt64/std/1000-8                    	  24586	    49592 ns/op
BenchmarkSortInt64/radix/10000-8                 	   8539	   136973 ns/op
BenchmarkSortInt64/std/10000-8                   	   1090	  1103024 ns/op
BenchmarkSortInt64/radix/100000-8                	    798	  1493399 ns/op
BenchmarkSortInt64/std/100000-8                  	     84	 13716647 ns/op
BenchmarkSortInt64/radix/1000000-8               	     74	 16567415 ns/op
BenchmarkSortInt64/std/1000000-8                 	      6	168054870 ns/op
BenchmarkSortFloat64/radix/1000-8                	  59556	    20545 ns/op
BenchmarkSortFloat64/std/1000-8                  	  14757	    81297 ns/op
BenchmarkSortFloat64/radix/10000-8               	   6640	   173604 ns/op
BenchmarkSortFloat64/std/10000-8                 	    933	  1284367 ns/op
BenchmarkSortFloat64/radix/100000-8              	    556	  2119231 ns/op
BenchmarkSortFloat64/std/100000-8                	     74	 16072978 ns/op
BenchmarkSortFloat64/radix/1000000-8             	     54	 20893954 ns/op
BenchmarkSortFloat64/std/1000000-8               	      6	195648988 ns/op
PASS
ok  	github.com/orisano/radix	67.594s
```

## References
https://lemire.me/blog/2021/04/09/how-fast-can-you-sort-arrays-of-integers-in-java/

## Author
Nao Yonashiro

## License
Apache 2.0
