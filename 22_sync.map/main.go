package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	//Store;   func (m *Map) Store(key, value interface{})
	m.Store(1, "a")
	m.Store(2, "b")

	//LoadOrStore
	//若key不存在，则存入key和value，返回false和输入的value
	v, ok := m.LoadOrStore("1", "aaa")
	fmt.Println(ok, v) //false aaa

	//若key已存在，则返回true和key对应的value，不会修改原来的value
	v, ok = m.LoadOrStore(1, "bbb")
	fmt.Println(ok, v) //true a

	//Load
	//load返回给定key的value，一个bool值表示key值是否在map中存在；若key不存在，返回nil
	v, ok = m.Load(1)
	if ok {
		fmt.Println("it's an existing key,value is ", v)
	} else {
		fmt.Println("it's an unknown key")
	}

	//Range；   func (m *Map) Range(f func(key, value interface{}) bool)
	//遍历sync.Map, 要求输入一个func作为参数
	f := func(k, v interface{}) bool {
		//这个函数的入参、出参的类型都已经固定，不能修改
		//可以在函数体内编写自己的代码，调用map中的k,v
		//如果函数返回false，则range停止迭代
		fmt.Println(k, v)
		return true
	}
	m.Range(f)

	//Delete
	m.Delete(1)
	fmt.Println(m.Load(1))

}
