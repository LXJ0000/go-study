package main

func isAlienSorted(words []string, order string) bool {
	mp := map[byte]int{}
	for i, c := range order {
		mp[byte(c)] = i
	}
	cmp := func(w1, w2 string) bool {
		if w1 == w2 {
			return true
		}
		var i int
		for i = 0; i < len(w1) && i < len(w2); i++ {
			if mp[w1[i]] == mp[w2[i]] {
				continue
			}
			return mp[w1[i]] < mp[w2[i]]
		}
		if i < len(w2) {
			return true
		}
		return false
	}
	for i := 0; i < len(words)-1; i++ {
		if !cmp(words[i], words[i+1]) {
			return false
		}
	}
	return true
}
