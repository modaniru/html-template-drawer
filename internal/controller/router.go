package controller

import (
	"fmt"
	"html/template"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/modaniru/html-template-drawer/internal/controller/middleware"
	"github.com/modaniru/html-template-drawer/internal/entity"
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
	// routing
	// r.router.GET("/", r.mainPage)
	r.router.GET("/article", r.LoadHtmlPage())
	r.router.GET("/article/create", r.ArticleFromPage)
	r.router.POST("/article", r.SubmitArticle)
	r.router.GET("/course", r.ListOfCourseArticles)
	r.router.GET("/", r.ListOfCourses)
	r.router.GET("/course/create", r.CourseFormPage)
	r.router.POST("/course", r.SubmitCourse)

	return r.router
}

func (r *Router) CourseFormPage(c *gin.Context) {
	courseForm.Execute(c.Writer, nil)
}

func (r *Router) SubmitCourse(c *gin.Context) {
	imageLink := c.PostForm("image")
	if _, err := url.ParseRequestURI(imageLink); err != nil {
		c.JSON(400, err.Error())
	}
	id, err := r.storage.Courses.SaveCourse(c, entity.CourseForm{Name: c.PostForm("name"), Image: imageLink})
	if err != nil {
		c.JSON(400, err.Error())
	}
	c.JSON(200, "success: "+id)
}

func (r *Router) ArticleFromPage(c *gin.Context) {
	courses, err := r.storage.Courses.GetAllCourses(c)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	articleForm.Execute(c.Writer, courses)
}

func (r *Router) SubmitArticle(c *gin.Context) {
	err := r.storage.Articles.SaveArticle(c, entity.ArticleForm{
		Title:        c.PostForm("title"),
		TemplateName: c.PostForm("name"),
		CourseId:     c.PostForm("course"),
	})
	if err != nil {
		c.JSON(400, err.Error())
	}
	c.JSON(200, "success")
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
	courseArticlesPage.Execute(c.Writer, list)
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
