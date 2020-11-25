package main

/**
https://leetcode-cn.com/problems/add-two-numbers/
两数相加
执行用时：8 ms, 在所有 Go 提交中击败了92.26%的用户
内存消耗：4.9 MB, 在所有 Go 提交中击败了64.44%的用户
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		nodeOne     = l1
		nodeTwo     = l2
		sumNode     = &ListNode{}
		currentNode = sumNode
		carry       = 0
	)
	for {
		if nodeOne != nil {
			currentNode.Val += nodeOne.Val
			nodeOne = nodeOne.Next
		}
		if nodeTwo != nil {
			currentNode.Val += nodeTwo.Val
			nodeTwo = nodeTwo.Next
		}
		carry = currentNode.Val / 10
		currentNode.Val = currentNode.Val % 10

		if nodeOne == nil && nodeTwo == nil {
			if carry != 0 {
				currentNode.Next = &ListNode{Val: carry, Next: nil}
			}
			break
		} else {
			currentNode.Next = &ListNode{Val: carry, Next: nil}
		}
		currentNode = currentNode.Next
	}
	return sumNode
}

func main() {
	nodeOne := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val:  9,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	nodeTwo := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}

	//nodeThree := &ListNode{
	//	Val: 2,
	//	Next: &ListNode{
	//		Val: 4,
	//		Next: &ListNode{
	//			Val:  3,
	//			Next: nil,
	//		},
	//	},
	//}
	//nodeFour := &ListNode{
	//	Val: 5,
	//	Next: &ListNode{
	//		Val: 6,
	//		Next: &ListNode{
	//			Val:  4,
	//			Next: nil,
	//		},
	//	},
	//}
	addTwoNumbers(nodeOne, nodeTwo)
}
