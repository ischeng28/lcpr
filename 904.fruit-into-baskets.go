/*
 * @lc app=leetcode.cn id=904 lang=golang
 * @lcpr version=30204
 *
 * [904] 水果成篮
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func totalFruit(fruits []int) (ans int) {
	mp := map[int]int{}
	left := 0
	for right, _ := range fruits {
		mp[fruits[right]]++
		for len(mp) > 2 {
			mp[fruits[left]]--
			if mp[fruits[left]] == 0 {
				delete(mp, fruits[left])
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,2,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,2,2]\n
// @lcpr case=end

// @lcpr case=start
// [3,3,3,1,2,1,1,2,3,3,4]\n
// @lcpr case=end

*/

