/*
 * @lc app=leetcode.cn id=1652 lang=golang
 * @lcpr version=30204
 *
 * [1652] 拆炸弹
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func decrypt(code []int, k int) []int {
	n := len(code)
	ans := make([]int, n)
	r := k + 1
	if k < 0 {
		k = -k
		r = n
	}
	sum := 0
	for _, v := range code[r-k : r] {
		sum += v
	}
	for i, _ := range ans {
		ans[i] = sum
		sum += code[r%n] - code[(r-k)%n]
		r++
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [5,7,1,4]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4]\n0\n
// @lcpr case=end

// @lcpr case=start
// [2,4,9,3]\n-2\n
// @lcpr case=end

*/

