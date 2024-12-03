package main

import (
	"fmt"
)

// структура элемента
type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node // голова списка
	tail *Node // последний узел списка
	size int   // размер списка
}

// вставка в начало списка
func (l *LinkedList) addNewHead(element int) {
	node := &Node{
		data: element,
		next: nil,
	}

	// если список был пуст
	if l.head == nil {
		l.tail = node
	} else {
		// прежний head сдвигаем на один узел назад
		// созданный узел в качестве next ссылается на head
		node.next = l.head
	}
	// устанавливаем новый узел в качестве головы списка
	l.head = node
	// увеличиваем размер списка
	l.size++
}

func (l *LinkedList) addNewTail(element int) {
	node := &Node{
		data: element,
		next: nil,
	}

	// если список пуст
	if l.tail == nil {
		l.head = node
	} else {
		l.tail.next = node
	}
	// записываем новый узел в качестве tail
	l.tail = node
	l.size++
}

func (l *LinkedList) insert(after int, element int) {
	// находим after
	search := l.head
	for search != nil {
		if search.data == after {
			// нашли нужный узел
			break
		}
		// если нет, то ищем дальше
		search = search.next
	}

	if search != nil {
		node := &Node{
			data: element,
			next: nil,
		}

		if search == l.tail {
			l.tail = node
		}

		node.next = search.next
		search.next = node
	}
	l.size++
}

func (l *LinkedList) search(element int) *Node {
	current := l.head
	for current != nil {
		if current.data == element{
			return current
		}
		current = current.next
	}
	return nil
}

func (l *LinkedList) printer() {
	var element *Node = l.head
	for element != nil {
		fmt.Print(element.data, " ")
		element = element.next
	}
	fmt.Print("\n")
}

func main() {
	lst := LinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
	lst.addNewHead(6)
	lst.addNewTail(10)
	lst.addNewHead(5)
	lst.insert(6, 8)
	lst.printer()
	fmt.Println(lst.search(6))
}
