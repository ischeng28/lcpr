/*
 * @lc app=leetcode.cn id=2535 lang=golang
 * @lcpr version=30204
 *
 * [2535] 数组元素和与数字和的绝对差
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func differenceOfSum(nums []int) (sum int) {
	sum = 0
	for _, num := range nums {
		sum += num
		for num > 0 {
			sum -= num % 10
			num = num / 10
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [1,15,6,3]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4]\n
// @lcpr case=end

*/

