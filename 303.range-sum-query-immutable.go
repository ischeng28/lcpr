/*
 * @lc app=leetcode.cn id=303 lang=golang
 * @lcpr version=30204
 *
 * [303] 区域和检索 - 数组不可变
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
type NumArray []int

func Constructor(nums []int) NumArray {
	arr := make([]int, len(nums)+1)
	for i, v := range nums {
		arr[i+1] = arr[i] + v
	}
	return arr
}

func (this *NumArray) SumRange(left int, right int) int {
	return (*this)[right+1] - (*this)[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
// @lc code=end

/*
// @lcpr case=start
// ["NumArray", "sumRange", "sumRange", "sumRange"][[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]\n
// @lcpr case=end

*/

