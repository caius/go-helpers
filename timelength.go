package main

import (
	"fmt"
)

type TimePeriod struct{
	seconds int
	minutes int
	hours int
	days int
	weeks int
}
// TODO: add inspect to TimePeriod

func TimeLength(num int) (output TimePeriod) {
	return
}

func main() {
	// 3 weeks, 5 days, 5 hours, 34 minutes and 2 seconds
	sec := 1
	min := sec * 60
	hour := min * 60
	day := hour * 24
	week := day * 7

	input := week * 3 + 5 * day + 7 * hour + 34 * min + 2 * sec
	output := TimeLength(input)
	fmt.Println(output)
}
