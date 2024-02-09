package sort

import (
	"fmt"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	nums := []int{1, 5, 2, 4, 3, 45, 23, 4}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	fmt.Println(nums)
}
