package main


func main() {
	lst := &DoublyLinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
	lst.addNewHead(4)
	lst.addNewHead(3)
	lst.addNewHead(1)
	lst.addNewHead(0)
	lst.addNewHead(-1)
	lst.addNewTail(10)
	lst.addNewTail(13)
	lst.headPrinter()
	lst.tailPrinter()

	lst.insert(4, 5)
	lst.insert(10, 11)
	lst.headPrinter()
	lst.tailPrinter()
}