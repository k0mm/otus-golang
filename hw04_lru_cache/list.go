package hw04lrucache

type List interface {
	Len() int
	Front() *listItem
	Back() *listItem
	PushFront(v interface{}) *listItem
	PushBack(v interface{}) *listItem
	Remove(i *listItem)
	MoveToFront(i *listItem)
}

type listItem struct {
	value interface{}
	next  *listItem
	prev  *listItem
}

type list struct {
	len   int
	front *listItem
	back  *listItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *listItem {
	return l.front
}

func (l *list) Back() *listItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *listItem {
	item := listItem{
		value: v,
		prev:  nil,
	}

	if l.front != nil {
		l.front.prev = &item
		item.next = l.front
	} else {
		l.back = &item
	}

	l.len++
	l.front = &item
	return &item
}

func (l *list) PushBack(v interface{}) *listItem {
	item := listItem{
		value: v,
		next:  nil,
	}

	if l.back != nil {
		l.back.next = &item
		item.prev = l.back
	} else {
		l.front = &item
	}

	l.len++
	l.back = &item
	return &item
}

func (l *list) Remove(i *listItem) {
	if i.prev != nil {
		i.prev.next = i.next
	} else {
		l.front = i.next
	}

	if i.next != nil {
		i.next.prev = i.prev
	} else {
		l.back = i.prev
	}

	l.len--
}

func (l *list) MoveToFront(i *listItem) {
	if i == l.front {
		return
	}

	if i.next != nil {
		i.next.prev = i.prev
	} else {
		l.back = i.prev
	}

	if i.prev != nil {
		i.prev.next = i.next
		i.next = l.front
		l.front.prev = i
	}

	i.prev = nil
	l.front = i
}

func NewList() List {
	return &list{}
}
