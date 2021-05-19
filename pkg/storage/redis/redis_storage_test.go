package redis

import (
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/go-redis/redis"
)

// func TestNew(t *testing.T) {
// 	uri, err := url.Parse("redis://localhost:6379")
// 	// uri.RawQuery = url.Values{}.Encode()
// 	// fmt.Println("uri.RawQuery", uri.RawQuery)
// 	// fmt.Println("uri.String", uri.String())
// 	redisOpt, err := redis.ParseURL(uri.String())
// 	redisClient := redis.NewClient(redisOpt)
// 	_, err = redisClient.Ping().Result()
// 	log.Println("redis connected successfully", redisClient)
// 	redisStorage := New(redisClient, "test-prefix")
// 	log.Printf("redisStorage = %+v\n\n", redisStorage)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, redisClient)
// 	assert.NotNil(t, redisStorage)
// 	assert.Equal(t, redisStorage.PrefixKey, "test-prefix")
// }

func GetRedisClient() *redis.Client {
	uri, _ := url.Parse("redis://localhost:6379")
	redisOpt, _ := redis.ParseURL(uri.String())
	return redis.NewClient(redisOpt)
}
func TestNew(t *testing.T) {
	redisClient := GetRedisClient()
	type args struct {
		client    *redis.Client
		prefixKey string
	}
	tests := []struct {
		name string
		args args
		want *Storage
	}{
		{"Initiliasing redis instance", args{redisClient, "test:"}, &Storage{redisClient, "test:"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.client, tt.args.prefixKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_Set(t *testing.T) {
	redisClient := GetRedisClient()
	type fields struct {
		Client    *redis.Client
		PrefixKey string
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{"Redis Set", fields{redisClient, "test:"}, args{"set", `{"id":"","company_id":"","state":"","scope":[""],"expires":"","expires_in":0,"access_mode":"","access_token":"","current_user":null,"refresh_token":"","is_new":false}`}, "OK", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				Client:    tt.fields.Client,
				PrefixKey: tt.fields.PrefixKey,
			}
			got, err := s.Set(tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Set() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_Setex(t *testing.T) {
	redisClient := GetRedisClient()
	type fields struct {
		Client    *redis.Client
		PrefixKey string
	}
	type args struct {
		key   string
		value interface{}
		ttl   time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{name: "TestStorage_Setex", fields: fields{redisClient, "test:"}, args: args{"setex", `{"id":"","company_id":"","state":"","scope":[""],"expires":"","expires_in":0,"access_mode":"","access_token":"","current_user":null,"refresh_token":"","is_new":false}`, time.Duration(time.Second * 3600)}, want: "OK", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				Client:    tt.fields.Client,
				PrefixKey: tt.fields.PrefixKey,
			}
			got, err := s.Setex(tt.args.key, tt.args.value, tt.args.ttl)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Setex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.Setex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_Hset(t *testing.T) {
	redisClient := GetRedisClient()

	type fields struct {
		Client    *redis.Client
		PrefixKey string
	}
	type args struct {
		key     string
		hashKey string
		value   interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{name: "TestStorage_Hset", fields: fields{redisClient, "test:"}, args: args{"hset", "id", "1"}, want: true, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				Client:    tt.fields.Client,
				PrefixKey: tt.fields.PrefixKey,
			}
			got, err := s.Hset(tt.args.key, tt.args.hashKey, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Hset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.Hset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_Get(t *testing.T) {
	redisClient := GetRedisClient()

	type fields struct {
		Client    *redis.Client
		PrefixKey string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{name: "TestStorage_Get", fields: fields{redisClient, "test:"}, args: args{"set"}, want: `{"id":"","company_id":"","state":"","scope":[""],"expires":"","expires_in":0,"access_mode":"","access_token":"","current_user":null,"refresh_token":"","is_new":false}`, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				Client:    tt.fields.Client,
				PrefixKey: tt.fields.PrefixKey,
			}
			got, err := s.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_Hget(t *testing.T) {
	redisClient := GetRedisClient()

	type fields struct {
		Client    *redis.Client
		PrefixKey string
	}
	type args struct {
		key     string
		hashKey string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{name: "TestStorage_Hget", fields: fields{redisClient, "test:"}, args: args{"hset", "id"}, want: "1", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				Client:    tt.fields.Client,
				PrefixKey: tt.fields.PrefixKey,
			}
			got, err := s.Hget(tt.args.key, tt.args.hashKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Hget() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.Hget() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_Hgetall(t *testing.T) {
	redisClient := GetRedisClient()

	type fields struct {
		Client    *redis.Client
		PrefixKey string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]string
		wantErr bool
	}{
		{name: "TestStorage_Hgetall", fields: fields{redisClient, "test:"}, args: args{"hset"}, want: map[string]string{"id": "1"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				Client:    tt.fields.Client,
				PrefixKey: tt.fields.PrefixKey,
			}
			got, err := s.Hgetall(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Hgetall() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.Hgetall() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_Del(t *testing.T) {
	redisClient := GetRedisClient()

	type fields struct {
		Client    *redis.Client
		PrefixKey string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{name: "TestStorage_Del", fields: fields{redisClient, "test:"}, args: args{"set"}, want: 1, wantErr: false},
		{name: "TestStorage_Del", fields: fields{redisClient, "test:"}, args: args{"setex"}, want: 1, wantErr: false},
		{name: "TestStorage_Del", fields: fields{redisClient, "test:"}, args: args{"hset"}, want: 1, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				Client:    tt.fields.Client,
				PrefixKey: tt.fields.PrefixKey,
			}
			got, err := s.Del(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Del() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.Del() = %v, want %v", got, tt.want)
			}
		})
	}
}
