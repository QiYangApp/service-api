package cron

import (
	"sync"
)

var singleton *Manage
var once = sync.Once{}

func Client() *Manage {
	once.Do(func() {
		singleton = &Manage{}
		singleton.Init()
	})

	return singleton
}

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
