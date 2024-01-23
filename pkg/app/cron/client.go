package cron

import (
	"github.com/robfig/cron/v3"
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

func Client() *cron.Cron {
	return Instance().cron
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
