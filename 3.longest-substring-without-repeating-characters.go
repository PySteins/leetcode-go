package leetcode

import "strings"

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
// 无重复字符的最长子串

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	var l, r, max int
	for ; r < len(s); r++ {

		// 查询窗口内有没有重复字符串
		if index := strings.IndexByte(s[l:r], s[r]); index != -1 {
			l += index + 1
		}

		if r-l+1 > max {
			max = r - l + 1
		}
	}

	return max
}
