package storage

import (
	"testing"

	goredis "github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

func TestNewRedisStorage(t *testing.T) {
	type args struct {
		client    *goredis.Client
		prefixKey string
	}
	tests := []struct {
		name string
		args args
		want *Storage
	}{
		{name: "TestNewRedisStorage", args: args{&goredis.Client{}, ""}, want: &Storage{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRedisStorage(tt.args.client, tt.args.prefixKey)
			assert.NotNil(t, got)
		})
	}
}

func TestNewMemoryStorage(t *testing.T) {
	type args struct {
		prefixKey string
	}
	tests := []struct {
		name string
		args args
		want *Storage
	}{
		{name: "TestNewMemoryStorage", args: args{""}, want: &Storage{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMemoryStorage(tt.args.prefixKey)
			assert.NotNil(t, got)
		})
	}
}
