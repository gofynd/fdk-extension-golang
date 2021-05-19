package routes

import (
	"fdk-extension-golang/pkg/extension"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupProxyRoutes(t *testing.T) {
	type args struct {
		ext *extension.Extension
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "TestSetupProxyRoutes", args: args{&extension.Extension{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SetupProxyRoutes(tt.args.ext)
			assert.NotNil(t, got)
			assert.NotNil(t, got1)
		})
	}
}
