/*
 * @lc app=leetcode.cn id=3264 lang=java
 * @lcpr version=30204
 *
 * [3264] K 次乘运算后的最终数组 I
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

import java.util.PriorityQueue;

class Solution {
    public int[] getFinalState(int[] nums, int k, int multiplier) {
        PriorityQueue<Integer> pq = new PriorityQueue<>((a, b) -> {
            if (nums[a] == nums[b]) {
                return a - b;
            }
            return nums[a] - nums[b];
        });
        for (int i = 0; i < nums.length; i++) {
            pq.offer(i);
        }
        while (k > 0) {
            int idx = pq.poll();
            nums[idx] *= multiplier;
            pq.offer(idx);
            k--;
        }
        return nums;
    }
}
// @lc code=end

/*
// @lcpr case=start
// [2,1,3,5,6]\n5\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n3\n4\n
// @lcpr case=end

 */
