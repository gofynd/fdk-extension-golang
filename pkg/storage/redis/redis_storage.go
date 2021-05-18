package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

//Storer holds redis storage behaviour
type Storer interface {
	Set(string, interface{}) (string, error)
	Setex(string, interface{}, time.Duration) (string, error)
	Hset(string, string, interface{}) (bool, error)
	Get(string) (string, error)
	Hget(string, string) (string, error)
	Hgetall(string) (map[string]string, error)
	Del(string) (int64, error)
}

//Storage holds redis storage object properties
type Storage struct {
	Client    *redis.Client
	PrefixKey string
}

//New returns new redis storage instance
func New(client *redis.Client, prefixKey string) *Storage {
	return &Storage{
		Client:    client,
		PrefixKey: prefixKey,
	}
}

//Set wrapper for redis set command
func (s *Storage) Set(key string, value interface{}) (string, error) {
	return s.Client.Set(s.PrefixKey+key, value, 0).Result()
}

//Setex wrapper for redis setex command
func (s *Storage) Setex(key string, value interface{}, ttl time.Duration) (string, error) {
	// return s.Client.Set(s.PrefixKey+key, value, ttl).Result()
	return s.Client.Set(fmt.Sprintf("%s%s", s.PrefixKey, key), value, ttl).Result()
	// return s.Client.SetXX(fmt.Sprintf("%s%s", s.PrefixKey, key), value, ttl).Result()
}

//Hset wrapper for redis hset command
func (s *Storage) Hset(key, hashKey string, value interface{}) (bool, error) {
	return s.Client.HSet(s.PrefixKey+key, hashKey, value).Result()
}

//Get wrapper for redis get command
func (s *Storage) Get(key string) (string, error) {
	return s.Client.Get(s.PrefixKey + key).Result()
}

//Hget wrapper for redis hget command
func (s *Storage) Hget(key, hashKey string) (string, error) {
	return s.Client.HGet(s.PrefixKey+key, hashKey).Result()
}

//Hgetall wrapper for redis hgetall command
func (s *Storage) Hgetall(key string) (map[string]string, error) {
	return s.Client.HGetAll(s.PrefixKey + key).Result()
}

//Del wrapper for redis del command
func (s *Storage) Del(key string) (int64, error) {
	return s.Client.Del(s.PrefixKey + key).Result()
}
