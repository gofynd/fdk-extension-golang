package pkg

import (
	"fdk-extension-golang/pkg/extension"
	"fdk-extension-golang/pkg/models"
	"fdk-extension-golang/pkg/routes"
	"fdk-extension-golang/pkg/session"
	"fdk-extension-golang/pkg/storage"

	"github.com/gin-gonic/gin"
	"github.com/gofynd/fdk-client-golang/sdk/application"
	"github.com/gofynd/fdk-client-golang/sdk/platform"
)

//FDK holds fdk object properties
type FDK struct {
	Extension              *extension.Extension
	FDKHandler             *gin.Engine
	APIRoutes              *gin.Engine
	ApplicationProxyRoutes *gin.Engine
	GetPlatformClient      func(string) (*platform.PlatformClient, error)
	GetApplicationClient   func(string, string) (*application.Client, error)
}

//FDKInput holds set fdk input properties
type FDKInput struct {
	APIKey      string             `json:"api_key"`
	APISecret   string             `json:"api_secret"`
	BaseURL     string             `json:"base_url"`
	Scopes      []string           `json:"scopes"`
	ExtCallback models.ExtCallback `json:"callbacks"`
	Storage     *storage.Storage   `json:"storage"`
	AccessMode  string             `json:"access_mode"`
	Cluster     string             `json:"cluster"`
}

//SetupFDK returns fdk instance
func SetupFDK(fdkInput *FDKInput) (*FDK, error) {
	ext, err := extension.New(fdkInput.APIKey, fdkInput.APISecret, fdkInput.BaseURL, fdkInput.AccessMode, fdkInput.Cluster, fdkInput.Storage, fdkInput.Scopes, fdkInput.ExtCallback)
	if err != nil {
		return &FDK{}, err
	}
	router := routes.SetupRoutes(ext)
	apiRoutes, applicationProxyRoutes := routes.SetupProxyRoutes(ext)

	getPlatformClient := func(companyID string) (*platform.PlatformClient, error) {
		client := &platform.PlatformClient{}
		if !ext.IsOnlineAccessMode() {
			sid, err := session.GenerateSessionID(false, models.Option{
				CompanyID: companyID,
				Cluster:   ext.Cluster,
			})
			if err != nil {
				return &platform.PlatformClient{}, err
			}
			sessionStorage := session.NewSessionStorage(ext)
			session, err := sessionStorage.GetSession(sid)
			if err != nil {
				return &platform.PlatformClient{}, err
			}
			client = ext.GetPlatformClient(companyID, platform.RawToken{
				ExpiresIn:    session.ExpiresIn,
				AccessToken:  session.AccessToken,
				RefreshToken: session.RefreshToken,
			})
		}
		return client, nil
	}

	getApplicationClient := func(applicationId, applicationToken string) (*application.Client, error) {
		applicationConfig, err := application.NewAppConfig(applicationId, applicationToken, ext.Cluster, &application.Options{})
		if err != nil {
			return &application.Client{}, err
		}
		applicationClient := application.NewAppClient(applicationConfig)
		return applicationClient, nil
	}

	return &FDK{
		FDKHandler:             router,
		Extension:              ext,
		APIRoutes:              apiRoutes,
		ApplicationProxyRoutes: applicationProxyRoutes,
		GetPlatformClient:      getPlatformClient,
		GetApplicationClient:   getApplicationClient,
	}, nil

}
