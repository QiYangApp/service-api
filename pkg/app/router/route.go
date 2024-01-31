package router

type inject interface {
	// PostConstruct Triggered after dependency injection is completed. You can continue to decorate the controller here.
	PostConstruct()
}

type Inject struct{}

func (c *Inject) PostConstruct() {}
