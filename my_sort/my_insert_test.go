package my_sort

import (
	"math"
	"sort"
	"testing"
)


func TestMInsert(t *testing.T) {
	to := generate(0,math.MaxInt32,10)
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

// BenchmarkMInsert10-4   	   20000	     8 3373 ns/op	       0 B/op	       0 allocs/op
func BenchmarkMInsert10(b *testing.B) {
	b.ReportAllocs()
	generates := generates(0,math.MaxInt32,10,1000)
	b.ResetTimer()
	for i := 0 ; i < b.N;i++ {
		b.StopTimer()
		inputs := make([][]int,len(generates))
		for j := 0 ;j < len(generates);j++{
			input := make([]int,len(generates[j]))
			copy(input,generates[j])
			inputs[j] = input
		}
		b.StartTimer()
		for j := 0 ; j < len(inputs); j ++ {
			MInsert(inputs[i % len(inputs)])
		}
	}
}


// BenchmarkMInsert100-4   	    2000	    66 2287 ns/op	       0 B/op	       0 allocs/op
func BenchmarkMInsert100(b *testing.B) {
	b.ReportAllocs()
	generates := generates(0,math.MaxInt32,100,1000)
	b.ResetTimer()
	for i := 0 ; i < b.N;i++ {
		b.StopTimer()
		inputs := make([][]int,len(generates))
		for j := 0 ;j < len(generates);j++{
			input := make([]int,len(generates[j]))
			copy(input,generates[j])
			inputs[j] = input
		}
		b.StartTimer()
		for j := 0 ; j < len(inputs); j ++ {
			MInsert(inputs[i % len(inputs)])
		}
	}
}

// BenchmarkMInsert1000-4   	     200	   639 3116 ns/op	       0 B/op	       0 allocs/op
func BenchmarkMInsert1000(b *testing.B) {
	b.ReportAllocs()
	generates := generates(0,math.MaxInt32,1000,1000)
	b.ResetTimer()
	for i := 0 ; i < b.N;i++ {
		b.StopTimer()
		inputs := make([][]int,len(generates))
		for j := 0 ;j < len(generates);j++{
			input := make([]int,len(generates[j]))
			copy(input,generates[j])
			inputs[j] = input
		}
		b.StartTimer()
		for j := 0 ; j < len(inputs); j ++ {
			MInsert(inputs[i % len(inputs)])
		}
	}
}
//  BenchmarkMInsert10000-4   	      20	  8388 2469 ns/op	       0 B/op	       0 allocs/op
func BenchmarkMInsert10000(b *testing.B) {
	b.ReportAllocs()
	generates := generates(0,math.MaxInt32,10000,1000)
	b.ResetTimer()
	for i := 0 ; i < b.N;i++ {
		b.StopTimer()
		inputs := make([][]int,len(generates))
		for j := 0 ;j < len(generates);j++{
			input := make([]int,len(generates[j]))
			copy(input,generates[j])
			inputs[j] = input
		}
		b.StartTimer()
		for j := 0 ; j < len(inputs); j ++ {
			MInsert(inputs[i % len(inputs)])
		}
	}
}

// BenchmarkMInsert100000-4   	       1	39 3747 6169 ns/op	       0 B/op	       0 allocs/op
func BenchmarkMInsert100000(b *testing.B) {
	b.ReportAllocs()
	generates := generates(0,math.MaxInt32,100000,1000)
	b.ResetTimer()
	for i := 0 ; i < b.N;i++ {
		b.StopTimer()
		inputs := make([][]int,len(generates))
		for j := 0 ;j < len(generates);j++{
			input := make([]int,len(generates[j]))
			copy(input,generates[j])
			inputs[j] = input
		}
		b.StartTimer()
		for j := 0 ; j < len(inputs); j ++ {
			MInsert(inputs[i % len(inputs)])
		}
	}
}