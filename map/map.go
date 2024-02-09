package main

import "fmt"

func main() {
	mp := map[int]int{
		1: 10,
		2: 20,
		3: 30,
	}
	k := mp
	k[1] = 11
	for key, val := range mp {
		fmt.Println(key, val)
	}
}
