/*
 * @lc app=leetcode.cn id=806 lang=golang
 * @lcpr version=30204
 *
 * [806] 写字符串需要的行数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func numberOfLines(widths []int, s string) []int {
	row := 1
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += widths[s[i]-'a']
		if sum > 100 {
			row++
			sum = widths[s[i]-'a']
		}
	}
	return []int{row, sum}

}

// @lc code=end

/*
// @lcpr case=start
// [10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10\n"abcdefghijklmnopqrstuvwxyz"\n
// @lcpr case=end

// @lcpr case=start
// [4,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10\n"bbbcccdddaaa"\n
// @lcpr case=end

*/

