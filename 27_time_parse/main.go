package main

import (
	"fmt"
	"time"
)

func main() {
	str := "2020-02-20  11:15:40"
	t, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t)
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, err = time.Parse(longForm, "Feb 20, 2020 at 11:15am (PST)")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t)

	const shortForm = "2006-Jan-02"
	t, _ = time.Parse(shortForm, "2020-Feb-20")
	fmt.Println(t)

	t1 := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local)
	fmt.Printf("Go launched at %s\n", t1)
}
