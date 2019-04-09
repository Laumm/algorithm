package my_sort

import (
	"math"
	"sort"
	"testing"
)

func TestMMerger(t *testing.T) {
	to := generate(0,math.MaxInt32,100)
	get := make([]int,len(to))
	want := make([]int,len(to))
	copy(get,to)
	copy(want,to)
	MMerger(get)
	sort.Ints(want)
	if !compare(get,want) {
		t.Fatalf("Insert sorted  test err \n\t Get %v \n\t Want %v",get,want)
	}
	return
}

// BenchmarkMMerger10-4   	    5000	    31 1329 ns/op	   80000 B/op	    1000 allocs/op (1000  次排序)
func BenchmarkMMerger10(b *testing.B) {
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
			MMerger(inputs[i % len(inputs)])
		}
	}
}

// BenchmarkMMerger100-4   	     500	   362 5604 ns/op	  896000 B/op	    1000 allocs/op
func BenchmarkMMerger100(b *testing.B) {
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
			MMerger(inputs[i % len(inputs)])
		}
	}
}

// BenchmarkMMerger1000-4   	      30	  4284 2496 ns/op	 8192000 B/op	    1000 allocs/op
func BenchmarkMMerger1000(b *testing.B) {
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
			MMerger(inputs[i % len(inputs)])
		}
	}
}

// BenchmarkMMerger10000-4   	       3	 4 4099 2282 ns/op	81920000 B/op	    1000 allocs/op
func BenchmarkMMerger10000(b *testing.B) {
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
			MMerger(inputs[i % len(inputs)])
		}
	}
}

// BenchmarkMMerger100000-4   	       1	60 9849 4002 ns/op	802816000 B/op	    1000 allocs/op
func BenchmarkMMerger100000(b *testing.B) {
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
			MMerger(inputs[i % len(inputs)])
		}
	}
}
