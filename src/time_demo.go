package main

import (
	"fmt"
	"time"
)

func main() {
	myTimeString := "2015-01-01 00:00:00"
	timeTpl := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local") // 指定时区
	theTime, _ := time.ParseInLocation(timeTpl, myTimeString, loc)
	timeStamp := theTime.Unix() // 日期转时间戳
	fmt.Println("This time is", theTime)
	fmt.Println("This time stamp is", timeStamp)

	thisTime := time.Unix(timeStamp, 0).Format(timeTpl) // 时间戳转日期
	fmt.Println("This time is", thisTime)
}
