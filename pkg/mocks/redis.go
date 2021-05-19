package mocks

import (
	"time"

	"github.com/go-redis/redis"
)

//RedisMock holds redis storage object properties
type RedisMock struct {
	Client    *redis.Client
	PrefixKey string
}

//NewRedisMock returns new redis storage instance
func NewRedisMock(client *redis.Client, prefixKey string) *RedisMock {
	return &RedisMock{
		Client:    client,
		PrefixKey: prefixKey,
	}
}

//Set wrapper for redis set command
func (r *RedisMock) Set(key string, value interface{}) (string, error) {
	return "OK", nil
}

//Setex wrapper for redis setex command
func (r *RedisMock) Setex(key string, value interface{}, ttl time.Duration) (string, error) {
	return "OK", nil
}

//Hset wrapper for redis hset command
func (r *RedisMock) Hset(key, hashKey string, value interface{}) (bool, error) {
	return true, nil
}

//Get wrapper for redis get command
func (r *RedisMock) Get(key string) (string, error) {
	return `{"id":"","company_id":"","state":"","scope":[""],"expires":"","expires_in":0,"access_mode":"","access_token":"","current_user":null,"refresh_token":"","is_new":false}`, nil
}

//Hget wrapper for redis hget command
func (r *RedisMock) Hget(key, hashKey string) (string, error) {
	return "hget-test", nil
}

//Hgetall wrapper for redis hgetall command
func (r *RedisMock) Hgetall(key string) (map[string]string, error) {
	return map[string]string{"id": "1"}, nil
}

//Del wrapper for redis del command
func (r *RedisMock) Del(key string) (int64, error) {
	return 1, nil
}
