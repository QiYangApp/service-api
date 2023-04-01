package token

type CacheMethodInterface interface {
	init() CacheMethodInterface
	store(token string) (string, error)
	parser(token string) (Claims, error)
	exists(token string) bool
	refresh(token string) bool
	remove(token string) bool
}

type Cache struct {
	ExpiresTime int
	Storage       interface{}
}

type CacheRedis struct {
	CacheMethodInterface
	*Cache
}

func (c *CacheRedis) init() CacheMethodInterface {
	c.Storage =
}

func (c *CacheRedis) store(token string) (string, error) {

}

func (c *CacheRedis) parser(token string) (Claims, error) {

}
func (c *CacheRedis) exists(token string) bool {

}
func (c *CacheRedis) refresh(token string) bool {

}
func (c *CacheRedis) rmove(token string) bool {

}