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
	r.router.Use(func(ctx *gin.Context) {
		ctx.Header("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")
		ctx.Header("Expires", "Sat, 26 Jul 1997 05:00:00 GMT")
	})
	r.router.Use(middleware.JsonLoggerMiddleware())
	r.router.Use(middleware.ErrorHandler)
	// routing

	r.router.GET("/", r.mainPage)
	r.router.GET("/article", r.loadHtmlPageById())
	r.router.GET("/course", r.courseArticles)
	r.router.GET("/list", r.coursesList)

	g := r.router.Group("/admin", middleware.Permission)
	g.GET("/article/create", r.articleForm)
	g.POST("/article", r.articleFormSubmit)
	g.GET("/course/create", r.courseForm)
	g.POST("/course", r.courseFormSubmit)
	g.GET("/list", r.adminCourseList)
	g.POST("/course/delete", r.deleteCourse)
	g.GET("/course", r.adminArticleList)
	g.POST("/article/delete", r.deleteArticle)

	return r.router
}
