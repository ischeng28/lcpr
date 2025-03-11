/*
 * @lc app=leetcode.cn id=16 lang=golang
 * @lcpr version=30204
 *
 * [16] 最接近的三数之和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	minDiff := math.MaxInt
	res := 0
	n := len(nums)
	for left := 0; left < n-2; left++ {
		if left > 0 && nums[left] == nums[left-1] {
			continue
		}
		sum := nums[left] + nums[left+1] + nums[left+2]
		if sum > target {
			if sum-target < minDiff {
				res = sum
			}
			break
		}

		sum = nums[left] + nums[n-1] + nums[n-2]
		if sum < target {
			if target-sum < minDiff {
				res = sum
				minDiff = target - sum
			}
			continue
		}

		mid, right := left+1, n-1
		for mid < right {
			sum = nums[left] + nums[mid] + nums[right]
			if sum == target {
				return target
			}
			if sum < target {
				if target-sum < minDiff {
					res = sum
					minDiff = target - sum
				}
				mid++
			} else {
				if sum-target < minDiff {
					res = sum
					minDiff = sum - target
				}
				right--
			}
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [-1,2,1,-4]\n1\n
// @lcpr case=end

// @lcpr case=start
// [0,0,0]\n1\n
// @lcpr case=end

*/

