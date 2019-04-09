package my_sort

import (
	"testing"
)

func TestMCount(t *testing.T) {
	generate := []uint8{4,3,2,5,100,10,200,180}
	get := MCount(generate)
	want := []uint8{2,3,4,5,10,100,180,200}
	if len(want) != len(get){
		t.Fatalf("Test MCount err \t Get Len = %d ,Wan len = %d",len(get),len(want))
	}
	for i:=0 ;i <len(get); i ++  {
		if get[i] != want[i] {
			t.Fatalf("Test Mcount err \t Get %v ,wan %v",get,want)
		}
	}

}

