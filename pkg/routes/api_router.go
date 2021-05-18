package routes

import (
	"encoding/json"
	"fdk-extension-golang/pkg/er"
	"fdk-extension-golang/pkg/extension"
	"fdk-extension-golang/pkg/middlewares"
	"fdk-extension-golang/pkg/models"
	"fdk-extension-golang/pkg/session"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofynd/fdk-client-golang/sdk/application"
	"github.com/gofynd/fdk-client-golang/sdk/platform"
)

//SetupProxyRoutes ...
func SetupProxyRoutes(ext *extension.Extension) (*gin.Engine, *gin.Engine) {

	apiRoutes := gin.Default()
	applicationProxyRoutes := gin.Default()

	//middleware that sets platform client to request
	apiRoutes.Use(middlewares.SessionMiddleware(true, session.NewSessionStorage(ext)), func(c *gin.Context) {
		fdkSession := &session.Session{}
		if sessionVal, ok := c.Get("fdk-session"); ok {
			if val, ok := sessionVal.(*session.Session); ok {
				fdkSession = val
			}
		}
		if fdkSession.ID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": er.NewFdkSessionNotFoundError("Can not complete oauth process as session not found")})
			c.Abort()
			return
		}
		client := ext.GetPlatformClient(fdkSession.CompanyID, platform.RawToken{
			ExpiresIn:    fdkSession.ExpiresIn,
			AccessToken:  fdkSession.AccessToken,
			RefreshToken: fdkSession.RefreshToken,
		})
		c.Set("platform-client", client)
		c.Set("extension", ext)
		c.Next()
	})

	//middleware that sets application client to request
	applicationProxyRoutes.Use(func(c *gin.Context) {
		rawUser := c.GetHeader("x-user-data")
		rawApp := c.GetHeader("x-application-data")
		if rawUser != "" {
			user := &models.User{}
			err := json.Unmarshal([]byte(rawUser), user)
			if err != nil {
				c.Error(err)
				c.Abort()
				return
			}
			c.Set("x-user", user)
		}
		if rawApp != "" {
			app := &models.Application{}
			err := json.Unmarshal([]byte(rawApp), app)
			if err != nil {
				c.Error(err)
				c.Abort()
				return
			}
			applicationConfig, err := application.NewAppConfig(app.ID.String(), app.Token, ext.Cluster, &application.Options{})
			if err != nil {
				c.Error(err)
				c.Abort()
				return
			}
			c.Set("application", app)
			c.Set("application-config", applicationConfig)
			c.Set("application-client", application.NewAppClient(applicationConfig))
		}
		c.Next()
	})
	return apiRoutes, applicationProxyRoutes
}
