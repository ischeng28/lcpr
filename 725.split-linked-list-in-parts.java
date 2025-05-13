/*
 * @lc app=leetcode.cn id=725 lang=java
 * @lcpr version=30204
 *
 * [725] 分隔链表
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
    public ListNode[] splitListToParts(ListNode head, int k) {
        int len = 0;
        ListNode tmp = head;
        while (tmp != null) {
            len++;
            tmp = tmp.next;
        }
        int num = len / k, remain = len % k;
        ListNode[] res = new ListNode[k];
        ListNode current = head;

        for (int i = 0; i < k && current != null; i++) {
            res[i] = current;
            int partSize = num;
            if (remain > 0) {
                partSize++;
                remain--;
            }
            for (int j = 1; j < partSize; j++) {
                current = current.next;
            }
            ListNode nextNode = current.next;
            current.next = null;
            current = nextNode;
        }
        return res;

    }
}
// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n5\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5,6,7,8,9,10]\n3\n
// @lcpr case=end

 */
