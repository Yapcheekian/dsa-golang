package longestcommonprefix

func LongestCommonPrefix(strs []string) string {
	same := strs[0]
	for i := 1; i < len(strs); i++ {
		j := 0
		for ; j < len(strs[i]) && j < len(same); j++ {
			if same[j] != strs[i][j] {
				break
			}
		}
		same = same[:j]
	}
	return same
}
