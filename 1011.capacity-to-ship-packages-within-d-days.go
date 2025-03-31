/*
 * @lc app=leetcode.cn id=1011 lang=golang
 * @lcpr version=30204
 *
 * [1011] 在 D 天内送达包裹的能力
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func shipWithinDays(weights []int, days int) int {
	check := func(m int) bool {
		count := 0
		sum := 0
		for _, v := range weights {
			sum += v
			if sum > m {
				count++
				if count > days {
					return false
				}
				sum = v
			}
		}
		count++
		return count <= days
	}
	totalSum := 0
	for _, v := range weights {
		totalSum += v
	}

	left, right := slices.Max(weights)-1, totalSum
	for left+1 < right {
		mid := left + (right-left)/2
		if check(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	return right

}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5,6,7,8,9,10]\n5\n
// @lcpr case=end

// @lcpr case=start
// [3,2,2,4,1,4]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,1,1]\n4\n
// @lcpr case=end

*/

