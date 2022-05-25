package cache

type Cache interface {
	Type() string
	Connect()
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
}
