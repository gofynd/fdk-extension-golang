package memory

import "time"

//Storer holds memory storage behaviour
type Storer interface {
	Get(string) interface{}
	Set(string, interface{})
	Setex(string, interface{}, time.Duration)
	Del(string)
	Hget(string, string) interface{}
	Hset(string, string, interface{})
	Hgetall(string) map[string]interface{}
}

//Storage holds memory storage object properties
type Storage struct {
	Data      map[string]interface{}
	PrefixKey string
}

//New returns new memory storage instance
func New(prefixKey string) *Storage {
	return &Storage{
		Data:      make(map[string]interface{}),
		PrefixKey: prefixKey,
	}
}

//Get wrapper to get data by specific key
func (s *Storage) Get(key string) interface{} {
	return s.Data[s.PrefixKey+key]
}

//Set wrapper to set data by specific key
func (s *Storage) Set(key string, value interface{}) {
	s.Data[s.PrefixKey+key] = value
}

//Setex wrapper to set data by specific key with expiry
func (s *Storage) Setex(key string, value interface{}, ttl time.Duration) {
	//TODO: add support for ttl
	s.Data[s.PrefixKey+key] = value
}

//Del wrapper to delete data by specific key
func (s *Storage) Del(key string) {
	delete(s.Data, s.PrefixKey+key)
}

//Hget wrapper to get specific field value from hash set by specific key and hash key
func (s *Storage) Hget(key, hashKey string) interface{} {
	hashMap := s.Data[s.PrefixKey+key]
	if val, ok := hashMap.(map[string]interface{}); ok {
		return val[hashKey]
	}
	return nil
}

//Hset wrapper to set hash set to data by specific hash key
func (s *Storage) Hset(key, hashKey string, value interface{}) {
	hashMap := make(map[string]interface{})
	if val, ok := s.Data[s.PrefixKey+key]; ok {
		if _, ok := val.(map[string]interface{}); ok {
			hashMap = val.(map[string]interface{})
		}
	}
	hashMap[hashKey] = value
	s.Data[s.PrefixKey+key] = hashMap
}

//Hgetall wrapper to get all field's values from hash set by specific hash key
func (s *Storage) Hgetall(key string) map[string]interface{} {
	if val, ok := s.Data[s.PrefixKey+key].(map[string]interface{}); ok {
		return val
	}
	return nil
}
