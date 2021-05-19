package mocks

import (
	"fdk-extension-golang/pkg/storage"

	goredis "github.com/go-redis/redis"
)

//NewMockRedisStorage return redis storage instance
func NewMockRedisStorage(client *goredis.Client, prefixKey string) *storage.Storage {
	return &storage.Storage{
		RedisStorer: NewRedisMock(client, prefixKey),
	}
}

//NewMockMemoryStorage return memory storage instance
func NewMockMemoryStorage(prefixKey string) *storage.Storage {
	return &storage.Storage{
		MemoryStorer: NewMemoryMock(prefixKey),
	}
}
