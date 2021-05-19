package middlewares

import (
	"fdk-extension-golang/pkg/models"
	"fdk-extension-golang/pkg/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

//SessionMiddleware ...
func SessionMiddleware(strict bool, sessionStorage *session.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Read cookie from request
		cookie, err := c.Cookie(models.SESSIONCOOKIENAME)
		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}
		FDKSession, err := sessionStorage.GetSession(cookie)
		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}
		if strict && FDKSession.ID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}
		c.Set("fdk-session", FDKSession)
		c.Next()
	}
}
