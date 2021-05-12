package algorithm

import (
	"fmt"
	"testing"
	"time"
)

/*
描述：
通过2个协程，把0-100的数按顺序打印出来，并且一个协程只打印奇数，另一个打印偶数
*/

func TestPrintNum(t *testing.T) {

	var i = 0
	s := make(chan int)

	go func() {
		for {
			s <- i
			if i > 100 {
				break
			}

			if i%2 == 0 {
				fmt.Println("go 协程 1 打印偶数 ： ", i)
				i++
			}
		}
	}()

	go func() {
		for {
			<-s
			if i > 100 {
				break
			}

			if i%2 == 1 {
				fmt.Println("go 协程 2 打印奇数 ：", i)
				i++
			}
		}
	}()

	fmt.Println("over")
	time.Sleep(1 * time.Hour)
}
