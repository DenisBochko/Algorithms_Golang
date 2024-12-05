package main

import (
	"fmt"
)

// структура элемента
type Node struct {
	data int
	next *Node
	prev *Node
}

type DoublyLinkedList struct {
	head *Node // голова списка
	tail *Node // конец списка
	size int   // размер списка
}

func (l *DoublyLinkedList) addNewHead(element int) {
	node := &Node{
		data: element,
		next: nil,
		prev: nil,
	}
	if l.head == nil {
		l.head = node
		l.tail = node
		l.size++
		return
	}
	node.next = l.head
	l.head.prev = node
	l.head = node
	l.size++
}

func (l *DoublyLinkedList) addNewTail(element int) {
	node := &Node{
		data: element,
		next: nil,
		prev: nil,
	}

	if l.tail == nil {
		l.tail = node
		l.head = node
		l.size++
		return
	}
	l.tail.next = node
	node.prev = l.tail
	l.tail = node
	l.size++
}

func (l *DoublyLinkedList) insert(after int, element int) {
	node := &Node{
		data: element,
		next: nil,
		prev: nil,
	}

	if l.head == nil {
		fmt.Println("Список пуст! Используйте вставку в начало или конец!")
		return
	}

	current := l.head
	for current != nil {
		if current.data == after{
			break
		}
		current = current.next
	}

	if current != nil{
		node.next = current.next
		node.prev = current
		
		current.next.prev = node
		current.next = node
		l.size++
	}
}

func (l *DoublyLinkedList) headPrinter() {
	element := l.head
	for element != nil {
		fmt.Printf("%v ", element.data)
		element = element.next
	}
	fmt.Printf("\nsize = %v\n", l.size)
}

func (l *DoublyLinkedList) tailPrinter() {
	element := l.tail
	for element != nil {
		fmt.Printf("%v ", element.data)
		element = element.prev
	}
	fmt.Printf("\nsize = %v\n", l.size)
}
