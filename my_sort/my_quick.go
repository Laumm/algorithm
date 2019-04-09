package my_sort

import (
	"math/rand"
	"time"
)

// 快速排序
// 时间复杂度（期望） Θ(nlgn) 空间复杂度 Θ(1)
// 正序或者倒叙时性能最低 Θ(n^2)
// 不稳定

func quick(inputs [] int,p ,q int )  {
	if l := q -p ; l <= 0  { 	// 递归终止判定
		return
	}
	m := inputs[p] 	// 主元素为首元素
	i , j := p,p
	for j < q {
		if inputs[j+1] <=  m {
			inputs[i+1],inputs[j+1] = inputs[j+1],inputs[i+1]
			i ++
		}
		j ++
	}
	inputs[p] ,inputs[i] = inputs[i],inputs[p]  	// 首元素交换
	quick(inputs,p,i-1)
	quick(inputs,i+1 ,q)
	return
}

// 算法实现有问题，运行很慢
func MQuick(inputs [] int)  {
	quick(inputs , 0,len(inputs)-1)
}

func randQuick(inputs [] int,p ,q int)  {
	if l := q -p ; l <= 0  { 	// 递归终止判定
		return
	}
	rand.Seed(time.Now().Unix())
	r := rand.Intn(q-p)
	inputs[p],inputs[p+r] = inputs[p+r],inputs[p]
	m := inputs[p]
	i , j := p,p
	for j < q {
		if inputs[j+1] <=  m {
			inputs[i+1],inputs[j+1] = inputs[j+1],inputs[i+1]
			i ++
		}
		j ++
	}
	inputs[p] ,inputs[i] = inputs[i],inputs[p]  	// 首元素交换
	quick(inputs,p,i-1)
	quick(inputs,i+1 ,q)
	return
}

func MRandQuick(inputs [] int)  {
	randQuick(inputs , 0,len(inputs)-1)
}
