package my_sort

import (
	"math"
	"sort"
	"testing"
)

func TestMQuick(t *testing.T) {
		to := generate(0,math.MaxInt32,100)
		t.Log("Generate:",to)
		get := make([]int,len(to))
		want := make([]int,len(to))
		copy(get,to)
		copy(want,to)
		MQuick(get)
		sort.Ints(want)
		t.Log("Get:",get)
		t.Log("Want:",want)
		if !compare(get,want) {
			t.Fatalf("Insert sorted  test err \n\t Get %v \n\t Want %v",get,want)
		}
		return
}

func TestMRandQuick(t *testing.T) {
	to := generate(0,math.MaxInt32,100)
	t.Log("Generate:",to)
	get := make([]int,len(to))
	want := make([]int,len(to))
	copy(get,to)
	copy(want,to)
	MRandQuick(get)
	sort.Ints(want)
	t.Log("Get:",get)
	t.Log("Want:",want)
	if !compare(get,want) {
		t.Fatalf("Insert sorted  test err \n\t Get %v \n\t Want %v",get,want)
	}
	return
}

// BenchmarkMQuick10-4   	   10000	    14 4856 ns/op	       0 B/op	       0 allocs/op
func BenchmarkMQuick10(b *testing.B) {
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
			MQuick(inputs[i % len(inputs)])
		}
	}
}

// BenchmarkMQuick100-4   	     300	   478 8788 ns/op	       0 B/op	       0 allocs/op
func BenchmarkMQuick100(b *testing.B) {
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
			MQuick(inputs[i % len(inputs)])
		}
	}
}

// BenchmarkMQuick1000-4   	      20	  7959 2886 ns/op	       0 B/op	       0 allocs/op
func BenchmarkMQuick1000(b *testing.B) {
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
			MQuick(inputs[i % len(inputs)])
		}
	}
}
// BenchmarkMQuick10000-4   	       1	14 1775 4331 ns/op	       0 B/op	       0 allocs/op
func BenchmarkMQuick10000(b *testing.B) {
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
			MQuick(inputs[i % len(inputs)])
		}
	}
}
// 无效
func BenchmarkMQuick100000(b *testing.B) {
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
			MQuick(inputs[i % len(inputs)])
		}
	}
}

