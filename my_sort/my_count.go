package my_sort

// 计数排序
// 时间复杂度 Θ(k+n) 空间复杂度 Θ(k+n) [n 待排元素个数 ， k 元素值区间]
// 有局限性 ，适用于变化不大的数字
// 稳定性

func MCount(inputs [] uint8 ) []uint8 {
	var count [256]uint8
	outs  := make([]uint8,len(inputs))
	for i := 0;i < len(inputs) ;i ++ {	//  计数
		count[inputs[i]] ++
	}

	for i :=1 ;i <len(count);i++ {  	// 元素位置
		count[i] = count[i-1] + count[i]
	}

	for i := len(inputs) - 1 ;i >=0 ;i-- {
		outs[count[inputs[i]]-1] = inputs[i]
		count[inputs[i]] --
	}
	return outs
}

// 计数算法改进 ，处理整出排序	
func MCountImprove(){}