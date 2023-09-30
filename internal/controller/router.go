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
	// r.router.GET("/", r.mainPage)
	r.router.GET("/article", r.LoadHtmlPage())
	r.router.GET("/course", r.ListOfCourseArticles)
	r.router.GET("/", r.ListOfCourses)

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

func (r *Router) ListOfCourseArticles(c *gin.Context) {
	courseId := c.Query("id")
	fmt.Println(courseId)
	if courseId == "" {
		// TODO 404
		c.Abort()
		return
	}
	list, err := r.storage.Articles.GetCourseArticles(c, courseId)
	if err != nil {
		// TODO error
		fmt.Println(err.Error())
		c.Abort()
		return
	}
	fmt.Println(list)
	c.JSON(200, list)
}

func (r *Router) ListOfCourses(c *gin.Context) {
	list, err := r.storage.Courses.GetAllCourses(c)
	fmt.Println(list)
	if err != nil {
		c.JSON(400, err)
		return
	}
	coursesPage.Execute(c.Writer, list)
}
