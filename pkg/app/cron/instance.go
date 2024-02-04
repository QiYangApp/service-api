package cron

import "github.com/robfig/cron/v3"

type Manage struct {
	cron *cron.Cron
}

func (i *Manage) Init() *Manage {
	i.cron = cron.New(
		cron.WithSeconds(),
		cron.WithLogger(Logger{}),
	)

	i.cron.Start()

	return i
}

func (i *Manage) Start() {
	i.cron.Start()
}

func (i *Manage) Stop() {
	i.cron.Stop()
}

func (i *Manage) AddFun(time TimeInterface, f func()) (cron.EntryID, error) {
	return i.cron.AddFunc(time.toString(), f)
}

func (i *Manage) AddJob(time TimeInterface, job cron.Job) (cron.EntryID, error) {
	return i.cron.AddJob(time.toString(), job)
}
