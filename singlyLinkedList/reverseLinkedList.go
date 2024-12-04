package main

import "fmt"

func main() {
	lst := LinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
	// Создать список с цыклом: (выведет true)
	// lst.addNewHead(5)
	// lst.addNewHead(4)
	// lst.addNewHead(3)
	// fmt.Println("указатель на 4 = ", lst.head)
	// lst.tail.next = lst.head
	// lst.addNewHead(2)
	// lst.addNewHead(1)
	// // lst.printer()
	// fmt.Println(hasCycle(lst.head))

	lst.addNewHead(5)
	lst.addNewHead(4)
	lst.addNewHead(3)
	lst.addNewHead(2)
	lst.addNewHead(1)
	lst.addNewHead(0)
	lst.printer()
	fmt.Println(middleOfList(lst.head))
}

func (l *LinkedList) reverseLinkedList() {
	var prev *Node = nil
	current := l.head

	for current != nil {
		next := current.next
		current.next = prev

		prev = current
		current = next
	}

	l.head = prev
}

func hasCycle(head *Node) bool {
	if head == nil || head.next == nil {
		return false
	}

	slow := head
	fast := head.next

	for slow != fast {
		if fast == nil || slow == nil {
			return false
		}
		slow = slow.next
		fast = fast.next.next
	}
	return true
}

func middleOfList(head *Node) *Node {
	slow := head
	fast := head
	
	// Очень важно !!!
	for fast != nil && fast.next != nil{
		slow = slow.next
		fast = fast.next.next
	}

	return slow
}
