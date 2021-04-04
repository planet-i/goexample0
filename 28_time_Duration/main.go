package main

import (
	"fmt"
	"time"
)

func main(){

    var wait time.Duration = 500 * time.Millisecond
 
    startingTime := time.Now().UTC()
    time.Sleep(600 * time.Millisecond)
	endingTime := time.Now().UTC()
    var duration time.Duration = endingTime.Sub(startingTime)
 
    if duration >= wait {
        fmt.Printf("Wait %v\n", wait)
		fmt.Printf("Duration [%v]\n", duration)
//------将Duration类型转化为Go中的int64或float64类型
		fmt.Printf("Nanoseconds [%d]\n", duration.Nanoseconds())
		fmt.Printf("Microseconds [%d]\n", duration. Microseconds())
		fmt.Printf("Milliseconds [%d]\n", duration.Milliseconds())
		fmt.Printf("Seconds [%.3f]\n", duration.Seconds())
		fmt.Printf("Minutes [%.3f]\n", duration.Minutes())
		fmt.Printf("Hours[%.3f]\n", duration.Hours())
	}else {
        fmt.Printf("Time DID NOT Elapsed : Wait[%d] Duration[%d]\n", wait, duration)
	}

	d1 := time.Since(endingTime)
	d2 := time.Until(endingTime)
	fmt.Printf("Since %v\n", d1)  //现在时间-参数时间
	fmt.Printf("Until %v\n", d2)  //参数时间-现在时间
	UseOther()
}

func UseOther(){
	d, err := time.ParseDuration("1h15m30.918273645s")//带符号的十进制数字序列->Duration型
	if err != nil {
		panic(err)
	}

	round := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}

	for _, m := range round {
		fmt.Printf("   d.Round(%6s) = %s\n", m, d.Round(m).String())  //类似四舍五入
		fmt.Printf("d.Truncate(%6s) = %s\n", m, d.Truncate(m).String())//类似向下取整
	}//若m<0 ,返回d不变
	t1 := time.Date(2016, time.August, 15, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2017, time.February, 16, 0, 0, 0, 0, time.UTC)
	fmt.Println(t2.Sub(t1).String()) //Duration -> 相似于“ 72h3m0.5s”形式的字符串
}
/* type Duration int64
const (
    Nanosecond Duration = 1
    Microsecond = 1000 * Nanosecond
    Millisecond = 1000 * Microsecond
    Second = 1000 * Millisecond
    Minute = 60 * Second
    Hour = 60 * Minute
) 
Printf和%v操作符，本地化显示一个 Duration 类型
- 纳秒值除以1e6得到了int64类型表示的毫秒值
- 秒值乘以1e3得到了float64类型表示的毫秒值
*/