package slist_test

import (
	"fmt"
	"github.com/tv42/intrusive/slist"
)

func Example() {

	type MyItem struct {
		slist.Item
		ID int
	}

	// helper function to check for nil before type asserting; slist.Next
	// returns an interface{}, so this intermediate step is needed
	next := func(i *MyItem) *MyItem {
		n := slist.Next(i)
		if n == nil {
			return nil
		}
		return n.(*MyItem)
	}

	var i1 = MyItem{ID: 1}
	var i2 = MyItem{ID: 2}
	slist.InsertAfter(&i1, &i2)
	for cursor := &i1; cursor != nil; cursor = next(cursor) {
		fmt.Println(cursor.ID)
	}
	// Output:
	// 1
	// 2
}
