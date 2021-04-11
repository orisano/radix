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
