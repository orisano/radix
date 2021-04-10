# radix
[![Test](https://github.com/orisano/radix/actions/workflows/go.yml/badge.svg)](https://github.com/orisano/radix/actions/workflows/go.yml)
## Benchmarks
```
goos: darwin
goarch: amd64
pkg: github.com/orisano/radix
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkSortInt32/std/1000-8         	   19708	     53759 ns/op
BenchmarkSortInt32/radix/1000-8       	  144142	      8031 ns/op
BenchmarkSortInt32/std/10000-8        	    1088	   1094595 ns/op
BenchmarkSortInt32/radix/10000-8      	   17654	     68060 ns/op
BenchmarkSortInt32/std/100000-8       	      85	  14037974 ns/op
BenchmarkSortInt32/radix/100000-8     	    1687	    684250 ns/op
BenchmarkSortInt32/std/1000000-8      	       6	 168432310 ns/op
BenchmarkSortInt32/radix/1000000-8    	     159	   7402114 ns/op
BenchmarkSortInt64/std/1000-8         	   24103	     50039 ns/op
BenchmarkSortInt64/radix/1000-8       	   76006	     15448 ns/op
BenchmarkSortInt64/std/10000-8        	    1063	   1104567 ns/op
BenchmarkSortInt64/radix/10000-8      	    7221	    148631 ns/op
BenchmarkSortInt64/std/100000-8       	      86	  13775945 ns/op
BenchmarkSortInt64/radix/100000-8     	     716	   1651069 ns/op
BenchmarkSortInt64/std/1000000-8      	       6	 168548893 ns/op
BenchmarkSortInt64/radix/1000000-8    	      56	  18255245 ns/op
BenchmarkSortFloat64/std/1000-8       	   14665	     81712 ns/op
BenchmarkSortFloat64/radix/1000-8     	   57018	     21110 ns/op
BenchmarkSortFloat64/std/10000-8      	     924	   1284250 ns/op
BenchmarkSortFloat64/radix/10000-8    	    5956	    190296 ns/op
BenchmarkSortFloat64/std/100000-8     	      74	  16395907 ns/op
BenchmarkSortFloat64/radix/100000-8   	     559	   2142369 ns/op
BenchmarkSortFloat64/std/1000000-8    	       6	 195785586 ns/op
BenchmarkSortFloat64/radix/1000000-8  	      48	  24025723 ns/op
PASS
ok  	github.com/orisano/radix	54.569s
```

## References
https://lemire.me/blog/2021/04/09/how-fast-can-you-sort-arrays-of-integers-in-java/

## Author
Nao Yonashiro

## License
Apache 2.0
