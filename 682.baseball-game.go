/*
 * @lc app=leetcode.cn id=682 lang=golang
 * @lcpr version=30204
 *
 * [682] 棒球比赛
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func calPoints(operations []string) int {
	stack := []int{}
	for _, v := range operations {
		if v == "+" {
			stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
		} else if v == "D" {
			stack = append(stack, stack[len(stack)-1]*2)
		} else if v == "C" {
			stack = stack[:len(stack)-1]
		} else {
			num, _ := strconv.Atoi(v)
			stack = append(stack, num)
		}
	}
	sum := 0
	for _, v := range stack {
		sum += v
	}
	return sum
}

// @lc code=end

/*
// @lcpr case=start
// ["5","2","C","D","+"]\n
// @lcpr case=end

// @lcpr case=start
// ["5","-2","4","C","D","9","+","+"]\n
// @lcpr case=end

// @lcpr case=start
// ["1"]\n
// @lcpr case=end

*/

