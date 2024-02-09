package sort

import (
	"fmt"
	"slices"
	"testing"
)

func TestSort(t *testing.T) {
	nums := []int{1, 9, 2, 6, 4, 6, 3}
	slices.Sort(nums)
	fmt.Println(nums)
}

func Sum(nums []int) (res [][]int) {

	// 保证a <= b <= c
	slices.Sort(nums)
	n := len(nums)
	st := map[int]struct{}{}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			k := 0 - nums[i] - nums[j]
			if _, ok := st[k]; ok {
				res = append(res, []int{nums[i], nums[j], k})
			}
		}
		st[nums[i]] = struct{}{}
	}
	return
}
