package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println("date")

	// 获取当前时间
	now := time.Now()
	fmt.Println("now=", now)

	// 通过date日期获取年月日
	fmt.Println("年=", now.Year())
	fmt.Println("月=", now.Month())
	fmt.Println("月=", int(now.Month()))
	fmt.Println("日=", now.Day())
	fmt.Println("时=", now.Hour())
	fmt.Println("分=", now.Minute())
	fmt.Println("秒=", now.Second())

	// 格式化日期和时间
	// 第一种
	fmt.Printf("年月日: %d-%d-%d %d:%d:%d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	dateStr := fmt.Sprintf("%d-%d-%d %d:%d:%d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println("dateStr=", dateStr)

	// 第二种
	// 2006-01-02 15:04:05 这个是固定的参考值
	fmt.Println("年月日时分秒", now.Format("2006-01-02 15:04:05"))
	fmt.Println("年月日", now.Format("2006-01-02"))
	fmt.Println("时分秒", now.Format("15:04:05"))
	fmt.Println("年", now.Format("2006"))

	// 时间常量，
	/**
	const (
		Nanosecond  Duration = 1
		Microsecond          = 1000 * Nanosecond
		Millisecond          = 1000 * Microsecond
		Second               = 1000 * Millisecond
		Minute               = 60 * Second
		Hour                 = 60 * Minute
	)
	*/
	// 每隔1s打印一个数组
	i := 0
	for {
		i++
		fmt.Println("i=", i)
		// 休眠
		// time.Sleep(time.Second)

		// 100ms
		time.Sleep(time.Millisecond * 100)
		if i == 5 {
			break
		}
	}

	// 获取当前unix的时间戳 和 unixnano 时间戳 （可以获取随机数字）
	// 获取到当前时间的 unix 时间戳 和 unixnano 时间戳
	fmt.Printf("unix=%v,unixnano=%v", now.Unix(), now.UnixNano())

	// 统计函数运行使用的时间

	start := time.Now().Unix()
	testTime()
	end := time.Now().Unix()

	fmt.Println("testTime() 运行消耗时间为", end-start)

}

func testTime() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += strconv.Itoa(i)
	}
}
