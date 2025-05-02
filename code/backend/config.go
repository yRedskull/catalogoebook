package main

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
	"time"

	br "github.com/anargu/gin-brotli"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
)

func noCache() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "no-store, no-cache, must-revalidate, private")
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "0")

		c.Next()
	}
}

func SecurityHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Next()
	}
}

func CompressionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ae := c.GetHeader("Accept-Encoding")

		if strings.Contains(ae, "br") {
			br.Brotli(br.DefaultCompression)(c)
		} else if strings.Contains(ae, "gzip") {
			gzip.Gzip(gzip.DefaultCompression)(c)
		}

		c.Next()
	}
}

func ConfigServer() (*gin.Engine, error) {
	r := gin.New()

	r.Use(cors.Default())
	r.Use(SecurityHeaders())
	r.Use(noCache())
	r.SetTrustedProxies([]string{"127.0.0.1"})

	r.Use(CompressionMiddleware())

	// MODE DEBUG - NOT PRODUCTION!!!
	gin.SetMode(gin.DebugMode)
	r.Use(gin.Logger())

	/* gin.SetMode(gin.ReleaseMode) */

	r.Use(gin.Recovery())

	r.SetFuncMap(template.FuncMap{})
	r.LoadHTMLGlob("../frontend/templates/*.html")

	r.Static("/static", "../frontend/static")

	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
        Rate:  1 * time.Second, 
        Limit: 6,               
    })

    r.Use(ratelimit.RateLimiter(store, &ratelimit.Options{
        KeyFunc: func(c *gin.Context) string {
            return c.ClientIP()
        },
        ErrorHandler: func(c *gin.Context, info ratelimit.Info) {
            c.Header("X-Rate-Limit-Limit", fmt.Sprintf("%d", info.Limit))
            c.Header("X-Rate-Limit-Reset", fmt.Sprintf("%d", info.ResetTime.Unix()))
            c.HTML(http.StatusTooManyRequests, "429.html", nil)
        },
    }))

	return r, nil
}
