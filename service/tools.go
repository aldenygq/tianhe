package service

import (
	"time"
	"unicode"
)

func CompareTwoDay(day1,day2 string) int64{
	// 创建两个日期
	d1, _ := time.Parse("2006-01-02", day1)
	d2, _ := time.Parse("2006-01-02", day2)
	
	// 比较日期
	if d1.Before(d2) {
		return 0
	} else if d1.After(d2) {
		return 1
	} else {
		return 2
	}
	return 3
}

func getDay() string {
	now := time.Now().Format("2006-01-02")
	
	// fmt.Println("Current date:", now.Format("2006-01-02"))
	return now
}
func IsContainChinese(str string) bool {
	for _, r := range str {
		if unicode.In(r, unicode.Scripts["Han"]) {
			return true
		}
	}
	return false
}