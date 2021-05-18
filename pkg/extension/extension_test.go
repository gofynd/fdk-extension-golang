package extension

import (
	"fdk-extension-golang/pkg/mocks"
	"fdk-extension-golang/pkg/models"
	"fdk-extension-golang/pkg/storage"
	"reflect"
	"testing"

	goredis "github.com/go-redis/redis"
	"github.com/gofynd/fdk-client-golang/sdk/platform"
	"github.com/stretchr/testify/assert"
)

var (
	redisMock   = mocks.NewMockRedisStorage(&goredis.Client{}, "")
	extCallback = mocks.GetExtCallback()
)

func TestNew(t *testing.T) {
	type args struct {
		apiKey       string
		apiSecret    string
		baseURL      string
		accessMode   string
		cluster      string
		storage      *storage.Storage
		scopes       []string
		extCallbacks models.ExtCallback
	}

	tests := []struct {
		name    string
		args    args
		want    *Extension
		wantErr bool
	}{
		{
			name:    "TestNew",
			args:    args{"000001", "tetsjskdjalsjdl", "https://light-hound-71.loca.lt", "online", "https://nice-turtle-9.loca.lt", redisMock, []string{"company/profiles"}, extCallback},
			want:    &Extension{"000001", "tetsjskdjalsjdl", redisMock, "https://light-hound-71.loca.lt", extCallback, "online", "https://nice-turtle-9.loca.lt", []string{"company/profiles"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.apiKey, tt.args.apiSecret, tt.args.baseURL, tt.args.accessMode, tt.args.cluster, tt.args.storage, tt.args.scopes, tt.args.extCallbacks)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.BaseURL, tt.want.BaseURL) {
				t.Errorf("New() = %v\n want %v", got, tt.want)
			}
		})
	}
}

func TestExtension_GetAuthCallback(t *testing.T) {
	type fields struct {
		APIKey      string
		APISecret   string
		Storage     *storage.Storage
		BaseURL     string
		ExtCallback models.ExtCallback
		AccessMode  string
		Cluster     string
		Scopes      []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "TestExtension_GetAuthCallback",
			fields: fields{BaseURL: "https://light-hound-71.loca.lt"},
			want:   "https://light-hound-71.loca.lt/fp/auth",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Extension{
				APIKey:      tt.fields.APIKey,
				APISecret:   tt.fields.APISecret,
				Storage:     tt.fields.Storage,
				BaseURL:     tt.fields.BaseURL,
				ExtCallback: tt.fields.ExtCallback,
				AccessMode:  tt.fields.AccessMode,
				Cluster:     tt.fields.Cluster,
				Scopes:      tt.fields.Scopes,
			}
			if got := e.GetAuthCallback(); got != tt.want {
				t.Errorf("Extension.GetAuthCallback() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtension_IsOnlineAccessMode(t *testing.T) {
	type fields struct {
		APIKey      string
		APISecret   string
		Storage     *storage.Storage
		BaseURL     string
		ExtCallback models.ExtCallback
		AccessMode  string
		Cluster     string
		Scopes      []string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "TestExtension_IsOnlineAccessMode",
			fields: fields{AccessMode: "online"},
			want:   true,
		},
		{
			name:   "TestExtension_IsOnlineAccessMode",
			fields: fields{AccessMode: "offline"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Extension{
				APIKey:      tt.fields.APIKey,
				APISecret:   tt.fields.APISecret,
				Storage:     tt.fields.Storage,
				BaseURL:     tt.fields.BaseURL,
				ExtCallback: tt.fields.ExtCallback,
				AccessMode:  tt.fields.AccessMode,
				Cluster:     tt.fields.Cluster,
				Scopes:      tt.fields.Scopes,
			}
			if got := e.IsOnlineAccessMode(); got != tt.want {
				t.Errorf("Extension.IsOnlineAccessMode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtension_GetPlatformClient(t *testing.T) {
	type fields struct {
		APIKey      string
		APISecret   string
		Storage     *storage.Storage
		BaseURL     string
		ExtCallback models.ExtCallback
		AccessMode  string
		Cluster     string
		Scopes      []string
	}
	type args struct {
		companyID string
		session   platform.RawToken
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *platform.PlatformClient
	}{
		{
			name:   "TestExtension_GetPlatformClient",
			fields: fields{"000001", "tetsjskdjalsjdl", redisMock, "https://light-hound-71.loca.lt", extCallback, "online", "https://nice-turtle-9.loca.lt", []string{"company/profiles"}},
			args: args{"1", platform.RawToken{ExpiresIn: 3600,
				AccessToken:  "oa-304e62737b31638ce16545f74da4eb99f2cedc50",
				RefreshToken: "oa-5f9739dc84117844f4c3a1aac8a92d411768df04",
				TokenType:    "bearer",
				CurrentUser:  map[string]interface{}{}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Extension{
				APIKey:      tt.fields.APIKey,
				APISecret:   tt.fields.APISecret,
				Storage:     tt.fields.Storage,
				BaseURL:     tt.fields.BaseURL,
				ExtCallback: tt.fields.ExtCallback,
				AccessMode:  tt.fields.AccessMode,
				Cluster:     tt.fields.Cluster,
				Scopes:      tt.fields.Scopes,
			}
			got := e.GetPlatformClient(tt.args.companyID, tt.args.session)
			assert.NotNil(t, got)
		})
	}
}
