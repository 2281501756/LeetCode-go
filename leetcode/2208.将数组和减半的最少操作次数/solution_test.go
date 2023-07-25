package leetcode

import (
	"container/heap"
	"testing"
)

func TestMinimumOperationsToHalveArraySum(t *testing.T) {

}

//leetcode submit region begin(Prohibit modification and deletion)

type PriorityQueue []float64

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Pop() any {
	old, n := *pq, len(*pq)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}
func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(float64))
}

func halveArray(nums []int) int {
	pq := &PriorityQueue{}
	sum1, sum2 := 0.0, 0.0
	for i := 0; i < len(nums); i++ {
		sum1 += float64(nums[i])
		heap.Push(pq, float64(nums[i]))
	}
	res := 0
	for sum2 < sum1/2 {
		x := heap.Pop(pq).(float64)
		sum2 += x / 2
		heap.Push(pq, x/2)
		res++
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
