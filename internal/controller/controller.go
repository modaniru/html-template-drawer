package controller

import (
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/modaniru/html-template-drawer/internal/entity"
)

func (r *Router) courseForm(c *gin.Context) {
	courseForm.Execute(c.Writer, nil)
}

func (r *Router) courseFormSubmit(c *gin.Context) {
	imageLink := c.PostForm("image")
	if _, err := url.ParseRequestURI(imageLink); err != nil {
		c.JSON(400, err.Error())
	}
	id, err := r.service.CourseService.SaveCourse(c, entity.CourseForm{Name: c.PostForm("name"), Image: imageLink})
	if err != nil {
		c.JSON(400, err.Error())
	}
	c.JSON(200, "success: "+id)
}

func (r *Router) articleForm(c *gin.Context) {
	courses, err := r.service.CourseService.GetAllCourses(c)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	articleForm.Execute(c.Writer, courses)
}

func (r *Router) articleFormSubmit(c *gin.Context) {
	err := r.service.ArticleService.SaveArticle(c, entity.ArticleForm{
		Title:        c.PostForm("title"),
		TemplateName: c.PostForm("name"),
		CourseId:     c.PostForm("course"),
	})
	if err != nil {
		c.JSON(400, err.Error())
	}
	c.JSON(200, "success")
}

func (r *Router) loadHtmlPageById() gin.HandlerFunc {
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

func (r *Router) courseArticles(c *gin.Context) {
	courseId := c.Query("id")
	fmt.Println(courseId)
	if courseId == "" {
		// TODO 404
		c.Abort()
		return
	}
	list, err := r.service.ArticleService.GetCourseArticles(c, courseId)
	if err != nil {
		// TODO error
		fmt.Println(err.Error())
		c.Abort()
		return
	}
	fmt.Println(list)
	courseArticlesPage.Execute(c.Writer, list)
}

func (r *Router) coursesList(c *gin.Context) {
	list, err := r.service.CourseService.GetAllCourses(c)
	fmt.Println(list)
	if err != nil {
		c.JSON(400, err)
		return
	}
	coursesPage.Execute(c.Writer, list)
}
