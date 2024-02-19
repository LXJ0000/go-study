package main

import (
	"math/rand"
)

type RandomizedSet struct {
	nums []int
	mp   map[int]int
}

//func Constructor() RandomizedSet {
//	return RandomizedSet{
//		nums: []int{},
//		mp:   map[int]int{},
//	}
//}

func (r *RandomizedSet) Insert(val int) bool {
	if _, ok := r.mp[val]; !ok {
		r.mp[val] = len(r.nums)
		r.nums = append(r.nums, val)
		return true
	}
	return false
}

func (r *RandomizedSet) Remove(val int) bool {
	if idx, ok := r.mp[val]; ok {
		last := r.nums[len(r.nums)-1]

		r.mp[last] = idx
		r.nums[idx] = last

		r.nums = r.nums[:len(r.nums)-1]
		delete(r.mp, val)
		return true
	}
	return false
}

func (r *RandomizedSet) GetRandom() int {
	return r.nums[rand.Intn(len(r.nums))]
}
