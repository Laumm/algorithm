package my_sort

import (
	"math"
	"sort"
	"testing"
)


func TestMInsert(t *testing.T) {
	to := generate(0,math.MaxInt8,10)
	get := make([]int,len(to))
	want := make([]int,len(to))
	copy(get,to)
	copy(want,to)
	MInsert(get)
	sort.Ints(want)
	if !compare(get,want) {
		t.Fatalf("Insert sorted  test err \n\t Get %v \n\t Want %v",get,want)
	}
	return
}

// 测试 五位数 20000	     88929 ns/op	       0 B/op	       0 allocs/op
func BenchmarkMInsert10(b *testing.B) {
	b.ReportAllocs()
	for i := 0 ; i < b.N;i++ {
		for j := 0 ; j < len(to10) ; j ++ {
			MInsert(to10[j])
		}
	}
}

// 测试六位数   2000	    696255 ns/op	       0 B/op	       0 allocs/op
func BenchmarkMInsert100(b *testing.B) {
	b.ReportAllocs()
	for i := 0 ; i < b.N;i++ {
		for j := 0 ; j < len(to100) ; j ++ {
			MInsert(to100[j])
		}
	}
}

//测试七位数 200	   8060842 ns/op	       0 B/op	       0 allocs/op
func BenchmarkMInsert1000(b *testing.B) {
	b.ReportAllocs()
	for i := 0 ; i < b.N;i++ {
		for j := 0 ; j < len(to1000) ; j ++ {
			MInsert(to1000[j])
		}
	}
}
//  1	24290005834 ns/op	       0 B/op	       0 allocs/op
func BenchmarkMInsert10000(b *testing.B) {
	b.ReportAllocs()
	for i := 0 ; i < b.N;i++ {
		for j := 0 ; j < len(to10000) ; j ++ {
			MInsert(to10000[j])
		}
	}
}

func BenchmarkMInsert100000(b *testing.B) {
	b.ReportAllocs()
	for i := 0 ; i < b.N;i++ {
		for j := 0 ; j < len(to100000) ; j ++ {
			MInsert(to100000[j])
		}
	}
}