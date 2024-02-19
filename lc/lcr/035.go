package main

import (
	"math"
	"slices"
)

func findMinDifference(timePoints []string) (res int) {
	getTime := func(t string) int {
		return int(((t[0]-'0')*10+t[1]-'0')*60 + (t[3]-'0')*10 + t[4] - '0')
	}
	var tr []int
	for _, t := range timePoints {
		k := getTime(t)
		tr = append(tr, k)
	}
	slices.Sort(tr)
	res = math.MaxInt32
	res = min(res, int(math.Abs(float64(tr[0]-tr[len(tr)-1]))))
	for i := 0; i < len(tr)-1; i++ {
		res = min(res, tr[i+1]-tr[i])
	}
	return
}
