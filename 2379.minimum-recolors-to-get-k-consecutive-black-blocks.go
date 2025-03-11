/*
 * @lc app=leetcode.cn id=2379 lang=golang
 * @lcpr version=30204
 *
 * [2379] 得到 K 个黑块的最少涂色次数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minimumRecolors(blocks string, k int) int {
	res := 0
	blackNum := 0
	for i := 0; i < len(blocks); i++ {
		if blocks[i] == 'B' {
			blackNum++
		}
		if i < k-1 {
			continue
		}
		res = max(res, blackNum)
		if blocks[i-k+1] == 'B' {
			blackNum--
		}
	}
	return k - res
}

// @lc code=end

/*
// @lcpr case=start
// "WBBWWBBWBW"\n7\n
// @lcpr case=end

// @lcpr case=start
// "WBWBBBW"\n2\n
// @lcpr case=end

*/

