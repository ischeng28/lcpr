/*
 * @lc app=leetcode.cn id=1441 lang=golang
 * @lcpr version=30204
 *
 * [1441] 用栈操作构建数组
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func buildArray(target []int, n int) []string {
	prev := 1
	ans := []string{}
	for _, v := range target {
		for v > prev {
			ans = append(ans, "Push", "Pop")
			prev++
		}
		ans = append(ans, "Push")
		prev++
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [1,3]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n4\n
// @lcpr case=end

*/

