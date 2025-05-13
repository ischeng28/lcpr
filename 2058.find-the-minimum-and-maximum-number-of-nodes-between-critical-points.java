/*
 * @lc app=leetcode.cn id=2058 lang=java
 * @lcpr version=30204
 *
 * [2058] 找出临界点之间的最小和最大距离
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public int[] nodesBetweenCriticalPoints(ListNode head) {
        ListNode a = head, b = head.next, c = head.next.next;
        int first = 0, last = 0, minDistance = Integer.MAX_VALUE;
        int pre = 0;
        for (int i = 1; c != null; i++) {
            if (b.val > a.val && b.val > c.val || b.val < a.val && b.val < c.val) {
                if (first == 0) {
                    first = i;
                }
                last = i;
                if (first != last) {
                    minDistance = Math.min(minDistance, last - pre);
                }
                pre = i;
            }
            a = b;
            b = c;
            c = c.next;
        }
        if (first == last) {
            return new int[] { -1, -1 };
        }
        return new int[] { minDistance, last - first };
    }
}
// @lc code=end

/*
// @lcpr case=start
// [3,1]\n
// @lcpr case=end

// @lcpr case=start
// [5,3,1,2,5,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,3,2,2,3,2,2,2,7]\n
// @lcpr case=end

// @lcpr case=start
// [2,3,3,2]\n
// @lcpr case=end

 */
