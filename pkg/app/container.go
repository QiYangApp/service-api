package app

type Container struct {
	instances map[string]any
}

func (c *Container) Instance(key string, instance any) *Container {
	c.instances[key] = instance

	return c
}

func (c *Container) Make(key string, instance any) *Container {
	c.instances[key] = instance

	return c
}
