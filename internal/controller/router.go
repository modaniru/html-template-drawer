package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/modaniru/html-template-drawer/internal/controller/middleware"
)

type Router struct {
	router *gin.Engine
}

func NewRouter(router *gin.Engine) *Router {
	return &Router{router: router}
}

func (r *Router) GetRouter() *gin.Engine {
	// load html files
	r.router.LoadHTMLGlob("template/*")
	// load static files
	r.router.Static("/static", "static")
	// log middleware
	r.router.Use(middleware.JsonLoggerMiddleware())
	// routing
	r.router.GET("/first", r.LoadHtmlPage("first.html"))
	r.router.GET("/second", r.LoadHtmlPage("second.html"))
	r.router.GET("/third", r.LoadHtmlPage("third.html"))
	return r.router
}

func (r *Router) LoadHtmlPage(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, name, nil)
	}
}
