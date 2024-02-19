package main

import (
	"cmp"
	"fmt"
	"slices"
	"sort"
)

func main() {
	a := []int{1, 6, 5, 2, 7, 4, 1, 4}
	b := "industrialising"
	c := []float64{1.1, 123, 345, 123, 234, 345, 123, 23421}

	slices.Sort(a)
	fmt.Println(a)

	a = []int{1, 6, 5, 2, 7, 4, 1, 4}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	fmt.Println(a)

	sort.Ints(a)

	slices.SortFunc(c, func(a, b float64) int {
		return cmp.Compare(a, b)
		// a >  b  1
		// a == b  0
		// a <  b -1
	})

	newB := []byte(b)
	slices.Sort(newB)
	fmt.Println(string(newB))

}

func groupAnagrams(strs []string) (res [][]string) {
	mp := map[string]int{}
	for _, v := range strs {
		newStr := []byte(v)
		slices.Sort(newStr)
		if idx, ok := mp[string(newStr)]; ok {
			res[idx] = append(res[idx], v)
		} else {
			mp[string(newStr)] = len(res)
			res = append(res, []string{v})
		}
	}
	return
}
