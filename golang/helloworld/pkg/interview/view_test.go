/**
 * @Author: XiaoLongBao
 * @Description: 面试题目
 * @File:  view_test
 * @Program: helloworld
 * @Date: 2021-04-28 15:54
 */
package interview

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")

}
func TestTeacher(t *testing.T) {
	teacher := Teacher{}
	teacher.ShowA()
}

/*
 上述解释： 被组合的类型People所包含的方法虽然升级成了外部类型Teacher这个组合类型的方法（一定要是匿名字段），
 但它们的方法(ShowA())调用时接受者并没有发生变化。此时People类型并不知道自己会被什么类型组合，
 当然也就无法调用方法时去使用未知的组合者Teacher类型的功能。
*/

func TestDeferPanic(t *testing.T) {
	defer func() {
		fmt.Println("打印前")
	}()
	defer func() {
		fmt.Println("打印中")
	}()
	defer func() {
		fmt.Println("打印后")
	}()

	panic("触发异常")
}

/*
defer 是后进先出。 panic 需要等defer 结束后才会向上传递。
出现panic恐慌时候，会先按照defer的后入先出的顺序执行，最后才会执行panic。
*/

func TestDeferFunc(t *testing.T) {
	var calc = func(index string, a, b int) int {
		ret := a + b
		t.Log(index, a, b, ret)
		return ret
	}

	a := 1
	b := 2

	defer calc("1", a, calc("10", a, b))

	a = 0

	defer calc("2", a, calc("20", a, b))

	b = 1
}

func TestDeferFunc2(t *testing.T) {
	t.Log("main", inc(t))
}

func inc(t *testing.T) int {
	test := &test{num: 0}
	defer test.Inc(3, t).Inc(2, t).Inc(1, t)
	t.Log("inc", test.num)
	return test.num
}

type test struct {
	num int
}

func (test *test) Inc(flag int, t *testing.T) *test {
	test.num++
	t.Log("test", flag, test.num)
	return test
}

/*
虽然defer是后进先出，但是defer里面函数 "calc" 优先从上到下运行。
第一个defer里面的calc函数执行结果打印
1-- 10，1，2，3
第二个defer里面的calc函数执行结果打印
2-- 20，0，2，2
这时候在依据defer的后进先出的原则开始继续计算
第二个defer优先执行
3-- 2，0，2，2
然后执行第一个defer函数
4-- 1，1，3，4
ps：注意 a 和 b 的作用域
*/

func TestChanSelect(t *testing.T) {
	intChan := make(chan int, 1)
	stringChan := make(chan string, 1)

	intChan <- 1
	stringChan <- "hello"

	select {
	case v := <-intChan:
		t.Log(v)
	case v := <-stringChan:
		panic(v)
	}
}

/*
select会随机选择一个可用通用做收发操作。所以代码是有肯触发异常，也有可能不会。
单个chan如果无缓冲时，将会阻塞。但结合 select可以在多个chan间等待执行。
三个原则：
1、select 中只要有一个case能return，则立刻执行。

2、当如果同一时间有多个case均能return则伪随机方式抽取任意一个执行。

3、如果没有一个case能return则可以执行”default”块
*/

func TestArray(t *testing.T) {
	a := make([]int, 0)
	a = append(a, 1, 2, 3)
	t.Log(a)

	a2 := make([]int, 5)
	a2 = append(a2, 1, 2, 3)
	t.Log(a2)
}

/*
make初始化是由默认值的哦
*/

func TestChar(t *testing.T) {
	var c rune

	t.Log(c)

	t.Log(string(c))

	t.Log(string(c) == "")

	c = 'a'

	t.Log(c)

	t.Log(string(c))
}

/*
golang char的简单理解
*/

type UserAage struct {
	ages map[string]int
	sync.Mutex
}

func (u *UserAage) Add(name string, age int) {
	u.Lock()

	defer u.Unlock()

	u.ages[name] = age
}

func (u *UserAage) Get(name string) int {
	if age, ok := u.ages[name]; ok {
		return age
	}

	return -1
}

func TestMapThreadSafe(t *testing.T) {
	u := &UserAage{
		ages: make(map[string]int),
	}
	go u.Add("aa", 333)

	go func(t *testing.T) {
		defer func() { //defer就是把匿名函数压入到defer栈中，等到执行完毕后或者发生异常后调用匿名函数
			err := recover() //recover是内置函数，可以捕获到异常
			if err != nil {  //说明有错误
				fmt.Println("err=", err)
				//当然这里可以把错误的详细位置发送给开发人员
				//send email to admin
			}
		}()

		age := u.Get("aa")
		t.Log(age)
	}(t)

	time.Sleep(100 * time.Hour)
}

/*
可能会出现fatal error: concurrent map read and map write
*/

type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{}, len(set.s))
	go func() {
		set.RLock()
		for elem, v := range set.s {
			ch <- elem
			println("Iter", elem, v)
		}
		close(ch)
		set.RUnlock()
	}()
	return ch
}

func TestIter(t *testing.T) {
	th := &threadSafeSet{
		s: []interface{}{
			"1", "2",
		},
	}

	v := <-th.Iter()
	fmt.Sprintf("%s%v", "ch", v)
	// time.Sleep(100 * time.Hour)
}

/*
看到这道题，我也在猜想出题者的意图在哪里。
chan?sync.RWMutex?go?chan缓存池?迭代?
所以只能再读一次题目，就从迭代入手看看。
既然是迭代就会要求set.s全部可以遍历一次。
但是chan是为缓存的，那就代表这写入一次就会阻塞。
*/
