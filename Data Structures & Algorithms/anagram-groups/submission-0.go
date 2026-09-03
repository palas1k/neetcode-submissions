func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, w := range strs {
		s := []byte(w)
		sort.Slice(s, func(i, j int) bool {return s[i] < s[j] })
		m[string(s)] = append(m[string(s)], w)
	}

	res := make([][]string, 0, len(m))
	for _, v := range(m) {
		res = append(res, v)
	}
	return res
}
