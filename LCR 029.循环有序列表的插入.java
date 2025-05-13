/*
 * @lc app=leetcode.cn id=LCR 029 lang=java
 * @lcpr version=30204
 *
 * [LCR 029] 循环有序列表的插入
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/*
// Definition for a Node.
class Node {
    public int val;
    public Node next;

    public Node() {}

    public Node(int _val) {
        val = _val;
    }

    public Node(int _val, Node _next) {
        val = _val;
        next = _next;
    }
};
*/

class Solution {
    public Node insert(Node head, int insertVal) {
        Node node = new Node(insertVal);
        if (head == null) {
            node.next = node;
            return node;
        }
        if (head.next == head) {
            head.next = node;
            node.next = head;
            return head;
        }
        Node current = head, next = head.next;
        while (next != head) {
            if (insertVal >= current.val && insertVal <= next.val) {
                break;
            }
            if (current.val > next.val) {
                if (insertVal > current.val || insertVal < next.val) {
                    break;
                }
            }
            current = current.next;
            next = next.next;
        }
        current.next = node;
        node.next = next;
        return head;
    }
}
// @lc code=end

/*
// @lcpr case=start
// [3,4,1]\n2\n
// @lcpr case=end

// @lcpr case=start
// []\n1\n
// @lcpr case=end

// @lcpr case=start
// [1]\n0\n
// @lcpr case=end

 */
