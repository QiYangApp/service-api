package authorize

type AuthorizedType[T interface{}, R *interface{}] interface {
	Authorizing(params T) R

	Authorized(params)
}
