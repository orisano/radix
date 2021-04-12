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
	"math"
	"runtime"
	"sync"
)

func encodeInt32(x int32) uint32 {
	return uint32(x - math.MinInt32)
}

func encodeInt64(x int64) uint64 {
	return uint64(x - math.MinInt64)
}

func encodeFloat32(x float32) uint32 {
	v := math.Float32bits(x)
	return v ^ (((^v >> 31) - 1) | 0x80000000)
}

func encodeFloat64(x float64) uint64 {
	v := math.Float64bits(x)
	return v ^ (((^v >> 63) - 1) | 0x8000000000000000)
}

func SortInt32(a []int32) {
	var level0, level1, level2, level3 [257]int
	b := make([]int32, len(a))
	for _, v := range a {
		x := encodeInt32(v)
		level0[((x>>0)&0xff)+1]++
		level1[((x>>8)&0xff)+1]++
		level2[((x>>16)&0xff)+1]++
		level3[((x>>24)&0xff)+1]++
	}
	for i := 1; i < 257; i++ {
		level0[i] += level0[i-1]
		level1[i] += level1[i-1]
		level2[i] += level2[i-1]
		level3[i] += level3[i-1]
	}
	for _, v := range a {
		x := (encodeInt32(v) >> 0) & 0xff
		b[level0[x]] = v
		level0[x]++
	}
	for _, v := range b {
		x := (encodeInt32(v) >> 8) & 0xff
		a[level1[x]] = v
		level1[x]++
	}
	for _, v := range a {
		x := (encodeInt32(v) >> 16) & 0xff
		b[level2[x]] = v
		level2[x]++
	}
	for _, v := range b {
		x := (encodeInt32(v) >> 24) & 0xff
		a[level3[x]] = v
		level3[x]++
	}
}

func SortUint32(a []uint32) {
	var level0, level1, level2, level3 [257]int
	b := make([]uint32, len(a))
	for _, v := range a {
		x := v
		level0[((x>>0)&0xff)+1]++
		level1[((x>>8)&0xff)+1]++
		level2[((x>>16)&0xff)+1]++
		level3[((x>>24)&0xff)+1]++
	}
	for i := 1; i < 257; i++ {
		level0[i] += level0[i-1]
		level1[i] += level1[i-1]
		level2[i] += level2[i-1]
		level3[i] += level3[i-1]
	}
	for _, v := range a {
		x := (v >> 0) & 0xff
		b[level0[x]] = v
		level0[x]++
	}
	for _, v := range b {
		x := (v >> 8) & 0xff
		a[level1[x]] = v
		level1[x]++
	}
	for _, v := range a {
		x := (v >> 16) & 0xff
		b[level2[x]] = v
		level2[x]++
	}
	for _, v := range b {
		x := (v >> 24) & 0xff
		a[level3[x]] = v
		level3[x]++
	}
}

func SortInt64(a []int64) {
	var level0, level1, level2, level3 [257]int
	var level4, level5, level6, level7 [257]int
	b := make([]int64, len(a))
	for _, v := range a {
		x := encodeInt64(v)
		level0[((x>>0)&0xff)+1]++
		level1[((x>>8)&0xff)+1]++
		level2[((x>>16)&0xff)+1]++
		level3[((x>>24)&0xff)+1]++
		level4[((x>>32)&0xff)+1]++
		level5[((x>>40)&0xff)+1]++
		level6[((x>>48)&0xff)+1]++
		level7[((x>>56)&0xff)+1]++
	}
	for i := 1; i < 257; i++ {
		level0[i] += level0[i-1]
		level1[i] += level1[i-1]
		level2[i] += level2[i-1]
		level3[i] += level3[i-1]
		level4[i] += level4[i-1]
		level5[i] += level5[i-1]
		level6[i] += level6[i-1]
		level7[i] += level7[i-1]
	}
	for _, v := range a {
		x := (encodeInt64(v) >> 0) & 0xff
		b[level0[x]] = v
		level0[x]++
	}
	for _, v := range b {
		x := (encodeInt64(v) >> 8) & 0xff
		a[level1[x]] = v
		level1[x]++
	}
	for _, v := range a {
		x := (encodeInt64(v) >> 16) & 0xff
		b[level2[x]] = v
		level2[x]++
	}
	for _, v := range b {
		x := (encodeInt64(v) >> 24) & 0xff
		a[level3[x]] = v
		level3[x]++
	}
	for _, v := range a {
		x := (encodeInt64(v) >> 32) & 0xff
		b[level4[x]] = v
		level4[x]++
	}
	for _, v := range b {
		x := (encodeInt64(v) >> 40) & 0xff
		a[level5[x]] = v
		level5[x]++
	}
	for _, v := range a {
		x := (encodeInt64(v) >> 48) & 0xff
		b[level6[x]] = v
		level6[x]++
	}
	for _, v := range b {
		x := (encodeInt64(v) >> 56) & 0xff
		a[level7[x]] = v
		level7[x]++
	}
}

func SortUint64(a []uint64) {
	var level0, level1, level2, level3 [257]int
	var level4, level5, level6, level7 [257]int
	b := make([]uint64, len(a))
	for _, v := range a {
		x := v
		level0[((x>>0)&0xff)+1]++
		level1[((x>>8)&0xff)+1]++
		level2[((x>>16)&0xff)+1]++
		level3[((x>>24)&0xff)+1]++
		level4[((x>>32)&0xff)+1]++
		level5[((x>>40)&0xff)+1]++
		level6[((x>>48)&0xff)+1]++
		level7[((x>>56)&0xff)+1]++
	}
	for i := 1; i < 257; i++ {
		level0[i] += level0[i-1]
		level1[i] += level1[i-1]
		level2[i] += level2[i-1]
		level3[i] += level3[i-1]
		level4[i] += level4[i-1]
		level5[i] += level5[i-1]
		level6[i] += level6[i-1]
		level7[i] += level7[i-1]
	}
	for _, v := range a {
		x := (v >> 0) & 0xff
		b[level0[x]] = v
		level0[x]++
	}
	for _, v := range b {
		x := (v >> 8) & 0xff
		a[level1[x]] = v
		level1[x]++
	}
	for _, v := range a {
		x := (v >> 16) & 0xff
		b[level2[x]] = v
		level2[x]++
	}
	for _, v := range b {
		x := (v >> 24) & 0xff
		a[level3[x]] = v
		level3[x]++
	}
	for _, v := range a {
		x := (v >> 32) & 0xff
		b[level4[x]] = v
		level4[x]++
	}
	for _, v := range b {
		x := (v >> 40) & 0xff
		a[level5[x]] = v
		level5[x]++
	}
	for _, v := range a {
		x := (v >> 48) & 0xff
		b[level6[x]] = v
		level6[x]++
	}
	for _, v := range b {
		x := (v >> 56) & 0xff
		a[level7[x]] = v
		level7[x]++
	}
}

func SortFloat32(a []float32) {
	var level0, level1, level2, level3 [257]int
	b := make([]float32, len(a))
	for _, v := range a {
		x := encodeFloat32(v)
		level0[((x>>0)&0xff)+1]++
		level1[((x>>8)&0xff)+1]++
		level2[((x>>16)&0xff)+1]++
		level3[((x>>24)&0xff)+1]++
	}
	for i := 1; i < 257; i++ {
		level0[i] += level0[i-1]
		level1[i] += level1[i-1]
		level2[i] += level2[i-1]
		level3[i] += level3[i-1]
	}
	for _, v := range a {
		x := (encodeFloat32(v) >> 0) & 0xff
		b[level0[x]] = v
		level0[x]++
	}
	for _, v := range b {
		x := (encodeFloat32(v) >> 8) & 0xff
		a[level1[x]] = v
		level1[x]++
	}
	for _, v := range a {
		x := (encodeFloat32(v) >> 16) & 0xff
		b[level2[x]] = v
		level2[x]++
	}
	for _, v := range b {
		x := (encodeFloat32(v) >> 24) & 0xff
		a[level3[x]] = v
		level3[x]++
	}
}

func SortFloat64(a []float64) {
	var level0, level1, level2, level3 [257]int
	var level4, level5, level6, level7 [257]int
	b := make([]float64, len(a))
	for _, v := range a {
		x := encodeFloat64(v)
		level0[((x>>0)&0xff)+1]++
		level1[((x>>8)&0xff)+1]++
		level2[((x>>16)&0xff)+1]++
		level3[((x>>24)&0xff)+1]++
		level4[((x>>32)&0xff)+1]++
		level5[((x>>40)&0xff)+1]++
		level6[((x>>48)&0xff)+1]++
		level7[((x>>56)&0xff)+1]++
	}
	for i := 1; i < 257; i++ {
		level0[i] += level0[i-1]
		level1[i] += level1[i-1]
		level2[i] += level2[i-1]
		level3[i] += level3[i-1]
		level4[i] += level4[i-1]
		level5[i] += level5[i-1]
		level6[i] += level6[i-1]
		level7[i] += level7[i-1]
	}
	for _, v := range a {
		x := (encodeFloat64(v) >> 0) & 0xff
		b[level0[x]] = v
		level0[x]++
	}
	for _, v := range b {
		x := (encodeFloat64(v) >> 8) & 0xff
		a[level1[x]] = v
		level1[x]++
	}
	for _, v := range a {
		x := (encodeFloat64(v) >> 16) & 0xff
		b[level2[x]] = v
		level2[x]++
	}
	for _, v := range b {
		x := (encodeFloat64(v) >> 24) & 0xff
		a[level3[x]] = v
		level3[x]++
	}
	for _, v := range a {
		x := (encodeFloat64(v) >> 32) & 0xff
		b[level4[x]] = v
		level4[x]++
	}
	for _, v := range b {
		x := (encodeFloat64(v) >> 40) & 0xff
		a[level5[x]] = v
		level5[x]++
	}
	for _, v := range a {
		x := (encodeFloat64(v) >> 48) & 0xff
		b[level6[x]] = v
		level6[x]++
	}
	for _, v := range b {
		x := (encodeFloat64(v) >> 56) & 0xff
		a[level7[x]] = v
		level7[x]++
	}
}

type parallelSorter struct {
	begin, end int
	histo      [256]int
}

func (ps *parallelSorter) computeHistogram(b uint8, src []int32) {
	for i := range ps.histo {
		ps.histo[i] = 0
	}
	chunk := src[ps.begin:ps.end]
	for i := range chunk {
		s := encodeInt32(chunk[i])
		t := uint8((s >> b) & 0xff)
		ps.histo[t]++
	}
}

type smallSlice struct {
	array [32]int32
	n     int
}

func (s *smallSlice) Copy(dst []int32) {
	copy(dst[:s.n], s.array[:s.n])
}

func (ps *parallelSorter) scatter(b uint8, src, dst []int32) {
	var bufs [256]smallSlice
	chunk := src[ps.begin:ps.end]
	for i := range chunk {
		s := encodeInt32(chunk[i])
		t := uint8((s >> b) & 0xff)

		n := bufs[t].n
		bufs[t].array[n] = chunk[i]
		bufs[t].n = n + 1
		if bufs[t].n == 32 {
			p := ps.histo[t]
			dst := dst[p : p+32 : p+32]
			dst[0] = bufs[t].array[0]
			dst[1] = bufs[t].array[1]
			dst[2] = bufs[t].array[2]
			dst[3] = bufs[t].array[3]
			dst[4] = bufs[t].array[4]
			dst[5] = bufs[t].array[5]
			dst[6] = bufs[t].array[6]
			dst[7] = bufs[t].array[7]
			dst[8] = bufs[t].array[8]
			dst[9] = bufs[t].array[9]
			dst[10] = bufs[t].array[10]
			dst[11] = bufs[t].array[11]
			dst[12] = bufs[t].array[12]
			dst[13] = bufs[t].array[13]
			dst[14] = bufs[t].array[14]
			dst[15] = bufs[t].array[15]
			dst[16] = bufs[t].array[16]
			dst[17] = bufs[t].array[17]
			dst[18] = bufs[t].array[18]
			dst[19] = bufs[t].array[19]
			dst[20] = bufs[t].array[20]
			dst[21] = bufs[t].array[21]
			dst[22] = bufs[t].array[22]
			dst[23] = bufs[t].array[23]
			dst[24] = bufs[t].array[24]
			dst[25] = bufs[t].array[25]
			dst[26] = bufs[t].array[26]
			dst[27] = bufs[t].array[27]
			dst[28] = bufs[t].array[28]
			dst[29] = bufs[t].array[29]
			dst[30] = bufs[t].array[30]
			dst[31] = bufs[t].array[31]
			ps.histo[t] += 32
			bufs[t].n = 0
		}
	}
	for i := 0; i < 256; i++ {
		i := uint8(i)
		p := ps.histo[i]
		copy(dst[p:], bufs[i].array[:bufs[i].n])
	}
}

func ParallelSort(a []int32) {
	p := runtime.GOMAXPROCS(0)
	n := len(a)
	ss := make([]parallelSorter, p)
	for i := 0; i < p-1; i++ {
		t := (n - ss[i].begin) / (p - i)
		end := ss[i].begin + t
		ss[i].end = end
		ss[i+1].begin = end
	}
	ss[p-1].end = n

	b := make([]int32, n)
	for i := 0; i < 4; i++ {
		shift := uint8(i * 8)
		var histoWg sync.WaitGroup
		for j := range ss {
			histoWg.Add(1)
			go func(j int) {
				ss[j].computeHistogram(shift, a)
				histoWg.Done()
			}(j)
		}
		histoWg.Wait()
		s := 0
		for j := 0; j < 256; j++ {
			j := uint8(j)
			for k := range ss {
				t := s + ss[k].histo[j]
				ss[k].histo[j] = s
				s = t
			}
		}
		var scatterWg sync.WaitGroup
		for j := range ss {
			scatterWg.Add(1)
			go func(j int) {
				ss[j].scatter(shift, a, b)
				scatterWg.Done()
			}(j)
		}
		scatterWg.Wait()
		a, b = b, a
	}
}
