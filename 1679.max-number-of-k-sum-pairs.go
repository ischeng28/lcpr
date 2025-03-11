/*
 * @lc app=leetcode.cn id=1679 lang=golang
 * @lcpr version=30204
 *
 * [1679] K 和数对的最大数目
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxOperations(nums []int, k int) int {
	mp := map[int]int{}
	res := 0
	for _, v := range nums {
		if mp[k-v] > 0 {
			res++
			mp[k-v]--
		} else {
			mp[v]++
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4]\n5\n
// @lcpr case=end

// @lcpr case=start
// [3,1,3,4,3]\n6\n
// @lcpr case=end

*/

