package my_sort

import (
	"math"
	"math/rand"
	"time"
)

var to10 [1000][]int
var to100[1000][]int
var to1000[1000][]int
var to10000[1000][]int
var to100000[1000][]int
func init()  {

	for i := 0 ; i < 1000 ; i ++ {
		to10[i] = generate(0,math.MaxInt8,10)
		to100[i] = generate(0,math.MaxInt8,100)
		to1000[i] = generate(0,math.MaxInt8,1000)
		to10000[i] = generate(0,math.MaxInt8,10000)
		to100000[i] = generate(0,math.MaxInt8,100000)
	}


}

// 随机生成正数切片
// 0 <= min <= 元素值 <= max 个数 == len
// len < 0 或者 min > max 返回  nil
func generate (min ,max , len int) [] int{
	if min > max || len  <= 0 {
		return  nil
	}
	results := make([]int,0,len)
	rand.Seed(time.Now().UnixNano())
	for i := 0 ; i <len ; i ++ {
		r := rand.Intn(max - min + 1  )
		results =append(results,r + min)
	}
	return  results
}

func compare(a,b [] int) bool {
	if len(a) != len(b) {
		return  false
	}
	for i := 0 ; i <len(a); i ++ {
		if a[i] != b[i] {
			return  false
		}
	}
	return  true
}
