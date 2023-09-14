package cron

import (
	"sync"
)

var Singleton *Instance
var once = sync.Once{}

func S() *Instance {
	once.Do(func() {
		Singleton = &Instance{}
		Singleton.Init()
	})

	return Singleton
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
