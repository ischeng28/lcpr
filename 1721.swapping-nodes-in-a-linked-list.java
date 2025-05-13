/*
 * @lc app=leetcode.cn id=1721 lang=java
 * @lcpr version=30204
 *
 * [1721] 交换链表中的节点
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
    public ListNode swapNodes(ListNode head, int k) {
        ListNode dummy = new ListNode(0, head);
        ListNode fast = dummy;
        ListNode slow = dummy;
        for (int i = 0; i < k; i++) {
            fast = fast.next;
        }
        ListNode p1 = fast;
        while (fast != null) {
            fast = fast.next;
            slow = slow.next;
        }
        int tmp = p1.val;
        p1.val = slow.val;
        slow.val = tmp;
        return dummy.next;
    }
}
// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [7,9,6,6,7,8,3,0,9,5]\n5\n
// @lcpr case=end

// @lcpr case=start
// [1]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n2\n
// @lcpr case=end

 */
