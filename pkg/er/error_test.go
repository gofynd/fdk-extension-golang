package er

import (
	"reflect"
	"testing"
)

func TestNewFdkInvalidExtensionJSON(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want *FdkInvalidExtensionJSON
	}{
		{name: "TestNewFdkInvalidExtensionJSON", args: args{"Invalid extension"}, want: &FdkInvalidExtensionJSON{"Invalid extension"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFdkInvalidExtensionJSON(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFdkInvalidExtensionJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFdkSessionNotFoundError(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want *FdkSessionNotFoundError
	}{
		{name: "TestNewFdkSessionNotFoundError", args: args{"Session not found"}, want: &FdkSessionNotFoundError{"Session not found"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFdkSessionNotFoundError(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFdkSessionNotFoundError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFdkInvalidOAuthError(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want *FdkInvalidOAuthError
	}{
		{name: "TestNewFdkInvalidOAuthError", args: args{"Invalid OAuth client"}, want: &FdkInvalidOAuthError{"Invalid OAuth client"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFdkInvalidOAuthError(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFdkInvalidOAuthError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFdkClusterMetaMissingError(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want *FdkClusterMetaMissingError
	}{
		{name: "TestNewFdkClusterMetaMissingError", args: args{"Missing cluster meta"}, want: &FdkClusterMetaMissingError{"Missing cluster meta"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFdkClusterMetaMissingError(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFdkClusterMetaMissingError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFdkInvalidExtensionJSON_Error(t *testing.T) {
	type fields struct {
		Message string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "TestFdkInvalidExtensionJSON_Error", fields: fields{"Invalid extension"}, want: "Invalid extension"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FdkInvalidExtensionJSON{
				Message: tt.fields.Message,
			}
			if got := f.Error(); got != tt.want {
				t.Errorf("FdkInvalidExtensionJSON.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFdkClusterMetaMissingError_Error(t *testing.T) {
	type fields struct {
		Message string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "TestFdkClusterMetaMissingError_Error", fields: fields{"missing cluster meta"}, want: "missing cluster meta"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FdkClusterMetaMissingError{
				Message: tt.fields.Message,
			}
			if got := f.Error(); got != tt.want {
				t.Errorf("FdkClusterMetaMissingError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFdkSessionNotFoundError_Error(t *testing.T) {
	type fields struct {
		Message string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "TestFdkSessionNotFoundError_Error", fields: fields{"session not found"}, want: "session not found"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FdkSessionNotFoundError{
				Message: tt.fields.Message,
			}
			if got := f.Error(); got != tt.want {
				t.Errorf("FdkSessionNotFoundError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFdkInvalidOAuthError_Error(t *testing.T) {
	type fields struct {
		Message string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "TestFdkInvalidOAuthError_Error", fields: fields{"invalid oauth"}, want: "invalid oauth"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FdkInvalidOAuthError{
				Message: tt.fields.Message,
			}
			if got := f.Error(); got != tt.want {
				t.Errorf("FdkInvalidOAuthError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
