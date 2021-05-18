package routes

import (
	"fdk-extension-golang/pkg/er"
	"fdk-extension-golang/pkg/extension"
	"fdk-extension-golang/pkg/middlewares"
	"fdk-extension-golang/pkg/models"
	"fdk-extension-golang/pkg/session"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/gofynd/fdk-client-golang/sdk/platform"
	"github.com/google/uuid"
)

//SetupRoutes ...
func SetupRoutes(ext *extension.Extension) *gin.Engine {

	//default http routing
	// mux := http.NewServeMux()
	// mux.HandleFunc()

	//gorilla routing
	// r := mux.NewRouter()
	// s := r.PathPrefix("/fp").Subrouter()
	// s.HandleFunc("/setup", SetupHandler).Methods("")

	//gin routing
	router := gin.Default()
	fp := router.Group("/fp")

	//setup call is deperected
	fp.POST("/setup", func(c *gin.Context) {
		c.JSON(200, struct {
			Success bool
			Message string
		}{Success: true, Message: "This is deprected call"})
	})

	fp.GET("/install", func(c *gin.Context) {
		// ?company_id=1&client_id=123313112122
		companyID := c.Query("company_id")
		if companyID == "" {
			c.Error(fmt.Errorf("Invalid company id"))
			return
		}
		platformConfig := ext.GetPlatformConfig(companyID)
		platformConfig.SetOAuthClient()
		FDKSession := &session.Session{}
		sessionStorage := session.NewSessionStorage(ext)
		var err error

		if ext.IsOnlineAccessMode() {
			sid, err := session.GenerateSessionID(true, models.Option{})
			if err != nil {
				c.Error(err)
				return
			}
			FDKSession = session.New(sid, true)
		} else {
			sid, err := session.GenerateSessionID(false, models.Option{
				CompanyID: companyID,
				Cluster:   ext.Cluster,
			})
			if err != nil {
				c.Error(err)
				return
			}
			FDKSession, err = sessionStorage.GetSession(sid)
			if err != nil && err != redis.Nil {
				c.Error(err)
				return
			}
			if FDKSession.ID == "" || err == redis.Nil {
				FDKSession = session.New(sid, true)
			}
		}

		sessionExpires := time.Now().UTC().Add(time.Minute * time.Duration(15)).Format(time.RFC3339)
		if FDKSession.IsNew {
			FDKSession.CompanyID = companyID
			FDKSession.Scope = ext.Scopes
			FDKSession.Expires = sessionExpires
			FDKSession.AccessMode = ext.AccessMode
		} else {
			// 	if(session.expires) {
			// 		session.expires = new Date(session.expires);
			// 	}
		}

		//TODO set signed cookie
		c.SetSameSite(http.SameSiteNoneMode)
		c.SetCookie(models.SESSIONCOOKIENAME, FDKSession.ID, 900, "", "", true, true)
		// let redirectUrl;

		// if(process.env.NODE_ENV === "production") {
		// 	if(!session.access_token) {
		// 		session.state = uuidv4();
		// 		// start authorization flow
		// 		redirectUrl = platformConfig.oauthClient.startAuthorization({
		// 			scope: session.scope,
		// 			redirectUri: ext.getAuthCallback(),
		// 			state: session.state,
		// 			access_mode: ext.access_mode
		// 		});
		// 	} else {
		// 		redirectUrl = await ext.callbacks.install(req);
		// 	}
		// } else {
		// 	session.state = uuidv4();
		// 	// start authorization flow
		// 	redirectUrl = platformConfig.oauthClient.startAuthorization({
		// 		scope: session.scope,
		// 		redirectUri: ext.getAuthCallback(),
		// 		state: session.state,
		// 		access_mode: ext.access_mode
		// 	});
		// }

		//else case code
		FDKSession.State = uuid.New().String()
		// start authorization flow
		redirectURL, err := platformConfig.OAuthClient.StartAuthorization(platform.Option{
			Scope:       FDKSession.Scope,
			RedirectURI: ext.GetAuthCallback(),
			State:       FDKSession.State,
			AccessMode:  ext.AccessMode,
		})
		if err != nil {
			c.Error(err)
			return
		}
		err = sessionStorage.SaveSession(FDKSession)
		if err != nil {
			c.Error(err)
			return
		}
		c.Redirect(http.StatusFound, redirectURL)
	})

	fp.GET("/auth", middlewares.SessionMiddleware(false, session.NewSessionStorage(ext)), func(c *gin.Context) {
		// ?code=ddjfhdsjfsfh&client_id=jsfnsajfhkasf&company_id=1&state=jashoh
		var fdkSession *session.Session
		sessionVal, ok := c.Get("fdk-session")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": er.NewFdkSessionNotFoundError("fdk session not found")})
			return
		}
		if val, ok := sessionVal.(*session.Session); ok {
			fdkSession = val
		}

		if fdkSession.ID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": er.NewFdkSessionNotFoundError("Can not complete oauth process as session not found")})
			return
		}
		if fdkSession.State != c.Query("state") {
			c.JSON(http.StatusBadRequest, gin.H{"error": er.NewFdkInvalidOAuthError("Invalid oauth call")})
			return
		}
		platformConfig := ext.GetPlatformConfig(fdkSession.CompanyID)
		platformConfig.SetOAuthClient()
		err := platformConfig.OAuthClient.VerifyCallback(platform.Query{
			Code: c.Query("code"),
		})
		if err != nil {
			c.Error(err)
			return
		}
		token := platformConfig.OAuthClient.RawToken
		sessionExpires := time.Now().Add(time.Second * time.Duration(token.ExpiresIn))

		if ext.IsOnlineAccessMode() {
			fdkSession.Expires = sessionExpires.Format(time.RFC3339)
		} else {
			//for offline access mode , expiry is not set
			fdkSession.Expires = ""
		}

		fdkSession.AccessToken = token.AccessToken
		fdkSession.ExpiresIn = token.ExpiresIn
		fdkSession.CurrentUser = token.CurrentUser
		fdkSession.RefreshToken = token.RefreshToken

		err = session.NewSessionStorage(ext).SaveSession(fdkSession)
		if err != nil {
			c.Error(err)
			return
		}

		//TODO set signed cookie
		c.SetSameSite(http.SameSiteNoneMode)
		c.SetCookie(models.SESSIONCOOKIENAME, fdkSession.ID, token.ExpiresIn, "", "", true, true)

		//instead of setting to req body, context is used
		c.Set("fdk-session", fdkSession)
		c.Set("extension", ext)

		redirectURL := ext.ExtCallback.Auth(c.Keys)
		c.Redirect(http.StatusFound, redirectURL)

	})

	fp.POST("/uninstall", func(c *gin.Context) {
		type req struct {
			ClientID  string `json:"client_id"`
			CompanyID string `json:"company_id"`
		}
		var reqBody req
		sessionStorage := session.NewSessionStorage(ext)
		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if !ext.IsOnlineAccessMode() {
			sid, err := session.GenerateSessionID(false, models.Option{
				CompanyID: reqBody.CompanyID,
				Cluster:   ext.Cluster,
			})
			if err != nil {
				c.Error(err)
				return
			}
			fdkSession, err := sessionStorage.GetSession(sid)
			if err != nil {
				c.Error(err)
				return
			}
			client := ext.GetPlatformClient(reqBody.CompanyID, platform.RawToken{
				ExpiresIn:    fdkSession.ExpiresIn,
				AccessToken:  fdkSession.AccessToken,
				RefreshToken: fdkSession.RefreshToken,
			})
			c.Set("platform-client", client)
			err = sessionStorage.DeleteSession(sid)
			if err != nil {
				c.Error(err)
				return
			}
		}
		c.Set("extension", ext)
		ext.ExtCallback.Uninstall(c.Keys)
		c.JSON(http.StatusOK, gin.H{"success": true})
	})

	return router
}
