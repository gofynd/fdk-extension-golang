package routes

import (
	"fdk-extension-golang/pkg/extension"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var testRoute = &gin.Engine{}

func TestSetupRoutes(t *testing.T) {
	type args struct {
		ext *extension.Extension
	}
	tests := []struct {
		name string
		args args
		want *gin.Engine
	}{
		{name: "TestSetupRoutes", args: args{&extension.Extension{}}, want: &gin.Engine{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SetupRoutes(tt.args.ext)
			assert.NotNil(t, got)
			testRoute = got
		})
	}

}
