package extension

import (
	"fdk-extension-golang/pkg/er"
	"fdk-extension-golang/pkg/models"
	"fdk-extension-golang/pkg/storage"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/gofynd/fdk-client-golang/sdk/platform"
)

//Extension holds extension object properties
type Extension struct {
	APIKey      string             `json:"api_key"`
	APISecret   string             `json:"api_secret"`
	Storage     *storage.Storage   `json:"-"`
	BaseURL     string             `json:"base_url"`
	ExtCallback models.ExtCallback `json:"-"`
	AccessMode  string             `json:"access_mode"`
	Cluster     string             `json:"cluster"`
	Scopes      []string           `json:"scopes"`
}

//New returns new Extension instance
func New(apiKey, apiSecret, baseURL, accessMode, cluster string, storage *storage.Storage, scopes []string, extCallbacks models.ExtCallback) (*Extension, error) {
	//Validate data
	validate := validator.New()
	var err error
	err = validate.Var(apiKey, "required")
	if err != nil {
		return &Extension{}, er.NewFdkInvalidExtensionJSON(err.Error())
	}
	err = validate.Var(apiSecret, "required")
	if err != nil {
		return &Extension{}, er.NewFdkInvalidExtensionJSON(err.Error())
	}
	err = validate.Var(baseURL, "required,url")
	if err != nil {
		return &Extension{}, er.NewFdkInvalidExtensionJSON(err.Error())
	}
	err = validate.Var(scopes, "gt=0,dive,required")
	if err != nil {
		return &Extension{}, er.NewFdkInvalidExtensionJSON(err.Error())
	}
	err = validate.Var(cluster, "omitempty,required,url")
	if err != nil {
		return &Extension{}, er.NewFdkInvalidExtensionJSON(err.Error())
	}
	if accessMode == "" {
		accessMode = "offline"
	}
	if cluster == "" {
		cluster = "https://api.fynd.com"
	}

	if extCallbacks.Install == nil || extCallbacks.Auth == nil || extCallbacks.Uninstall == nil {
		return &Extension{}, er.NewFdkInvalidExtensionJSON("Missing some of callbacks. Please add all  `auth`, `install` and `uninstall` callbacks.")
	}
	return &Extension{
		APIKey:      apiKey,
		APISecret:   apiSecret,
		Storage:     storage,
		BaseURL:     baseURL,
		ExtCallback: extCallbacks,
		AccessMode:  accessMode,
		Cluster:     cluster,
		Scopes:      scopes,
	}, nil
}

//GetAuthCallback return the auth url
func (e *Extension) GetAuthCallback() string {
	return fmt.Sprintf("%s%s", e.BaseURL, "/fp/auth")
}

//IsOnlineAccessMode checks if access mode is online
func (e *Extension) IsOnlineAccessMode() bool {
	return e.AccessMode == "online"
}

//GetPlatformConfig return the platform config
func (e *Extension) GetPlatformConfig(companyID string) *platform.PlatformConfig {
	//Initialise platform config instance
	return platform.NewPlatformConfig(companyID, e.APIKey, e.APISecret, e.Cluster)
}

//GetPlatformClient return the platform client
func (e *Extension) GetPlatformClient(companyID string, session platform.RawToken) *platform.PlatformClient {
	//Initialise platform config instance
	platformConfig := e.GetPlatformConfig(companyID)

	//Set OAuthClient
	platformConfig.SetOAuthClient()

	//Set access token
	platformConfig.OAuthClient.SetAccessToken(session)

	//return platform client instance
	return platform.NewPlatformClient(platformConfig)
}
