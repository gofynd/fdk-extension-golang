package er

import "fmt"

//FdkInvalidExtensionJSON ...
type FdkInvalidExtensionJSON struct {
	Message string `json:"string"`
}

//FdkClusterMetaMissingError ...
type FdkClusterMetaMissingError struct {
	Message string `json:"string"`
}

//FdkSessionNotFoundError ...
type FdkSessionNotFoundError struct {
	Message string `json:"string"`
}

//FdkInvalidOAuthError ...
type FdkInvalidOAuthError struct {
	Message string `json:"string"`
}

//NewFdkInvalidExtensionJSON ...
func NewFdkInvalidExtensionJSON(message string) *FdkInvalidExtensionJSON {
	return &FdkInvalidExtensionJSON{message}
}

//NewFdkSessionNotFoundError ...
func NewFdkSessionNotFoundError(message string) *FdkSessionNotFoundError {
	return &FdkSessionNotFoundError{message}
}

//NewFdkInvalidOAuthError ...
func NewFdkInvalidOAuthError(message string) *FdkInvalidOAuthError {
	return &FdkInvalidOAuthError{message}
}

//NewFdkClusterMetaMissingError ...
func NewFdkClusterMetaMissingError(message string) *FdkClusterMetaMissingError {
	return &FdkClusterMetaMissingError{message}
}

func (f *FdkInvalidExtensionJSON) Error() string {
	return fmt.Sprintf("%s", f.Message)
}

func (f *FdkClusterMetaMissingError) Error() string {
	return fmt.Sprintf("%s", f.Message)
}

func (f *FdkSessionNotFoundError) Error() string {
	return fmt.Sprintf("%s", f.Message)
}

func (f *FdkInvalidOAuthError) Error() string {
	return fmt.Sprintf("%s", f.Message)
}
