/*
 * @lc app=leetcode.cn id=735 lang=golang
 * @lcpr version=30204
 *
 * [735] 小行星碰撞
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func asteroidCollision(asteroids []int) []int {
	stack := []int{}
	for _, v := range asteroids {
		if v > 0 {
			stack = append(stack, v)
			continue
		}
		if len(stack) == 0 || stack[len(stack)-1] < 0 {
			stack = append(stack, v)
			continue
		}
		for len(stack) > 0 && stack[len(stack)-1] > 0 && -v > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			stack = append(stack, v)
		} else if stack[len(stack)-1] < 0 {
			stack = append(stack, v)
		} else if stack[len(stack)-1] == -v {
			stack = stack[:len(stack)-1]
		} else {
			continue
		}
	}
	return stack
}

// @lc code=end

// @lc code=end

/*
// @lcpr case=start
// [5,10,-5]\n
// @lcpr case=end

// @lcpr case=start
// [8,-8]\n
// @lcpr case=end

// @lcpr case=start
// [10,2,-5]\n
// @lcpr case=end

*/

