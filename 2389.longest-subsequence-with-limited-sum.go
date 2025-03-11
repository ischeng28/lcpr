/*
 * @lc app=leetcode.cn id=2389 lang=golang
 * @lcpr version=30204
 *
 * [2389] 和有限的最长子序列
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func answerQueries(nums []int, queries []int) []int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	for i, v := range queries {
		index := searchInts(nums, v+1)
		queries[i] = index
	}
	return queries
}

func searchInts(nums []int, target int) int {
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
// [4,5,2,1]\n[3,10,21]\n
// @lcpr case=end

// @lcpr case=start
// [2,3,4,5]\n[1]\n
// @lcpr case=end

*/

