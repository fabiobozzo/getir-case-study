package kv

type Storage interface {
	Put(key, value string) error
	Get(key string) (string, error)
}
