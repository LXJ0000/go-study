package main

func lengthOfLongestSubstring(s string) (res int) {
	n := len(s)
	st := map[byte]struct{}{}
	last := 0
	for i := 0; i < n; i++ {
		for {
			if _, ok := st[s[i]-'a']; ok {
				delete(st, s[last]-'a')
				// st[s[last] - 'a'] --
				last++
			} else {
				break
			}
		}

		st[s[i]-'a'] = struct{}{}
		res = max(res, i-last+1)
	}
	return
}
