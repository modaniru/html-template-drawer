package controller

import (
	"fmt"
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
	// TODO localhost/?id=go_course
	// log middleware
	r.router.Use(middleware.JsonLoggerMiddleware())
	// routing
	r.router.GET("/", r.mainPage)
	r.router.GET("/articles", r.LoadHtmlPage())

	return r.router
}

func (r *Router) LoadHtmlPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		page := c.Query("id")
		fmt.Println(page)
		c.HTML(200, page+".html", nil)
	}
}

// This page must return list of course
func (r *Router) mainPage(c *gin.Context) {
	mainPage.Execute(c.Writer, nil)
}
