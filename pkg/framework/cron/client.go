package cron

import (
	"sync"
)

var singleton *Manage
var once = sync.Once{}

func Instance() *Manage {
	once.Do(func() {
		singleton = &Manage{}
	})

	return singleton
}

var Client = Instance().cron

func TaskTime() TimeInterface {
	return &Time{
		second: "*",
		minute: "*",
		hour:   "*",
		day:    "*",
		month:  "*",
		week:   "*",
	}
}
