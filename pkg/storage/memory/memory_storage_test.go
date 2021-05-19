package memory

import (
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	type args struct {
		prefixKey string
	}
	tests := []struct {
		name string
		args args
		want *Storage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.prefixKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_Get(t *testing.T) {
	type fields struct {
		Data      map[string]interface{}
		PrefixKey string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				Data:      tt.fields.Data,
				PrefixKey: tt.fields.PrefixKey,
			}
			if got := s.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_Set(t *testing.T) {
	type fields struct {
		Data      map[string]interface{}
		PrefixKey string
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				Data:      tt.fields.Data,
				PrefixKey: tt.fields.PrefixKey,
			}
			s.Set(tt.args.key, tt.args.value)
		})
	}
}

func TestStorage_Setex(t *testing.T) {
	type fields struct {
		Data      map[string]interface{}
		PrefixKey string
	}
	type args struct {
		key   string
		value interface{}
		ttl   time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				Data:      tt.fields.Data,
				PrefixKey: tt.fields.PrefixKey,
			}
			s.Setex(tt.args.key, tt.args.value, tt.args.ttl)
		})
	}
}

func TestStorage_Del(t *testing.T) {
	type fields struct {
		Data      map[string]interface{}
		PrefixKey string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				Data:      tt.fields.Data,
				PrefixKey: tt.fields.PrefixKey,
			}
			s.Del(tt.args.key)
		})
	}
}

func TestStorage_Hget(t *testing.T) {
	type fields struct {
		Data      map[string]interface{}
		PrefixKey string
	}
	type args struct {
		key     string
		hashKey string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				Data:      tt.fields.Data,
				PrefixKey: tt.fields.PrefixKey,
			}
			if got := s.Hget(tt.args.key, tt.args.hashKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.Hget() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_Hset(t *testing.T) {
	type fields struct {
		Data      map[string]interface{}
		PrefixKey string
	}
	type args struct {
		key     string
		hashKey string
		value   interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				Data:      tt.fields.Data,
				PrefixKey: tt.fields.PrefixKey,
			}
			s.Hset(tt.args.key, tt.args.hashKey, tt.args.value)
		})
	}
}

func TestStorage_Hgetall(t *testing.T) {
	type fields struct {
		Data      map[string]interface{}
		PrefixKey string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				Data:      tt.fields.Data,
				PrefixKey: tt.fields.PrefixKey,
			}
			if got := s.Hgetall(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.Hgetall() = %v, want %v", got, tt.want)
			}
		})
	}
}
