package main

import (
	"fmt"
)

func main() {
	s := make([]int, 3, 5)
	s[0] = 99
	s1 := append(s, 10)      //append向切片尾部添加数据，返回新的切片对象，数据被追加到原底层数组，
	s2 := append(s1, 20, 30) //如超过切片cap限制，则为新切片对象重新分配底层数组

	fmt.Println(s, len(s), cap(s))
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))

	x := [][]int{
		{1, 2},
		{10, 20, 30},
		{100},
	} //x是一个切片数组
	fmt.Println(x[1])
	x[2] = append(x[2], 3, 4) //对切片数组中的切片append，会重新分配底层数组，因为原切片len=map了
	fmt.Println(x[2])

	// for i,value := range s2{
	// 	if value ==  0 {
	// 		s2 = append(s2[:i], s2[i+1:]...)
	// 		i--
	// 	}
	// 	if i == index{
	// 		s2 = append(s2[:i], s2[i+1:])
	// 	}
	// }
	// for i := 0; i < len(slice); i++ {
	// 	fmt.Printf("slice[%v]=%v ", i, slice[i])
	// }
}
