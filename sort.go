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

func SortInt32(a []int32) {
	var level0, level1, level2, level3 [257]int
	b := make([]int32, len(a))
	for _, v := range a {
		x := uint32(v - math.MinInt32)
		level0[(x&0xff)+1]++
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
		x := (v - math.MinInt32) & 0xff
		b[level0[x]] = v
		level0[x]++
	}
	for _, v := range b {
		x := (uint32(v-math.MinInt32) >> 8) & 0xff
		a[level1[x]] = v
		level1[x]++
	}
	for _, v := range a {
		x := (uint32(v-math.MinInt32) >> 16) & 0xff
		b[level2[x]] = v
		level2[x]++
	}
	for _, v := range b {
		x := (uint32(v-math.MinInt32) >> 24) & 0xff
		a[level3[x]] = v
		level3[x]++
	}
}

func SortInt64(a []int64) {
	var level0, level1, level2, level3 [257]int
	var level4, level5, level6, level7 [257]int
	b := make([]int64, len(a))
	for _, v := range a {
		x := uint64(v - math.MinInt64)
		level0[(x&0xff)+1]++
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
		x := (v - math.MinInt64) & 0xff
		b[level0[x]] = v
		level0[x]++
	}
	for _, v := range b {
		x := (uint64(v-math.MinInt64) >> 8) & 0xff
		a[level1[x]] = v
		level1[x]++
	}
	for _, v := range a {
		x := (uint64(v-math.MinInt64) >> 16) & 0xff
		b[level2[x]] = v
		level2[x]++
	}
	for _, v := range b {
		x := (uint64(v-math.MinInt64) >> 24) & 0xff
		a[level3[x]] = v
		level3[x]++
	}
	for _, v := range a {
		x := (uint64(v-math.MinInt64) >> 32) & 0xff
		b[level4[x]] = v
		level4[x]++
	}
	for _, v := range b {
		x := (uint64(v-math.MinInt64) >> 40) & 0xff
		a[level5[x]] = v
		level5[x]++
	}
	for _, v := range a {
		x := (uint64(v-math.MinInt64) >> 48) & 0xff
		b[level6[x]] = v
		level6[x]++
	}
	for _, v := range b {
		x := (uint64(v-math.MinInt64) >> 56) & 0xff
		a[level7[x]] = v
		level7[x]++
	}
}
