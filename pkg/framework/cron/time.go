package cron

import (
	"fmt"
	"strconv"
)

type TimeInterface interface {
	toString() string

	AtMinute(minute int) TimeInterface

	EverySecond(second int) TimeInterface

	EveryMinute(minute int) TimeInterface

	EveryDay(day int) TimeInterface

	AtDay(day int) TimeInterface
}

type Time struct {
	TimeInterface

	second string
	minute string
	hour   string
	day    string
	month  string
	week   string
}

func (t *Time) toString() string {
	return fmt.Sprintf("%s %s %s %s %s %s", t.second, t.minute, t.hour, t.day, t.month, t.week)
}

func (t *Time) EverySecond(second int) TimeInterface {
	if second != 0 {
		t.second = fmt.Sprintf("%s/%d", "*", second)
	}

	return t
}

func (t *Time) EveryMinute(minute int) TimeInterface {
	if minute != 0 {
		t.minute = fmt.Sprintf("%s/%d", "*", minute)
	}

	return t
}

func (t *Time) EveryDay(day int) TimeInterface {
	if day != 0 {
		t.day = fmt.Sprintf("%s/%d", "*", day)
	}

	return t
}

func (t *Time) AtDay(day int) TimeInterface {
	if day != 0 {
		t.day = strconv.Itoa(day)
	}

	return t
}

func (t *Time) AtMinute(minute int) TimeInterface {
	if minute != 0 {
		t.minute = strconv.Itoa(minute)
	}

	return t
}
