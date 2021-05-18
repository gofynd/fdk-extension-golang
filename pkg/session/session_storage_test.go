package session

import (
	"fdk-extension-golang/pkg/extension"
	"fdk-extension-golang/pkg/mocks"
	"reflect"
	"testing"

	goredis "github.com/go-redis/redis"
)

var redisMock = mocks.NewMockRedisStorage(&goredis.Client{}, "")

func TestNewSessionStorage(t *testing.T) {
	type args struct {
		ext *extension.Extension
	}
	tests := []struct {
		name string
		args args
		want *Storage
	}{
		{name: "TestNewSessionStorage", args: args{&extension.Extension{Storage: redisMock}}, want: &Storage{&extension.Extension{Storage: redisMock}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSessionStorage(tt.args.ext); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSessionStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_SaveSession(t *testing.T) {
	type fields struct {
		extension *extension.Extension
	}
	type args struct {
		session *Session
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "TestStorage_SaveSession", fields: fields{&extension.Extension{Storage: redisMock}}, args: args{&Session{}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				extension: tt.fields.extension,
			}
			if err := s.SaveSession(tt.args.session); (err != nil) != tt.wantErr {
				t.Errorf("Storage.SaveSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_GetSession(t *testing.T) {
	type fields struct {
		extension *extension.Extension
	}
	type args struct {
		sessionID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Session
		wantErr bool
	}{
		{name: "TestStorage_GetSession", fields: fields{&extension.Extension{Storage: redisMock}}, args: args{""}, want: &Session{Scope: []string{""}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				extension: tt.fields.extension,
			}
			got, err := s.GetSession(tt.args.sessionID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.GetSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.GetSession() = %+v \n want %+v", got, tt.want)
			}
		})
	}
}

func TestStorage_DeleteSession(t *testing.T) {
	type fields struct {
		extension *extension.Extension
	}
	type args struct {
		sessionID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "TestStorage_DeleteSession", fields: fields{&extension.Extension{Storage: redisMock}}, args: args{""}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				extension: tt.fields.extension,
			}
			if err := s.DeleteSession(tt.args.sessionID); (err != nil) != tt.wantErr {
				t.Errorf("Storage.DeleteSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
