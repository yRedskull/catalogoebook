package main

import (
	"html/template"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)



func SetupRoute() (*gin.Engine, error) {
	r := gin.Default()

	r.Use(cors.Default())
	r.Use(noCache())
	r.SetTrustedProxies([]string{"127.0.0.1"})

	r.SetFuncMap(template.FuncMap{})
	r.LoadHTMLGlob("templates/*.html")

	r.Static("/static", "./static")

	return r, nil
}
