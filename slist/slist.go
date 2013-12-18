// Package slist implements an intrusive singly linked list.
//
//
package slist

// Item contains the list metadata. Embed it as an anonymous field in
// your own structs.
type Item struct {
	next item
}

func (i *Item) get() *Item {
	return i
}

type item interface {
	get() *Item
}

// InsertAfter inserts newItem after marker.
func InsertAfter(marker item, newItem item) {
	m := marker.get()
	n := newItem.get()
	n.next = m.next
	m.next = newItem
}

// Next returns the next item after marker, or nil if none.
func Next(marker item) interface{} {
	return marker.get().next
}

// RemoveNext removes the next item after marker.
//
// Note: marker is not removed, the item after marker is. This is
// required because a singly linked list has no back pointers.
func RemoveNext(marker item) {
	i := marker.get()
	if i.next != nil {
		tmp := i.next.get()
		i.next = tmp.next
		tmp.next = nil
	}
}
