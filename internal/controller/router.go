package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/modaniru/html-template-drawer/internal/controller/middleware"
	"github.com/modaniru/html-template-drawer/internal/service"
)

type Router struct {
	router  *gin.Engine
	service *service.Service
}

func NewRouter(router *gin.Engine, service *service.Service) *Router {
	return &Router{router: router, service: service}
}

func (r *Router) GetRouter() *gin.Engine {
	// load html files
	r.router.LoadHTMLGlob("resources/template/*/*.html")
	// load css files
	r.router.Static("css", "./resources/css")
	r.router.Static("img", "./resources/img")
	r.router.Static("js", "./resources/js")
	r.router.Static("fonts", "./resources/fonts")
	// log middleware
	r.router.Use(middleware.JsonLoggerMiddleware())
	r.router.Use(middleware.ErrorHandler)
	// routing
	r.router.GET("/", r.mainPage)
	r.router.GET("/article", r.loadHtmlPageById())
	r.router.GET("/article/create", r.articleForm)
	r.router.POST("/article", r.articleFormSubmit)
	r.router.GET("/course", r.courseArticles)
	r.router.GET("/list", r.coursesList)
	r.router.GET("/course/create", r.courseForm)
	r.router.POST("/course", r.courseFormSubmit)

	return r.router
}
