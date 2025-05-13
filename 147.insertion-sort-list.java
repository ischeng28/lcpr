/*
 * @lc app=leetcode.cn id=147 lang=java
 * @lcpr version=30204
 *
 * [147] 对链表进行插入排序
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
    public ListNode insertionSortList(ListNode head) {
        ListNode dummy = new ListNode(0, head);
        ListNode last = head;
        ListNode current = head.next;
        while (current != null) {
            if (current.val >= last.val) {
                last = current;
            } else {
                ListNode prev = dummy;
                while (prev.next.val <= current.val) {
                    prev = prev.next;
                }
                last.next = current.next;
                current.next = prev.next;
                prev.next = current;
            }
            current = last.next;
        }
        return dummy.next;
    }
}
// @lc code=end

/*
// @lcpr case=start
// [4,2,1,3]\n
// @lcpr case=end

// @lcpr case=start
// [-1,5,3,4,0]\n
// @lcpr case=end

 */
