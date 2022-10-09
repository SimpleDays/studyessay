package algorithm

// 选择排序
// 原理：i ~ N-1 上找到最小值的小标， 然后进行交换，一直这么操作下去
// 时间复杂度 O(N²)
// 空间复杂度O(1)
func SelectSort(intArray []int) {

	if len(intArray) == 0 {
		return
	}

	for i := 0; i < len(intArray)-1; i++ { //i ~ N-1
		minIndex := i

		for j := i + 1; j < len(intArray); j++ { // i+1 ~ N 找到比当前位置更小的数的索引位
			if intArray[j] < intArray[minIndex] {
				minIndex = j
			}
		}

		tmp := intArray[i]
		intArray[i] = intArray[minIndex]
		intArray[minIndex] = tmp
	}
}

// 冒泡排序
// 原理： 比较相邻的元素。如果第一个比第二个大，就交换他们两个。
// 对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。
// 时间复杂度 O(N²)
// 空间复杂度O(1)
// 关于异或(^)直接交换数据 (两个数是独立的内存区域) （注意： 如果同一个内存区域进行异或会变成0)
// 假设 a = 甲 ，b = 乙
// a = a ^ b = 甲 ^ 乙
// b = a ^ b = 甲 ^ 乙 ^ 乙 = （根据异或的结合律/交换律） 甲 ^ (乙 ^ 乙) = 甲 ^ 0 = 甲
// a = a ^ b = 甲 ^ 乙 ^ 甲 = （根据异或的结合律/交换律） (甲 ^ 甲) ^ 乙 = 0 ^ 乙 = 乙
func BubbleSort(intArray []int) {

	if len(intArray) == 0 {
		return
	}

	for i := len(intArray) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if intArray[j] > intArray[j+1] {
				intArray[j] = intArray[j] ^ intArray[j+1]
				intArray[j+1] = intArray[j] ^ intArray[j+1]
				intArray[j] = intArray[j] ^ intArray[j+1]
			}
		}
	}
}
