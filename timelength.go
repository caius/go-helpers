package main

import (
	"fmt"
)

var (
	Sec = 1
	Min = Sec * 60
	Hour = Min * 60
	Day = Hour * 24
	Week = Day * 7
)

type TimePeriod struct{
	seconds int
	minutes int
	hours int
	days int
	weeks int
}
// TODO: add inspect to TimePeriod

func (tp TimePeriod) String() string {
	return fmt.Sprintf("%d weeks, %d days, %d hours, %d minutes and %d seconds", tp.weeks, tp.days, tp.hours, tp.minutes, tp.seconds)
}

func rb_divmod(x, y int) (q, r int) {
	q = x / y
	r = x % y
	return
}

func NewTimePeriod(num int) (output TimePeriod) {
	output.weeks, num = rb_divmod(num, Week)
	output.days, num = rb_divmod(num, Day)
	output.hours, num = rb_divmod(num, Hour)
	output.minutes, num = rb_divmod(num, Min)
	output.seconds = num

	return
}

func main() {
	input := (Week * 3) + (Day * 0) + (Hour * 7) + (Min * 5) + (Sec * 2)

	output := NewTimePeriod(input)
	fmt.Println(output)
}
