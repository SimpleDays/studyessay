/*
 * @Author: dongliang 342479980@qq.com
 * @Date: 2022-10-09 19:16:12
 * @LastEditors: dongliang 342479980@qq.com
 * @LastEditTime: 2022-10-09 23:05:22
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
