/*
 * @lc app=leetcode.cn id=150 lang=golang
 * @lcpr version=30204
 *
 * [150] 逆波兰表达式求值
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func evalRPN(tokens []string) int {
	nums := []int{}
	calc := func(op string) {
		b := nums[len(nums)-1]
		nums = nums[:len(nums)-1]

		a := nums[len(nums)-1]
		nums = nums[:len(nums)-1]

		if op == "*" {
			nums = append(nums, a*b)
		} else if op == "/" {
			nums = append(nums, a/b)
		} else if op == "+" {
			nums = append(nums, a+b)
		} else if op == "-" {
			nums = append(nums, a-b)
		}
	}

	isOp := func(s string) bool {
		return s == "+" || s == "-" || s == "*" || s == "/"
	}

	for _, s := range tokens {
		if !isOp(s) {
			v, _ := strconv.Atoi(s)
			nums = append(nums, v)
		} else {
			calc(s)
		}
	}

	return nums[0]
}

// @lc code=end

/*
// @lcpr case=start
// ["2","1","+","3","*"]\n
// @lcpr case=end

// @lcpr case=start
// ["4","13","5","/","+"]\n
// @lcpr case=end

// @lcpr case=start
// ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]\n
// @lcpr case=end

*/

