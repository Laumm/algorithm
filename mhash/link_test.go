package mhash

import (
	"testing"
)

func TestLink(t *testing.T) {
	l := new(link)
	l.set("1",1)
	_ ,ok := l.get("1")
	if ok != true {
		t.Fail()
		return
	}
	_,ok = l.get("2")
	if ok != false {
		t.Fail()
		return
	}
	l.del("1")
	if ok != false {
		t.Fail()
		return
	}
}
