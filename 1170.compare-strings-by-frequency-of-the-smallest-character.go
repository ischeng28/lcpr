/*
 * @lc app=leetcode.cn id=1170 lang=golang
 * @lcpr version=30204
 *
 * [1170] 比较字符串最小字母出现频次
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func numSmallerByFrequency(queries []string, words []string) []int {
	count := make([]int, len(words))
	for i, s := range words {
		count[i] = getFreq(s)
	}
	sort.Ints(count)
	res := make([]int, len(queries))
	for i, v := range queries {
		freq := getFreq(v)
		res[i] = len(count) - searchInts(count, freq)
	}
	return res
}

func searchInts(nums []int, target int) int {
	left, right := -1, len(nums)
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid
		} else {
			right = mid
		}
	}
	return right
}

func getFreq(s string) int {
	minChar := s[0]
	res := 1
	for i := 1; i < len(s); i++ {
		if s[i] < minChar {
			res = 1
			minChar = s[i]
		} else if s[i] == minChar {
			res++
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// ["cbd"]\n["zaaaz"]\n
// @lcpr case=end

// @lcpr case=start
// ["bbb","cc"]\n["a","aa","aaa","aaaa"]\n
// @lcpr case=end

*/

