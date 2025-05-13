/*
 * @lc app=leetcode.cn id=92 lang=java
 * @lcpr version=30204
 *
 * [92] 反转链表 II
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
    public ListNode reverseBetween(ListNode head, int left, int right) {
        ListNode dummy = new ListNode(0, head);
        ListNode p0 = dummy;
        for (int i = 0; i < left - 1; i++) {
            p0 = p0.next;
        }
        ListNode current = p0.next;
        ListNode pre = null;
        for (int j = 0; j < right - left + 1; j++) {
            ListNode next = current.next;
            current.next = pre;
            pre = current;
            current = next;
        }
        p0.next.next = current;
        p0.next = pre;

        return dummy.next;
    }
}
// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n4\n
// @lcpr case=end

// @lcpr case=start
// [5]\n1\n1\n
// @lcpr case=end

 */
