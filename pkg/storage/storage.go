package storage

import (
	"fdk-extension-golang/pkg/storage/memory"
	"fdk-extension-golang/pkg/storage/redis"

	goredis "github.com/go-redis/redis"
)

//Storage ...
type Storage struct {
	RedisStorer  redis.Storer  `json:"redis_storer"`
	MemoryStorer memory.Storer `json:"memory_storer"`
}

//NewRedisStorage return redis storage instance
func NewRedisStorage(client *goredis.Client, prefixKey string) *Storage {
	return &Storage{
		RedisStorer: redis.New(client, prefixKey),
	}
}

//NewMemoryStorage return memory storage instance
func NewMemoryStorage(prefixKey string) *Storage {
	return &Storage{
		MemoryStorer: memory.New(prefixKey),
	}
}
