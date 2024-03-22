package auth

type ConfigType[T any] interface {
	Value() T
}

// Config represents login config as far as the db is concerned
type Config[T any] struct {
	Val T
}

func (c *Config[T]) Value() T {
	return c.Val
}
