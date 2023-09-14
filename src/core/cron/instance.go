package cron

import "github.com/robfig/cron/v3"

type Instance struct {
	cron *cron.Cron
}

func (i *Instance) Init() *Instance {
	i.cron = cron.New(
		cron.WithSeconds(),
		cron.WithLogger(Logger{}),
	)

	return i
}

func (i *Instance) Start() {
	i.cron.Start()
}

func (i *Instance) Stop() {
	i.cron.Stop()
}

func (i *Instance) AddFun(time TimeInterface, f func()) (cron.EntryID, error) {
	return i.cron.AddFunc(time.toString(), f)
}

func (i *Instance) AddJob(time TimeInterface, job cron.Job) (cron.EntryID, error) {
	return i.cron.AddJob(time.toString(), job)
}
