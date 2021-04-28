/**
 * @Author: XiaoLongBao
 * @Description: OOP的面试题目
 * @File:  oop
 * @Program: helloworld
 * @Date: 2021-04-28 15:54
 */
package interview

import (
	"fmt"
	"testing"
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
