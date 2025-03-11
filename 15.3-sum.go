/*
 * @lc app=leetcode.cn id=15 lang=golang
 * @lcpr version=30204
 *
 * [15] 三数之和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func threeSum(nums []int) (res [][]int) {
	n := len(nums)
	if n < 3 {
		return res
	}
	sort.Ints(nums)
	for left := 0; left < n-2; left++ {
		if left > 0 && nums[left] == nums[left-1] {
			continue
		}
		mid, right := left+1, n-1
		for mid < right {
			sum := nums[left] + nums[mid] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[left], nums[mid], nums[right]})
				for mid++; mid < right && nums[mid] == nums[mid-1]; mid++ {
				}
				for right--; mid < right && nums[right] == nums[right+1]; right-- {
				}
			} else if sum < 0 {
				mid++
			} else {
				right--
			}
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [-1,0,1,2,-1,-4]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,1]\n
// @lcpr case=end

// @lcpr case=start
// [0,0,0]\n
// @lcpr case=end

*/

