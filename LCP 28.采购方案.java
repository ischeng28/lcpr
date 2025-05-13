/*
 * @lc app=leetcode.cn id=LCP 28 lang=java
 * @lcpr version=30204
 *
 * [LCP 28] 采购方案
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
class Solution {
    public int purchasePlans(int[] nums, int target) {
        int res = 0;
        Arrays.sort(nums);
        int left = 0;
        int right = nums.length - 1;
        while (left < right) {
            if (nums[left] + nums[right] > target) {
                right--;
            } else {
                res += right - left;
                left++;
            }
            res = res % 1000000007;
        }
        return res % 1000000007;
    }
}
// @lc code=end

/*
// @lcpr case=start
// [2,5,3,5]\n6`>\n
// @lcpr case=end

// @lcpr case=start
// [2,2,1,9]\n10`>\n
// @lcpr case=end

 */
