package leetcode

import (
    "testing"
)

func TestStatisticsFromALargeSample(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func sampleStats(count []int) []float64 {
	var minimum, maximum, mean, median, mode float64
	for i := 0; i <= 255; i++ {
		if count[i] > 0 {
			minimum = float64(i)
			break
		}
	}

	for i := 255; i >= 0; i-- {
		if count[i] > 0 {
			maximum = float64(i)
			break
		}
	}

	var sum, n float64
	max := 0
	for i := 0; i <= 255; i++ {
		sum += float64(count[i] * i)
		n += float64(count[i])
		if count[i] > max {
			max = count[i]
			mode = float64(i)
		}
	}
	mean = sum / n
	l, r := 0, 255
	for t := (int(n) - 1) >> 1; t >= count[l]; {
		t -= count[l]
		l++
	}
	for t := (int(n) - 1) >> 1; t >= count[r]; {
		t -= count[r]
		r--
	}
	median = float64(l+r) / 2

	return []float64{
		minimum,
		maximum,
		mean,
		median,
		mode,
	}
}

//leetcode submit region end(Prohibit modification and deletion)
