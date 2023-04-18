package helpers

func init() {
	once.Do(func() {
		PathInstance = new(Path).init()
	})
}
