/*
 * @Author: dongliang 342479980@qq.com
 * @Date: 2022-10-09 19:16:12
 * @LastEditors: dongliang 342479980@qq.com
 * @LastEditTime: 2022-10-10 00:58:55
 * @FilePath: /helloworld/pkg/algorithm/sort_test.go
 * @Description: 排序算法记录
 */
package algorithm

import "testing"

// 选择排序
// 原理：i ~ N-1 上找到最小值的小标， 然后进行交换，一直这么操作下去
// 时间复杂度 O(N²)
// 空间复杂度O(1)
func TestSelectSort(t *testing.T) {
	intArray := []int{3, 6, 99, 1, 2, 0, 55, 789, 321, 21, 2, 6}

	SelectSort(intArray)

	t.Log(intArray)
}

// 冒泡排序
// 原理： 比较相邻的元素。如果第一个比第二个大，就交换他们两个。
// 对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。
// 时间复杂度 O(N²)
// 空间复杂度O(1)
func TestBubbleSort(t *testing.T) {
	intArray := []int{3, 6, 99, 1, 2, 0, 55, 789, 333, 21, 2, 6, 12345}

	BubbleSort(intArray)

	t.Log(intArray)
}

// 插入排序
// 原理：在已排序序列中从后向前扫描，找到相应位置并插入。0~0， 0~1，。。。。。。0~N-1
// 时间复杂度 O(N²)
// 空间复杂度O(1)
func TestInsertionSort(t *testing.T) {
	intArray := []int{3, 6, 99, 1, 2, 0, 55, 789, 333, 21, 2, 6, 1234, 66}

	InsertionSort(intArray)

	t.Log(intArray)
}

// 算法题： 从一个int数组中找出一种数是出现奇数次，其他都是出现偶数次，要求时间复杂度O(N), 空间复杂度O(1)
func TestOnlyOddNumber(t *testing.T) {
	intArray := []int{2, 1, 3, 1, 3, 1, 3, 2, 1}

	eor := 0

	for _, num := range intArray {
		eor = eor ^ num
	}

	t.Logf("出现奇数次的数是: %v", eor)
}

// 算法题： 从一个int数组中找出二种数是出现奇数次，其他都是出现偶数次，要求时间复杂度O(N), 空间复杂度O(1)
func TestTowTimesOddNumber(t *testing.T) {
	intArray := []int{2, 1, 3, 1, 3, 1, 3, 2, 1, 6}

	eor := 0

	for _, num := range intArray {
		eor = eor ^ num
	}

	// eor = a^b
	// eor != 0
	// eor必然有一个位置上是1

	rightOne := eor & (^eor + 1) // 提取出最右的1  eor 与 eor取反 加 1

	onlyOne := 0 // eor'

	for _, num := range intArray {
		if num&rightOne == 1 {
			onlyOne = onlyOne ^ num
		}
	}

	t.Logf("出现二种数是奇数次的分别是：%v 和 %v", onlyOne, (eor ^ onlyOne))

}
