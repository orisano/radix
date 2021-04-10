# radix
## Benchmarks
```
goos: darwin
goarch: amd64
pkg: github.com/orisano/radix
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkSortInt32/std/1000-8         	   22212	     51059 ns/op
BenchmarkSortInt32/radix/1000-8       	  149677	      8023 ns/op
BenchmarkSortInt32/std/10000-8        	    1070	   1105393 ns/op
BenchmarkSortInt32/radix/10000-8      	   17588	     68600 ns/op
BenchmarkSortInt32/std/100000-8       	      82	  14088866 ns/op
BenchmarkSortInt32/radix/100000-8     	    1663	    681334 ns/op
BenchmarkSortInt32/std/1000000-8      	       6	 171085842 ns/op
BenchmarkSortInt32/radix/1000000-8    	     156	   7604423 ns/op
BenchmarkSortInt64/std/1000-8         	   23025	     51764 ns/op
BenchmarkSortInt64/radix/1000-8       	   75854	     16087 ns/op
BenchmarkSortInt64/std/10000-8        	    1078	   1105161 ns/op
BenchmarkSortInt64/radix/10000-8      	    7328	    152684 ns/op
BenchmarkSortInt64/std/100000-8       	      86	  13945998 ns/op
BenchmarkSortInt64/radix/100000-8     	     710	   1683827 ns/op
BenchmarkSortInt64/std/1000000-8      	       6	 171971064 ns/op
BenchmarkSortInt64/radix/1000000-8    	      58	  19354205 ns/op
PASS
ok  	github.com/orisano/radix	39.918s
```

## References
https://lemire.me/blog/2021/04/09/how-fast-can-you-sort-arrays-of-integers-in-java/

## Author
Nao Yonashiro

## License
Apache 2.0
