package main

import (
	"fmt"
	"slices"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//fmt.Println("hello world")
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	addTwoNumbers(l1, l2)

	//fmt.Println(calcLL)

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Total := convertToSlice(l1)
	l2Total := convertToSlice(l2)

	result := processResult(l1Total[0], l2Total[0])

	return result
}

func printLinkedList(ll *ListNode) {
	if ll != nil {
		fmt.Printf("%d ", ll.Val)
		printLinkedList(ll.Next)
	}
}

func convertToSlice(ll *ListNode) []int {
	var LLR []int

	if ll != nil {
		LLR = append(LLR, ll.Val)

		if ll.Next != nil {
			LLR = append(LLR, convertToSlice(ll.Next)...)
		} else {
			return LLR
		}
	}

	slices.Reverse(LLR)

	var totalString string
	for _, e := range LLR {
		totalString += strconv.Itoa(e)
	}

	total, err := strconv.Atoi(totalString)
	if err != nil {
		fmt.Println(err)
	}

	return []int{total}
}

func processResult(l1Total int, l2Total int) *ListNode {
	totalCalculated := l1Total + l2Total
	fmt.Println("calculated: ", totalCalculated)
	totalString := strconv.Itoa(totalCalculated)

	result := createLinkedList(totalString)

	return result

}

func createLinkedList(value string) *ListNode {
	var linkedList *ListNode

	value = reverseString(value)

	for i, l := range value {
		newValue := string(l)
		newValueInt, _ := strconv.Atoi(newValue)

		if i == 0 {
			linkedList = &ListNode{
				Val: newValueInt,
			}
		} else {
			addToLinkedList(newValueInt, linkedList)
		}
	}

	return linkedList
}

func addToLinkedList(value int, l *ListNode) {
	if l != nil {
		if l.Next != nil {
			addToLinkedList(value, l.Next)
		} else {
			l.Next = &ListNode{
				Val: value,
			}
		}
	}
}

func reverseString(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}

//func reverseList(ll *ListNode, reversed *ListNode) {
//	var llR []*ListNode
//
//	//if reversed == nil {
//	//	reversed = &ListNode{}
//	//}
//
//	if ll.Next != nil {
//		llR = append(llR, ll)
//		fmt.Println(ll.Val)
//		for i, l := range llR {
//			fmt.Println("iteration ", i)
//			fmt.Println(l.Val)
//		}
//		reverseList(ll.Next, nil)
//	} else {
//		fmt.Println(ll.Val)
//		llR = append(llR, ll)
//		for i, l := range llR {
//			fmt.Println("iteration ", i)
//			fmt.Println(l.Val)
//		}
//	}
//
//	fmt.Println("Length of llR ", len(llR))
//	//slices.Reverse(llR)
//
//	//fmt.Println(llR)
//	for i, l := range llR {
//		fmt.Println(l.Val)
//		if reversed == nil {
//			reversed = &ListNode{
//				Val:  l.Val,
//				Next: &ListNode{},
//			}
//		} else {
//			if l != nil {
//
//				next := i + 1
//				reversed.Val = l.Val
//
//				if next < len(llR) {
//					reversed.Next = llR[i+1]
//				}
//			}
//		}
//	}
//
//	printLinkedList(reversed)
//	//fmt.Println(reversed)
//	//return reversed
//}
//
//func NextNode(l1 *ListNode, l2 *ListNode, ll *ListNode) {
//
//	if l1 != nil && l2 != nil {
//		llNodeVal := l1.Val + l2.Val
//		if ll == nil {
//			ll = &ListNode{
//				Val:  llNodeVal,
//				Next: &ListNode{},
//			}
//		} else {
//			ll.Val = llNodeVal
//			ll.Next = &ListNode{}
//		}
//
//		//fmt.Println(llNodeVal)
//		fmt.Println(l1.Val)
//		if l1.Next != nil {
//			//fmt.Println(llNodeVal)
//			NextNode(l1.Next, l2.Next, ll.Next)
//		} else {
//			return
//		}
//	}
//}
