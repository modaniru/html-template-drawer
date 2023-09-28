package controller

import (
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/modaniru/html-template-drawer/internal/controller/middleware"
	"github.com/modaniru/html-template-drawer/internal/storage"
)

type Router struct {
	router  *gin.Engine
	storage *storage.Storage
}

func NewRouter(router *gin.Engine, storage *storage.Storage) *Router {
	return &Router{router: router, storage: storage}
}

var (
	mainPage    = template.Must(template.ParseFiles("resources/template/home.html"))
	coursesPage = template.Must(template.ParseFiles("resources/template/courses.html"))
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
	r.router.GET("/courses", r.ListOfCourses)

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

func (r *Router) ListOfCourses(c *gin.Context) {
	list, err := r.storage.Courses.GetAllCourses(c)
	fmt.Println(list)
	if err != nil {
		c.Abort()
		return
	}
	coursesPage.Execute(c.Writer, list)
}
