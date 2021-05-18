package pkg

import (
	"fdk-extension-golang/pkg/mocks"
	"reflect"
	"testing"

	goredis "github.com/go-redis/redis"
)

var (
	redisMock   = mocks.NewMockRedisStorage(&goredis.Client{}, "")
	extCallback = mocks.GetExtCallback()
)

func TestSetupFDK(t *testing.T) {
	type args struct {
		fdkInput *FDKInput
	}
	tests := []struct {
		name    string
		args    args
		want    *FDK
		wantErr bool
	}{
		{name: "TestSetupFDK", args: args{&FDKInput{"000001", "tetsjskdjalsjdl", "https://light-hound-71.loca.lt", []string{"company/profiles"}, extCallback, redisMock, "online", "https://nice-turtle-9.loca.lt"}}, want: &FDK{}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SetupFDK(tt.args.fdkInput)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetupFDK() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Extension.BaseURL, "https://light-hound-71.loca.lt") {
				t.Errorf("SetupFDK() = %+v \n want %+v", got.Extension.BaseURL, "https://light-hound-71.loca.lt")
			}
		})
	}
}
