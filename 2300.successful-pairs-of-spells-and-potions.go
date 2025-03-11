/*
 * @lc app=leetcode.cn id=2300 lang=golang
 * @lcpr version=30204
 *
 * [2300] 咒语和药水的成功对数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	for i, v := range spells {
		p := (v + int(success) - 1) / v
		po := searchLowerBound(potions, p)
		spells[i] = len(potions) - po
	}
	return spells
}

func searchLowerBound(nums []int, target int) int {
	left, right := -1, len(nums)
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}
	return right
}

// @lc code=end

/*
// @lcpr case=start
// [5,1,3]\n[1,2,3,4,5]\n7\n
// @lcpr case=end

// @lcpr case=start
// [3,1,2]\n[8,5,8]\n16\n
// @lcpr case=end

*/

