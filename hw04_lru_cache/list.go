package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	head *ListItem
	tail *ListItem
	len  int
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) insertAfter(node, newNode *ListItem) {
	newNode.Prev = node

	if node.Next == nil {
		newNode.Next = nil
		l.tail = newNode
	} else {
		newNode.Next = node.Next
		node.Next.Prev = newNode
	}
	node.Next = newNode
	l.len++
}

func (l *list) insertBefore(node, newNode *ListItem) {
	newNode.Next = node

	if node.Prev == nil {
		newNode.Prev = nil
		l.head = newNode
	} else {
		newNode.Prev = node.Prev
		node.Prev.Next = newNode
	}
	node.Prev = newNode
	l.len++
}

func (l *list) PushFront(v interface{}) *ListItem {
	n := new(ListItem)
	n.Value = v

	if l.head == nil {
		l.head = n
		l.tail = n
		n.Prev = nil
		n.Next = nil
		l.len++
	} else {
		l.insertBefore(l.head, n)
	}
	return n
}

func (l *list) PushBack(v interface{}) *ListItem {
	n := new(ListItem)
	n.Value = v

	if l.tail == nil {
		l.PushFront(n)
	} else {
		l.insertAfter(l.tail, n)
	}

	return n
}

func (l *list) Remove(i *ListItem) {
	if i.Prev == nil {
		l.head = i.Next
		i.Next.Prev = nil
	} else {
		i.Prev.Next = i.Next
	}

	if i.Next == nil {
		l.tail = i.Prev
		i.Prev.Next = nil
	} else {
		i.Next.Prev = i.Prev
	}

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.insertBefore(l.head, i)
}
