package radix

import (
	"math"
	"sync"
)

type float64Sorter struct{}

func (float64Sorter) encode64(x float64) uint64 {
	v := math.Float64bits(x)
	return v ^ (((^v >> 63) - 1) | 0x8000000000000000)
}

func (s float64Sorter) sort(a []float64) {
	b := make([]float64, len(a))

	var level0, level1, level2, level3 [256]int
	var level4, level5, level6, level7 [256]int
	for _, v := range a {
		x := s.encode64(v)
		level0[uint8((x>>0)&0xff)]++
		level1[uint8((x>>8)&0xff)]++
		level2[uint8((x>>16)&0xff)]++
		level3[uint8((x>>24)&0xff)]++
		level4[uint8((x>>32)&0xff)]++
		level5[uint8((x>>40)&0xff)]++
		level6[uint8((x>>48)&0xff)]++
		level7[uint8((x>>56)&0xff)]++
	}
	var l0, l1, l2, l3 int
	var l4, l5, l6, l7 int
	for i := 0; i < 256; i++ {
		i := uint8(i)
		l0, level0[i] = l0+level0[i], l0
		l1, level1[i] = l1+level1[i], l1
		l2, level2[i] = l2+level2[i], l2
		l3, level3[i] = l3+level3[i], l3
		l4, level4[i] = l4+level4[i], l4
		l5, level5[i] = l5+level5[i], l5
		l6, level6[i] = l6+level6[i], l6
		l7, level7[i] = l7+level7[i], l7
	}
	for _, v := range a {
		x := uint8((s.encode64(v) >> 0) & 0xff)
		b[level0[x]] = v
		level0[x]++
	}
	for _, v := range b {
		x := uint8((s.encode64(v) >> 8) & 0xff)
		a[level1[x]] = v
		level1[x]++
	}
	for _, v := range a {
		x := uint8((s.encode64(v) >> 16) & 0xff)
		b[level2[x]] = v
		level2[x]++
	}
	for _, v := range b {
		x := uint8((s.encode64(v) >> 24) & 0xff)
		a[level3[x]] = v
		level3[x]++
	}
	for _, v := range a {
		x := uint8((s.encode64(v) >> 32) & 0xff)
		b[level4[x]] = v
		level4[x]++
	}
	for _, v := range b {
		x := uint8((s.encode64(v) >> 40) & 0xff)
		a[level5[x]] = v
		level5[x]++
	}
	for _, v := range a {
		x := uint8((s.encode64(v) >> 48) & 0xff)
		b[level6[x]] = v
		level6[x]++
	}
	for _, v := range b {
		x := uint8((s.encode64(v) >> 56) & 0xff)
		a[level7[x]] = v
		level7[x]++
	}
}

type float64ParallelSorter struct {
	float64Sorter
}

func (s float64ParallelSorter) sort(a []float64, p int) {
	b := make([]float64, len(a))
	type smallSlice struct {
		array [32]float64
		n     int
	}
	cs := parallelContexts(len(a), p)
	for step := 0; step < 8; step++ {
		shift := uint8(step * 8)
		var histoWg sync.WaitGroup
		for j := range cs {
			histoWg.Add(1)
			go func(c *parallelContext) {
				for i := range c.histo {
					c.histo[i] = 0
				}
				for _, x := range a[c.begin:c.end] {
					c.histo[uint8((s.encode64(x)>>shift)&0xff)]++
				}
				histoWg.Done()
			}(&cs[j])
		}
		histoWg.Wait()
		acc := 0
		for j := 0; j < 256; j++ {
			j := uint8(j)
			for k := range cs {
				acc, cs[k].histo[j] = acc+cs[k].histo[j], acc
			}
		}
		var scatterWg sync.WaitGroup
		for j := range cs {
			scatterWg.Add(1)
			go func(c *parallelContext) {
				var bufs [256]smallSlice
				for _, x := range a[c.begin:c.end] {
					t := uint8((s.encode64(x) >> shift) & 0xff)
					n := bufs[t].n
					bufs[t].array[n] = x
					bufs[t].n = n + 1
					if bufs[t].n == 32 {
						from := c.histo[t]
						dst := b[from : from+32 : from+32]
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
						c.histo[t] += 32
						bufs[t].n = 0
					}
				}
				for i := 0; i < 256; i++ {
					i := uint8(i)
					from := c.histo[i]
					copy(b[from:], bufs[i].array[:bufs[i].n])
				}
				scatterWg.Done()
			}(&cs[j])
		}
		scatterWg.Wait()
		a, b = b, a
	}
}
