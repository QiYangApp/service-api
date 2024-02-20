package events

// laravel events

type EventFunc func()

type Event struct {
	Key   string
	Fun   EventFunc
	Share bool
}

type Dispatcher struct {
	events map[string]Event
}
