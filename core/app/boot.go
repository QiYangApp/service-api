package app

type EventFunc func(*App)

type Event struct {
	Key   string
	Fun   EventFunc
	Share bool
}
