/*
 * @lc app=leetcode.cn id=18 lang=golang
 * @lcpr version=30204
 *
 * [18] 四数之和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := [][]int{}
	for a := 0; a < n-3; a++ {
		if a > 0 && nums[a-1] == nums[a] {
			continue
		}
		sum := nums[a] + nums[a+1] + nums[a+2] + nums[a+3]
		if sum > target {
			break
		}
		sum = nums[a] + nums[n-1] + nums[n-2] + nums[n-3]
		if sum < target {
			continue
		}
		for b := a + 1; b < n-2; b++ {
			if b > a+1 && nums[b-1] == nums[b] {
				continue
			}
			sum = nums[a] + nums[b] + nums[b+1] + nums[b+2]
			if sum > target {
				break
			}
			sum = nums[a] + nums[b] + nums[n-1] + nums[n-2]
			if sum < target {
				continue
			}
			c, d := b+1, n-1
			for c < d {
				sum = nums[a] + nums[b] + nums[c] + nums[d]
				if sum == target {
					res = append(res, []int{nums[a], nums[b], nums[c], nums[d]})
					for c = c + 1; c < d && nums[c] == nums[c-1]; c++ {
					}
					for d = d - 1; c < d && nums[d] == nums[d+1]; d-- {
					}
				} else if sum < target {
					c++
				} else {
					d--
				}

			}
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [1,0,-1,0,-2,2]\n0\n
// @lcpr case=end

// @lcpr case=start
// [2,2,2,2,2]\n8\n
// @lcpr case=end

*/

