package slist_test

import (
	"github.com/tv42/intrusive/slist"
	"testing"
)

type testItem struct {
	slist.Item
	ID int
}

func TestOne(t *testing.T) {
	var one = testItem{ID: 1}
	if g, e := slist.Next(&one), interface{}(nil); g != e {
		t.Errorf("Next(one): %#v != %#v", g, e)
	}
}

func TestSimple(t *testing.T) {
	var i1 = testItem{ID: 1}
	var i2 = testItem{ID: 2}
	slist.InsertAfter(&i1, &i2)
	if g, e := slist.Next(&i1), &i2; g != e {
		t.Errorf("Next(i1): %#v != %#v", g, e)
	}
	if g, e := slist.Next(&i2), interface{}(nil); g != e {
		t.Errorf("Next(i2): %#v != %#v", g, e)
	}
}

func TestRemoveNext(t *testing.T) {
	var i1 = testItem{ID: 1}
	var i2 = testItem{ID: 2}
	slist.InsertAfter(&i1, &i2)
	slist.RemoveNext(&i1)
	// i2 is now removed from the list
	if g, e := slist.Next(&i1), interface{}(nil); g != e {
		t.Errorf("Next(i1): %#v != %#v", g, e)
	}
}
