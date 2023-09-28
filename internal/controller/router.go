package controller

import (
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/modaniru/html-template-drawer/internal/controller/middleware"
)

type Router struct {
	router *gin.Engine
}

func NewRouter(router *gin.Engine) *Router {
	return &Router{router: router}
}

var (
	mainPage = template.Must(template.ParseFiles("resources/template/home.html"))
)

func (r *Router) GetRouter() *gin.Engine {
	// load html files
	// r.router.LoadHTMLGlob("template/**/*")

	// load articles
	r.router.LoadHTMLGlob("resources/template/articles/*")
	// load css files
	r.router.Static("css", "./resources/css")
	// log middleware
	r.router.Use(middleware.JsonLoggerMiddleware())
	// routing
	r.router.GET("/", r.mainPage)

	return r.router
}

func (r *Router) LoadHtmlPage(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, name, nil)
	}
}

// This page must return list of course
func (r *Router) mainPage(c *gin.Context) {
	mainPage.Execute(c.Writer, nil)
}
