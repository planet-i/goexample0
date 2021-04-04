package main

import "fmt"

func main() {
	s1 := make([]int, 3, 5)
	s2 := make([]int, 5)
	s3 := []int{1, 2, 5: 6} //s3 := [...]int{1,2,5:6}数组的显示初始化形式

	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s3, len(s3), cap(s3))

}
