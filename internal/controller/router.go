package controller

import (
	"html/template"

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

var (
	mainPage           = template.Must(template.ParseFiles("resources/template/home.html"))
	coursesPage        = template.Must(template.ParseFiles("resources/template/courses.html"))
	courseArticlesPage = template.Must(template.ParseFiles("resources/template/articles.html"))
	courseForm         = template.Must(template.ParseFiles("resources/template/course_form.html"))
	articleForm        = template.Must(template.ParseFiles("resources/template/article_form.html"))
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
	r.router.Use(middleware.ErrorHandler)
	// routing
	// r.router.GET("/", r.mainPage)
	r.router.GET("/article", r.loadHtmlPageById())
	r.router.GET("/article/create", r.articleForm)
	r.router.POST("/article", r.articleFormSubmit)
	r.router.GET("/course", r.courseArticles)
	r.router.GET("/", r.coursesList)
	r.router.GET("/course/create", r.courseForm)
	r.router.POST("/course", r.courseFormSubmit)

	return r.router
}
