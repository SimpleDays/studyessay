package algorithm

import (
	"fmt"
	"testing"
)

// 描述
// 模拟筛选出把可以淘汰的服务器
// 用二维数组表示当前服务器的依赖关系
// 用一维数组表示提供需要淘汰的服务器
// 通过算法淘汰可以淘汰的服务器序列
// 规则：如果提供淘汰服务器之外还存在一定依赖关系则无法删除
// 例如：依赖服务器序列号：[[0,1,2], [0,4], [5,6]]  说明：0,1,2相互依赖，同时 0，4号机器也相互依赖，5，6也相互依赖
// 此时提供待淘汰服务器序列数组：[0,1,2,5,6] 因 0，1，2，4 相互依赖而 4 并非在淘汰列表中 所以 0，1，2 无法直接淘汰，5，6 因在一组依赖下，没有其他依赖
// 所以 5，6 可以 淘汰，最终返回可淘汰服务器序列数组：[5,6]

var toBeEliminatedList = []int{0, 1, 2, 5, 6}

var relationServers [][]int

func init() {
	// 初始化依赖服务器列表
	rss := make([][]int, 0)

	rss = append(rss, []int{0, 1, 2})

	rss = append(rss, []int{0, 4, 5})

	rss = append(rss, []int{5, 6})

	relationServers = rss
}

func TestRelationServer(t *testing.T) {
	t.Log(relationServers)
	t.Log(toBeEliminatedList)
	EliminatedServers(toBeEliminatedList)
}

func EliminatedServers(list []int) []int {

	// 1、把依赖服务器的并集转换成map类型
	UnionToMap()

	return nil
}

func UnionToMap() map[int][]int {
	unionMap := make(map[int][]int)

	rsLen := len(relationServers)

	relations := make(map[int][]int)

	for i := 0; i < rsLen; i++ {
		rss := relationServers[i]

		relations[i] = append(relations[i], []int{i}...)

		for j := i + 1; j < len(rss); j++ {
			y := rss[j]

			ok := hasRelation(y)

			if ok {
				relations[i] = append(relations[i])
			}
		}
	}

	fmt.Println(relations)
	return unionMap
}

func hasRelation(y int, rrs []int) (bool, []int) {

	flag := false
	r := make([]int, 0)

	for index, rss := range relationServers {
		for _, rs := range rss {
			if rs == y {
				if rrs == nil || len(rrs) == 0 {
					flag = true
					r = append(r, index)
				} else {
					var f = true
					for _, c := range rrs {
						if index == c {
							f = false
							break
						}
					}

					if f {
						flag = true
						rrs = append(rrs, index)
						r = rrs
					}
				}
				break
			}
		}
	}

	return flag, r
}
