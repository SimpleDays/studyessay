package algorithm

import (
	"fmt"
	"strings"
	"sync"
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
	gw := sync.WaitGroup{}

	gw.Add(101)

	go func() {
		for {
			s <- i
			if i > 100 {
				break
			}

			if i%2 == 0 {
				fmt.Println("go 协程 1 打印偶数 ： ", i)
				i++
				gw.Done()
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
				fmt.Println("go 协程 2 打印奇数 ： ", i)
				i++
				gw.Done()
			}
		}
	}()

	gw.Wait()
	fmt.Println("over")
}

/*
描述：
存在两个字符串数组， 一个是大写A-Z, 另一个小写a-z
存在两个协程，分别打印上述两个数组
要求输出结果 AaBbCc.....Zz
*/

func TestPrintLetter(t *testing.T) {
	capitals := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	lowercase := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	var i = 0
	s1 := make(chan int)
	s2 := make(chan int)
	gw := sync.WaitGroup{}

	gw.Add(52)

	go func() {
		for {
			<-s1
			if i == len(capitals) {
				break
			}

			fmt.Print(capitals[i])
			gw.Done()
			s2 <- i
		}
	}()

	time.Sleep(2 * time.Second)

	go func() {
		for {
			s1 <- i
			<-s2
			if i == len(lowercase) {
				break
			}

			fmt.Print(lowercase[i])
			i++
			gw.Done()
		}
	}()

	gw.Wait()
	fmt.Println("")
	fmt.Println("over")
}

func TestAbc(t *testing.T) {

	sw := sync.WaitGroup{}

	sw.Add(52)

	go func() {
		s := "abcdefghijklmnopqrstuvwxyz"

		for _, d := range s {
			fmt.Print("\"" + string(d) + "\",")
			sw.Done()
		}

		fmt.Println("")

		for _, d := range s {
			m := strings.ToUpper(string(d))
			fmt.Print("\"" + m + "\",")
			sw.Done()
		}
	}()

	sw.Wait()
	fmt.Println("")
	fmt.Println("over")
}
