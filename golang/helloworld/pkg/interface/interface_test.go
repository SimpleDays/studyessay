/**
 * @Author: XiaoLongBao
 * @Description: 对interface的了解学习
 * @File:  interface_test
 * @Program: hello world
 * @Date: 2021-03-17 09:27
 */
package _interface

import (
	"reflect"
	"testing"
)

/*
 * Go interface 的一个 “坑” 及原理分析 https://mp.weixin.qq.com/s/vNACbdSDxC9S0LOAr7ngLQ
 * 针对 interface 使用疑惑。
 */

// todo: 事例1
func TestInterfaceEg1(t *testing.T) {
	var v interface{}
	v = (*int)(nil)

	// 为什么不是 true。明明都已经强行置为 nil 了。是不是 Go 编译器有问题？
	t.Log(v == nil)
}

// todo: 事例2
func TestInterfaceEg2(t *testing.T) {
	var data *byte
	var in interface{}

	// 刚刚声明出来的 data 和 in 变量，确实是输出结果是 nil，判断结果也是 true
	t.Log(data, data == nil)
	t.Log(in, in == nil)

	in = data

	// 奇怪是怎么把变量 data 一赋予给变量 in，世界就变了？输出结果依然是 nil，但判定却变成了 false
	t.Log(in, in == nil)
}

/*
 * 分析：
 * interface 判断与想象中不一样的根本原因是，interface 并不是一个指针类型，虽然他看起来很像，以至于误导了不少人。
 * interface 共有两类数据结构: runtime.iface 和 runtime.eface
 * runtime.eface 结构体：表示不包含任何方法的空接口，也称为 empty interface
 * runtime.iface 结构体：表示包含方法的接口
 * 两者相应的底层数据结构
 * type eface struct {
 *   _type *_type
 *   data  unsafe.Pointer
 * }
 *
 * type iface struct {
 *   tab  *itab
 *   data unsafe.Pointer
 * }
 *
 * interface 不是单纯的值，而是分为类型和值
 * 所以传统认知的此 nil 并非彼 nil，必须得类型和值同时都为 nil 的情况下，interface 的 nil 判断才会为 true
 * 1、可以利用反射（reflect）来做 nil 的值判断，在反射中会有针对 interface 类型的特殊处理
 * 2、对值进行 nil 判断，再返回给 interface 设置
 * 3、返回具体的值类型，而不是返回 interface
 */

// todo: 利用反射（reflect）来做 nil 的值判断
func TestInterfaceReflect(t *testing.T) {
	var data *byte
	var in interface{}

	in = data
	t.Log(in, func(i interface{}) bool {
		vi := reflect.ValueOf(i)
		// 类型如果为指针时：
		if vi.Kind() == reflect.Ptr {
			return vi.IsNil()
		}

		return false
	}(in))
}
