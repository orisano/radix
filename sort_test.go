//   Copyright 2021 Nao Yonashiro
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package radix

import (
	"math/rand"
	"runtime"
	"sort"
	"strconv"
	"testing"
)

func TestSortInt32(t *testing.T) {
	rand.Seed(72)
	x := make([]int32, 100000)
	for i := range x {
		x[i] = int32(rand.Uint32())
	}
	SortInt32(x)
	if !sort.SliceIsSorted(x, func(i, j int) bool {
		return x[i] < x[j]
	}) {
		t.Error("x is not sorted")
	}
}

func TestParallelSortInt32(t *testing.T) {
	rand.Seed(72)
	x := make([]int32, 100000)
	for i := range x {
		x[i] = int32(rand.Uint32())
	}
	ParallelSortInt32(x, runtime.GOMAXPROCS(0))
	if !sort.SliceIsSorted(x, func(i, j int) bool {
		return x[i] < x[j]
	}) {
		t.Error("x is not sorted")
	}
}

func TestSortInt64(t *testing.T) {
	rand.Seed(72)
	x := make([]int64, 100000)
	for i := range x {
		x[i] = int64(rand.Uint64())
	}
	SortInt64(x)
	if !sort.SliceIsSorted(x, func(i, j int) bool {
		return x[i] < x[j]
	}) {
		t.Error("x is not sorted")
	}
}

func TestSortFloat32(t *testing.T) {
	rand.Seed(72)
	x := make([]float32, 100000)
	for i := range x {
		x[i] = rand.Float32()
	}
	SortFloat32(x)
	if !sort.SliceIsSorted(x, func(i, j int) bool {
		return x[i] < x[j]
	}) {
		t.Error("x is not sorted")
	}
}

func TestSortFloat64(t *testing.T) {
	rand.Seed(72)
	x := make([]float64, 100000)
	for i := range x {
		x[i] = rand.NormFloat64()
	}
	SortFloat64(x)
	if !sort.SliceIsSorted(x, func(i, j int) bool {
		return x[i] < x[j]
	}) {
		t.Error("x is not sorted")
	}
}

func BenchmarkSortInt32(b *testing.B) {
	sizes := []int{
		1000,
		10000,
		100000,
		1000000,
		10000000,
	}
	for _, size := range sizes {
		rand.Seed(72)
		x := make([]int32, size)
		for i := range x {
			x[i] = int32(rand.Uint32())
		}
		y := make([]int32, size)
		b.Run("parallel-radix/"+strconv.Itoa(size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				copy(y, x)
				b.StartTimer()
				ParallelSortInt32(y, runtime.GOMAXPROCS(0))
			}
		})
		b.Run("radix/"+strconv.Itoa(size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				copy(y, x)
				b.StartTimer()
				SortInt32(y)
			}
		})
		b.Run("std/"+strconv.Itoa(size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				copy(y, x)
				b.StartTimer()
				sort.Slice(y, func(i, j int) bool {
					return y[i] < y[j]
				})
			}
		})
	}
}

func BenchmarkSortInt64(b *testing.B) {
	sizes := []int{
		1000,
		10000,
		100000,
		1000000,
	}
	for _, size := range sizes {
		rand.Seed(72)
		x := make([]int64, size)
		for i := range x {
			x[i] = int64(rand.Uint64())
		}
		y := make([]int64, size)
		b.Run("radix/"+strconv.Itoa(size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				copy(y, x)
				b.StartTimer()
				SortInt64(y)
			}
		})
		b.Run("std/"+strconv.Itoa(size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				copy(y, x)
				b.StartTimer()
				sort.Slice(y, func(i, j int) bool {
					return y[i] < y[j]
				})
			}
		})
	}
}

func BenchmarkSortFloat64(b *testing.B) {
	sizes := []int{
		1000,
		10000,
		100000,
		1000000,
	}
	for _, size := range sizes {
		rand.Seed(72)
		x := make([]float64, size)
		for i := range x {
			x[i] = rand.NormFloat64()
		}
		y := make([]float64, size)
		b.Run("radix/"+strconv.Itoa(size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				copy(y, x)
				b.StartTimer()
				SortFloat64(y)
			}
		})
		b.Run("std/"+strconv.Itoa(size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				copy(y, x)
				b.StartTimer()
				sort.Float64s(y)
			}
		})
	}
}
