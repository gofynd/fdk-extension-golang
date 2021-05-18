# fdk-extension-golang
FDK Extension Helper Library


## Getting Started

Get started with the Golang FDK Extension Helper Library

### Usage

```
import "https://github.com/gofynd/fdk-extension-golang"
```

### Sample Usage 

```golang

  func main() {

	r := gin.Default()
	uri, err := url.Parse("redis://localhost:6379")
	if err != nil {
		log.Fatalf("url parsed failed : %s", err.Error())
	}
	redisOpt, err := redis.ParseURL(uri.String())
    redisClient := redis.NewClient(redisOpt)
    
     //Install callback
     func Install(contextKeys map[string]interface{}) string {
         return ""
     }

    //Auth callback
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


    //Uninstall callback
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

	fdk, err := pkg.SetupFDK(&pkg.FDKInput{
		APIKey:      "YOUR_API_KEY",
		APISecret:   "YOUR_API_SECRET",
		BaseURL:     "YOUR_EXTENSION_HOST_URK",
		Scopes:      []string{""},
		ExtCallback: models.ExtCallback{Auth: Auth, Install: Install, Uninstall: Uninstall},
		Storage:     storage.NewRedisStorage(redisClient, ""),
		AccessMode:  "ACCESS_MODE_VALUE",
		Cluster:     "CLUSTER_URL",
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
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		c.JSON(200, data)
		return
	})
	r.Run(":5070")
}
```


### Documentation


### SetupFDK(*pkg.FDKInput)
SetupFDK will return the fdk instance which enables to quickstart the extension development.

Provide values such as `api_key`, `api_secret`, `base_url`, `scopes`, `callbacks`, `storage`, `access_mode`, `cluster`

Use FDK instance to start the extension server 


### Define extension callback functions in below singature

### 1. Install(map[string]interface{}) string 
*Install* is the extension callback function invoked after user starts the extension installation.

### 2. Auth(map[string]interface{}) string 
*Auth* is the extension callback function invoked after user authorizes the extension service.
It will have `fdk-session` and `extension` present in the context keys which is passed as function argument.

### 3. Uninstall(map[string]interface{}) string 
*Uninstall* is the extension callback function invoked after user uninstall the extension service from the platform.
It will have `platform-client` (only if `access_mode` = `offline`) and `extension` present in the context keys which is passed as function argument.


*FDK instance* also provides `ApplicationProxyRoutes` and `APIRoutes` which registers middlewares to respective route groups.

`ApplicationProxyRoutes` sets various values in the context namely `application`, `application-config` and `application-client`.
`APIRoutes` sets values in the context namely `platform-client` and `extension`.

