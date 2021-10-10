package list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}
// 1. 暴力，遍历， hash/set 存储遍历过的

//func hasCycle(head *ListNode) bool {
//	var nodeSeen = make(map[int]bool)
//	for head != nil {
//		if _, ok := nodeSeen[head.Val];ok{
//			return true
//		}
//		nodeSeen[head.Val] = true
//		head = head.Next
//	}
//	return false
//}
// 2. 快慢指针
func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil { //不是的条件
		return false
	}
	slow := head
	fast := head.Next
	for fast != slow {
		if head == nil || head.Next == nil {//不是的条件
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}