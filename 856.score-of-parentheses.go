/*
 * @lc app=leetcode.cn id=856 lang=golang
 * @lcpr version=30204
 *
 * [856] 括号的分数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func scoreOfParentheses(s string) int {
    st := []int{0}
    for _, c := range s {
        if c == '(' {
            st = append(st, 0)
        } else {
            v := st[len(st)-1]
            st = st[:len(st)-1]
            st[len(st)-1] += max(2*v, 1)
        }
    }
    return st[0]
}
// @lc code=end

/*
// @lcpr case=start
// "()"\n
// @lcpr case=end

// @lcpr case=start
// "(())"\n
// @lcpr case=end

// @lcpr case=start
// "()()"\n
// @lcpr case=end

// @lcpr case=start
// "(()(()))"\n
// @lcpr case=end

*/

