package main

func main() {
	// ----- ONLY HERE FOR TESTING ----- \\
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// ========== THE SOLUTION FUNCTION ========== \\
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// ===== DUMMY NOTE TO START ===== \\
	dummy := &ListNode{}
	completeNode := dummy

	curr1, curr2 := list1, list2

	// ===== WHILE THE CURRENT NODES AREN'T NULL ===== \\
	for curr1 != nil && curr2 != nil {

		// IF NODE1'S VALUE IS BIGGER THAN NODE2'S \\
		if curr1.Val <= curr2.Val {
			completeNode.Next = curr1 // SET MERGED LIST'S NEXT TO CURR1
			curr1 = curr1.Next        // SET THE NEXT NODE1 AS THE CURRENT NODE1
		} else {
			completeNode.Next = curr2 // SET MERGED LIST'S NEXT TO CURR2
			curr2 = curr2.Next        // SET THE NEXT NODE2 AS THE CURRENT NODE2
		}
		completeNode = completeNode.Next // MOVE FORWARD
	}

	// APPEND AND ATTACH THE REMAING NODE(S) \\
	if curr1 != nil {
		completeNode.Next = curr1
	} else {
		completeNode.Next = curr2
	}

	return dummy.Next // RETURN THE HEAD OF THE MERGED LIST
}
