/*
 * @Author: dongliang 342479980@qq.com
 * @Date: 2022-10-09 19:16:12
 * @LastEditors: dongliang 342479980@qq.com
 * @LastEditTime: 2022-10-09 19:48:24
 * @FilePath: /helloworld/pkg/algorithm/sort_test.go
 * @Description: 排序算法记录
 */
package algorithm

import "testing"

// 选择排序
// 原理：i ~ N-1 上找到最小值的小标， 然后进行交换，一直这么操作下去
// 时间复杂度 O(N²)
func TestSelectSort(t *testing.T) {
	intArray := []int{3, 6, 99, 1, 2, 0, 55, 789, 321, 21, 2, 6}

	for i := 0; i < len(intArray)-1; i++ { //i ~ N-1
		minIndex := i

		for j := i + 1; j < len(intArray); j++ { // i+1 ~ N 找到比当前位置更小的数的索引位
			if intArray[j] < intArray[minIndex] {
				minIndex = j
			}
		}

		swap(intArray, i, minIndex)
	}

	t.Log(intArray)
}

func swap(intArray []int, i int, j int) {
	tmp := intArray[i]
	intArray[i] = intArray[j]
	intArray[j] = tmp
}
