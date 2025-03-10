package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	a, b := 123, 234
	fmt.Println("Before reversing: a =", a, "b =", b)

	for x, y := a, b; x > 0 || y > 0; {
		if x > 0 {
			a = a*10 + x%10
			x /= 10
		}
		if y > 0 {
			b = b*10 + y%10
			y /= 10
		}
	}

	fmt.Println("After reversing: a =", a, "b =", b)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	completeNode := dummy
	curr1, curr2 := l1, l2
	num1 := ""
	num2 := ""

	// Convert linked lists to reversed strings
	for curr1 != nil || curr2 != nil {
		if curr1 != nil {
			num1 += strconv.Itoa(curr1.Val)
			curr1 = curr1.Next
		}
		if curr2 != nil {
			num2 += strconv.Itoa(curr2.Val)
			curr2 = curr2.Next
		}
	}

	// Reverse the numbers (to get the actual numerical value)
	reversed1 := reverseString(num1)
	reversed2 := reverseString(num2)

	// Convert reversed strings to big.Int
	bigNum1 := new(big.Int)
	bigNum2 := new(big.Int)
	bigNum1.SetString(reversed1, 10)
	bigNum2.SetString(reversed2, 10)

	// Compute sum using big.Int
	sum := new(big.Int).Add(bigNum1, bigNum2)

	// Convert sum to string
	sumStr := sum.String()

	// Convert sum string back to a linked list in reverse order
	for i := len(sumStr) - 1; i >= 0; i-- {
		if completeNode.Next == nil {
			completeNode.Next = &ListNode{}
		}
		completeNode.Next.Val, _ = strconv.Atoi(string(sumStr[i]))
		completeNode = completeNode.Next
	}

	return dummy.Next

}

func reverseString(s string) string {
	runes := []rune(s) // Convert string to slice of runes
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] // Swap characters
	}
	return string(runes) // Convert slice of runes back to string
}
