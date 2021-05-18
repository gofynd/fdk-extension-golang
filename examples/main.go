package main

import (
	"fdk-extension-golang/pkg"
	"fdk-extension-golang/pkg/extension"
	"fdk-extension-golang/pkg/models"
	"fdk-extension-golang/pkg/session"
	"fdk-extension-golang/pkg/storage"
	"log"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/gofynd/fdk-client-golang/sdk/platform"
)

func main() {

	r := gin.Default()
	uri, err := url.Parse("redis://localhost:6379")
	if err != nil {
		log.Fatalf("url parsed failed : %s", err.Error())
	}
	redisOpt, err := redis.ParseURL(uri.String())
	redisClient := redis.NewClient(redisOpt)

	fdk, err := pkg.SetupFDK(&pkg.FDKInput{
		APIKey:      "609bedbfd17c659702c2a331",
		APISecret:   "o2_EfixOXzeuIuk",
		BaseURL:     "https://old-wasp-53.loca.lt",
		Scopes:      []string{"company/profiles"},
		ExtCallback: models.ExtCallback{Auth: Auth, Install: Install, Uninstall: Uninstall},
		Storage:     storage.NewRedisStorage(redisClient, ""),
		AccessMode:  "offline",
		Cluster:     "https://api.fyndx0.de",
	})
	r = fdk.FDKHandler

	r.GET("/_healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"ok": "ok"})
		return
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": true})
		return
	})

	applicationProxyRoutes := r.Group("")
	applicationProxyRoutes.Use(fdk.ApplicationProxyRoutes.Handlers...)
	applicationProxyRoutes.GET("/test/app", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": true})
		return
	})

	APIRoutes := r.Group("")
	APIRoutes.Use(fdk.APIRoutes.Handlers...)
	APIRoutes.GET("/test/platform", func(c *gin.Context) {

		platformClient := &platform.PlatformClient{}
		ext := &extension.Extension{}
		if platformClientContext, ok := c.Get("platform-client"); ok {
			platformClient = platformClientContext.(*platform.PlatformClient)
		}
		if extensionContext, ok := c.Get("extension"); ok {
			ext = extensionContext.(*extension.Extension)
			log.Printf("extension = %+v\n\n", ext)
		}
		data, err := platformClient.Catalog.GetProducts(platform.PlatformGetProductsXQuery{PageNo: 1, PageSize: 5})
		// data, err := platformClient.Lead.GetTickets(platform.PlatformGetTicketsXQuery{PageSize: 5})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		c.JSON(200, data)
		return
	})
	r.Run(":5070")
}

//Auth ...
func Auth(contextKeys map[string]interface{}) string {
	var ok bool
	fdkSession := &session.Session{}
	ext := &extension.Extension{}
	if _, ok := contextKeys["fdk-session"]; !ok {
		return ""
	}
	if fdkSession, ok = contextKeys["fdk-session"].(*session.Session); !ok {
		return ""
	}
	if _, ok := contextKeys["extension"]; !ok {
		return ""
	}
	if ext, ok = contextKeys["extension"].(*extension.Extension); !ok {
		return ""
	}
	log.Printf("fdkSession = %+v\n\n", fdkSession)
	return ext.BaseURL
}

//Install ...
func Install(contextKeys map[string]interface{}) string {
	return ""
}

//Uninstall ...
func Uninstall(contextKeys map[string]interface{}) string {

	platformClient := &platform.PlatformClient{}
	ext := &extension.Extension{}
	if platformClientContext, ok := contextKeys["platform-client"]; ok {
		platformClient = platformClientContext.(*platform.PlatformClient)
		log.Printf("platformClient = %+v\n\n", platformClient)
	}
	if extensionContext, ok := contextKeys["extension"]; ok {
		ext = extensionContext.(*extension.Extension)
		return ext.BaseURL
	}
	return ""
}
