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
