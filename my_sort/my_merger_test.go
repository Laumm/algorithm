package my_sort

import (
	"math"
	"sort"
	"testing"
)

func TestMMerger(t *testing.T) {
	to := generate(0,math.MaxInt8,100)
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

// BenchmarkMMerger10-4   	    5000	    371838 ns/op	   80000 B/op	    1000 allocs/op
func BenchmarkMMerger10(b *testing.B) {
	b.ReportAllocs()
	for i := 0 ; i < b.N;i++ {
		for j := 0 ; j < len(to10) ; j ++ {
			MMerger(to10[j])
		}
	}
}

// BenchmarkMMerger100-4   	     300	   4040263 ns/op	  896000 B/op	    1000 allocs/op
func BenchmarkMMerger100(b *testing.B) {
	b.ReportAllocs()
	for i := 0 ; i < b.N;i++ {
		for j := 0 ; j < len(to100) ; j ++ {
			MMerger(to100[j])
		}
	}
}

// BenchmarkMMerger1000-4   	      30	  42842496 ns/op	 8192000 B/op	    1000 allocs/op
func BenchmarkMMerger1000(b *testing.B) {
	b.ReportAllocs()
	for i := 0 ; i < b.N;i++ {
		for j := 0 ; j < len(to1000) ; j ++ {
			MMerger(to1000[j])
		}
	}
}

// BenchmarkMMerger10000-4   	       3	 475188885 ns/op	81920000 B/op	    1000 allocs/op
func BenchmarkMMerger10000(b *testing.B) {
	b.ReportAllocs()
	for i := 0 ; i < b.N;i++ {
		for j := 0 ; j < len(to10000) ; j ++ {
			MMerger(to10000[j])
		}
	}
}

// BenchmarkMMerger100000-4   	       1	10883710854 ns/op	802816000 B/op	    1000 allocs/op
func BenchmarkMMerger100000(b *testing.B) {
	b.ReportAllocs()
	for i := 0 ; i < b.N;i++ {
		for j := 0 ; j < len(to100000) ; j ++ {
			MMerger(to100000[j])
		}
	}
}

//BenchmarkMMergerTmp10-4   	    5000	    277234 ns/op	       0 B/op	       0 allocs/op
func BenchmarkMMergerTmp10(b *testing.B) {
	b.ReportAllocs()
	tmp := make([]int,10)
	for i := 0 ; i < b.N;i++ {
		for j := 0 ; j < len(to10) ; j ++ {
			MMergerTmp(to10[j],tmp)
		}
	}
}

// BenchmarkMMergerTmp100-4   	     500	   3967186 ns/op	       0 B/op	       0 allocs/op
func BenchmarkMMergerTmp100(b *testing.B) {
	b.ReportAllocs()
	tmp := make([]int,100)
	for i := 0 ; i < b.N;i++ {
		for j := 0 ; j < len(to100) ; j ++ {
			MMergerTmp(to100[j],tmp)
		}
	}
}

//BenchmarkMMergerTmp1000-4   	      30	  44497808 ns/op	       0 B/op	       0 allocs/op
func BenchmarkMMergerTmp1000(b *testing.B) {
	b.ReportAllocs()
	tmp := make([]int,1000)
	for i := 0 ; i < b.N;i++ {
		for j := 0 ; j < len(to1000) ; j ++ {
			MMergerTmp(to1000[j],tmp)
		}
	}
}

//BenchmarkMMergerTmp10000-4   	       1	1152212781 ns/op	   81920 B/op	       1 allocs/op
func BenchmarkMMergerTmp10000(b *testing.B) {
	b.ReportAllocs()
	tmp := make([]int,10000)
	for i := 0 ; i < b.N;i++ {
		for j := 0 ; j < len(to10000) ; j ++ {
			MMergerTmp(to10000[j],tmp)
		}
	}
}

//BenchmarkMMergerTmp100000-4   	       1	11582650411 ns/op	  802816 B/op	       1 allocs/op
func BenchmarkMMergerTmp100000(b *testing.B) {
	b.ReportAllocs()
	tmp := make([]int,100000)
	for i := 0 ; i < b.N;i++ {
		for j := 0 ; j < len(to100000) ; j ++ {
			MMergerTmp(to100000[j],tmp)
		}
	}
}