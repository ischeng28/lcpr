/*
 * @lc app=leetcode.cn id=1006 lang=golang
 * @lcpr version=30204
 *
 * [1006] 笨阶乘
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func clumsy(N int) int {
	nums := []int{}
	ops := []byte{}
	mp := map[byte]int{
		'*': 2,
		'/': 2,
		'+': 1,
		'-': 1,
	}

	cal := func() {
		b := nums[len(nums)-1]
		nums = nums[:len(nums)-1]

		a := nums[len(nums)-1]
		nums = nums[:len(nums)-1]

		op := ops[len(ops)-1]
		ops = ops[:len(ops)-1]

		ans := 0
		if op == '*' {
			ans = a * b
		} else if op == '/' {
			ans = a / b
		} else if op == '+' {
			ans = a + b
		} else if op == '-' {
			ans = a - b
		}
		nums = append(nums, ans)
	}

	cs := []byte{'*', '/', '+', '-'}
	j := 0
	for i := N; i > 0; i-- {
		op := cs[j%4]
		nums = append(nums, i)
		for len(ops) > 0 && mp[ops[len(ops)-1]] >= mp[op] {
			cal()
		}
		if i != 1 {
			ops = append(ops, op)
		}
		j++

	}
	for len(ops) > 0 {
		cal()
	}
	return nums[0]
}

// @lc code=end

/*
// @lcpr case=start
// 4\n
// @lcpr case=end

// @lcpr case=start
// 10\n
// @lcpr case=end

*/

