# radix
[![Test](https://github.com/orisano/radix/actions/workflows/go.yml/badge.svg)](https://github.com/orisano/radix/actions/workflows/go.yml)
## Benchmarks
```
goos: darwin
goarch: amd64
pkg: github.com/orisano/radix
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkSortInt32/parallel-radix/1000-8         	    9241	    133422 ns/op
BenchmarkSortInt32/radix/1000-8                  	  135092	      8080 ns/op
BenchmarkSortInt32/std/1000-8                    	   24223	     50666 ns/op
BenchmarkSortInt32/parallel-radix/10000-8        	    5595	    196793 ns/op
BenchmarkSortInt32/radix/10000-8                 	   17590	     68618 ns/op
BenchmarkSortInt32/std/10000-8                   	    1071	   1096447 ns/op
BenchmarkSortInt32/parallel-radix/100000-8       	    1982	    568277 ns/op
BenchmarkSortInt32/radix/100000-8                	    1705	    689857 ns/op
BenchmarkSortInt32/std/100000-8                  	      84	  13942514 ns/op
BenchmarkSortInt32/parallel-radix/1000000-8      	     296	   3896824 ns/op
BenchmarkSortInt32/radix/1000000-8               	     156	   7603199 ns/op
BenchmarkSortInt32/std/1000000-8                 	       6	 170011070 ns/op
BenchmarkSortInt32/parallel-radix/10000000-8     	      31	  32860509 ns/op
BenchmarkSortInt32/radix/10000000-8              	      15	  72790575 ns/op
BenchmarkSortInt32/std/10000000-8                	       1	1967581046 ns/op
BenchmarkSortInt64/radix/1000-8                  	   76771	     15643 ns/op
BenchmarkSortInt64/std/1000-8                    	   24112	     50633 ns/op
BenchmarkSortInt64/radix/10000-8                 	    7333	    151333 ns/op
BenchmarkSortInt64/std/10000-8                   	    1076	   1113783 ns/op
BenchmarkSortInt64/radix/100000-8                	     729	   1639533 ns/op
BenchmarkSortInt64/std/100000-8                  	      81	  13797465 ns/op
BenchmarkSortInt64/radix/1000000-8               	      61	  19137255 ns/op
BenchmarkSortInt64/std/1000000-8                 	       6	 169879131 ns/op
BenchmarkSortFloat64/radix/1000-8                	   56588	     21197 ns/op
BenchmarkSortFloat64/std/1000-8                  	   14817	     81180 ns/op
BenchmarkSortFloat64/radix/10000-8               	    5872	    192180 ns/op
BenchmarkSortFloat64/std/10000-8                 	     934	   1303996 ns/op
BenchmarkSortFloat64/radix/100000-8              	     524	   2199037 ns/op
BenchmarkSortFloat64/std/100000-8                	      70	  16114116 ns/op
BenchmarkSortFloat64/radix/1000000-8             	      49	  24124947 ns/op
BenchmarkSortFloat64/std/1000000-8               	       6	 196980290 ns/op
PASS
ok  	github.com/orisano/radix	66.290s
```

## References
https://lemire.me/blog/2021/04/09/how-fast-can-you-sort-arrays-of-integers-in-java/

## Author
Nao Yonashiro

## License
Apache 2.0
