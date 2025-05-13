/*
 * @lc app=leetcode.cn id=2807 lang=java
 * @lcpr version=30204
 *
 * [2807] 在链表中插入最大公约数
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
    public ListNode insertGreatestCommonDivisors(ListNode head) {
        ListNode newHead = head;
        while (head != null && head.next != null) {
            int gcd = gcd(head.val, head.next.val);
            ListNode newNode = new ListNode(gcd);
            newNode.next = head.next;
            head.next = newNode;
            head = head.next.next;
        }
        return newHead;
    }

    public int gcd(int a, int b) {
        if (b == 0) {
            return a;
        }
        return gcd(b, a % b);
    }
}
// @lc code=end

/*
// @lcpr case=start
// [18,6,10,3]\n
// @lcpr case=end

// @lcpr case=start
// [7]\n
// @lcpr case=end

 */
