package my_sort


// 归并排序 ，先归后并
// T(n) = 2T(n/2) + Θ(n)
// 空间复杂度Θ(nlgn), Θ(n),Θ(1) 取决于具体实现


// 归并排序
// 空间复杂度 Θ(n)
func MMerger(input []int)  {
	if len(input) == 0 {
		panic("empty list")
		return
	}
	tmp := make([]int,len(input))
	merger(input,tmp)
}

// 空间复杂度 Θ(n)
func MMergerTmp(input,tmp []int)  {
	if len(input) == 0  || len(tmp) < len(input){
		panic("empty list")
		return
	}
	merger(input,tmp)
}



func merger(input,tmp []int)  {
	if  len(input) == 1 {
		return
	}

	// 归操作
	left := input[: len(input)/2]
	right := input[ len(input)/2:]
	merger(left,tmp[:len(left)])
	merger(right,tmp[:len(right)])

	// 并操作
	for l ,r ,t := 0,0,0 ;;{
		if l < len(left) && r < len(right){
			if left[l] <= right[r] {
				tmp[t] = left[l]
				l ++
			}else {
				tmp[t] = right[r]
				r ++
			}
		}else if l < len(left) {	  // 其中一个比较完成
			tmp[t] = left[l]
			l ++
		} else if r < len(right) {
			tmp[t] = right[r]
			r ++
		}else {
				break
		}
		t ++
	}
	copy(left,tmp)
	copy(right,tmp[len(left):])
	return
}

