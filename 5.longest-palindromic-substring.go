package leetcode

// https://leetcode-cn.com/problems/longest-palindromic-substring/
// 最长回文子串

func longestPalindrome(s string) string {

	if len(s) < 2 {
		return s
	}

	var start, max int
	max = 1

	size := len(s)
	dp := make([][]bool, size)

	for i := range dp {
		dp[i] = make([]bool, size)
		dp[i][i] = true
	}

	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if s[i] == s[j] {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] {
				if j-i+1 > max {
					max = j - i + 1
					start = i
				}
			}
		}
	}

	return s[start : start+max]
}
