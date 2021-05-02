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

func SortInt32(a []int32) {
	int32Sorter{}.sort(a)
}

func SortInt64(a []int64) {
	int64Sorter{}.sort(a)
}

func SortUint32(a []uint32) {
	uint32Sorter{}.sort(a)
}

func SortUint64(a []uint64) {
	uint64Sorter{}.sort(a)
}

func SortFloat32(a []float32) {
	float32Sorter{}.sort(a)
}

func SortFloat64(a []float64) {
	float64Sorter{}.sort(a)
}

func ParallelSortInt32(a []int32, p int) {
	int32ParallelSorter{}.sort(a, p)
}

func ParallelSortInt64(a []int64, p int) {
	int64ParallelSorter{}.sort(a, p)
}

func ParallelSortUint32(a []uint32, p int) {
	uint32ParallelSorter{}.sort(a, p)
}

func ParallelSortUint64(a []uint64, p int) {
	uint64ParallelSorter{}.sort(a, p)
}

func ParallelSortFloat32(a []float32, p int) {
	float32ParallelSorter{}.sort(a, p)
}

func ParallelSortFloat64(a []float64, p int) {
	float64ParallelSorter{}.sort(a, p)
}

type parallelContext struct {
	begin, end int
	histo      [256]int
}

func parallelContexts(n, p int) []parallelContext {
	cs := make([]parallelContext, p)
	for i := 0; i < p-1; i++ {
		t := (n - cs[i].begin) / (p - i)
		end := cs[i].begin + t
		cs[i].end = end
		cs[i+1].begin = end
	}
	cs[p-1].end = n
	return cs
}
