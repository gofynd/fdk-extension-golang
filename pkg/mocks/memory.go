package mocks

import "time"

//MemoryMock holds memory storage object properties
type MemoryMock struct {
	Data      map[string]interface{}
	PrefixKey string
}

//NewMemoryMock returns new memory storage instance
func NewMemoryMock(prefixKey string) *MemoryMock {
	return &MemoryMock{
		Data:      make(map[string]interface{}),
		PrefixKey: prefixKey,
	}
}

//Get wrapper to get data by specific key
func (m *MemoryMock) Get(key string) interface{} {
	return m.Data[m.PrefixKey+key]
}

//Set wrapper to set data by specific key
func (m *MemoryMock) Set(key string, value interface{}) {
	m.Data[m.PrefixKey+key] = value
}

//Setex wrapper to set data by specific key with expiry
func (m *MemoryMock) Setex(key string, value interface{}, ttl time.Duration) {
	//TODO: add support for ttl
	m.Data[m.PrefixKey+key] = value
}

//Del wrapper to delete data by specific key
func (m *MemoryMock) Del(key string) {
	delete(m.Data, m.PrefixKey+key)
}

//Hget wrapper to get specific field value from hash set by specific key and hash key
func (m *MemoryMock) Hget(key, hashKey string) interface{} {
	hashMap := m.Data[m.PrefixKey+key]
	if val, ok := hashMap.(map[string]interface{}); ok {
		return val[hashKey]
	}
	return nil
}

//Hset wrapper to set hash set to data by specific hash key
func (m *MemoryMock) Hset(key, hashKey string, value interface{}) {
	hashMap := make(map[string]interface{})
	if val, ok := m.Data[m.PrefixKey+key]; ok {
		if _, ok := val.(map[string]interface{}); ok {
			hashMap = val.(map[string]interface{})
		}
	}
	hashMap[hashKey] = value
	m.Data[m.PrefixKey+key] = hashMap
}

//Hgetall wrapper to get all field's values from hash set by specific hash key
func (m *MemoryMock) Hgetall(key string) map[string]interface{} {
	if val, ok := m.Data[m.PrefixKey+key].(map[string]interface{}); ok {
		return val
	}
	return nil
}
