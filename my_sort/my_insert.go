package my_sort

// 插入排序
// 时间复杂度  T(n) = Θ(n^2)
// 正序（最好）T(n) = Θ(n)
// 倒叙（最坏）T(n) = Θ(n)
// 空间复杂度  S(n) = Θ(1)
// 稳定排序
// 测试发现 在待排序数较小时，增长较慢(约10倍) ，较大时增长非常快


// 插入排序算法
func MInsert(input []int) {
	for i := 1 ; i < len(input) ; i ++ {
		v := input[i]
		for j := i-1 ; j >= 0  ; j -- {
			if v >= input[j]  { // 等于保证排序稳定性
				copy(input[j+2:],input[j+1:i])
				input[j+1] = v
				break
			}
			if j == 0 {         //插入最小值
				copy(input[1:],input[:i])
				input[0] = v
			}
		}
	}
	return
}