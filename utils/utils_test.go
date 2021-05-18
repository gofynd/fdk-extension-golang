package utils

import (
	"fmt"
	"testing"
)

func TestGetHash(t *testing.T) {
	type args struct {
		dataString string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "Testing GetHash()", args: args{fmt.Sprintf("%s:%s", "http:localhost:9000", "1")}, want: "c4bbe93b605f4a2756ee571c31337d558475b850dca52043eaf55ddba4c208b6", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetHash(tt.args.dataString)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
