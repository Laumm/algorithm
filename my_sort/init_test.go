package my_sort

import (
	"math/rand"
	"time"
)


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

func generates (min ,max , len,num int) [][]int  {
	result := make([][]int,0,num)
	for i := 0 ;i < num ; i++ {
		result =append(result,generate(min,max,len))
	}
	return result
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
