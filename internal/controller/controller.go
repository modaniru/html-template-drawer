package controller

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/modaniru/html-template-drawer/internal/entity"
	"github.com/modaniru/html-template-drawer/pkg"
)

func (r *Router) courseForm(c *gin.Context) {
	c.HTML(200, "s_course_form.html", map[string]any{
		"key": c.Query("key"),
	})
}

func (r *Router) courseFormSubmit(c *gin.Context) {
	imageLink := c.PostForm("image")
	if _, err := url.ParseRequestURI(imageLink); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	_, err := r.service.CourseService.SaveCourse(c, entity.CourseForm{Name: c.PostForm("name"), Image: imageLink})
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Redirect(301, "/admin/course/create?key="+c.Query("key"))
}

func (r *Router) articleForm(c *gin.Context) {
	courses, err := r.service.CourseService.GetAllCourses(c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.HTML(200, "s_article_form.html", map[string]any{
		"courses": courses,
		"files":   pkg.GetAllArticlesFiles(),
		"key":     c.Query("key"),
	})
}

func (r *Router) articleFormSubmit(c *gin.Context) {
	err := r.service.ArticleService.SaveArticle(c, entity.ArticleForm{
		Title:        c.PostForm("title"),
		TemplateName: c.PostForm("name"),
		CourseId:     c.PostForm("course"),
	})
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Redirect(301, fmt.Sprintf("/admin/article/create?key=%s", c.Query("key")))
}

func (r *Router) loadHtmlPageById() gin.HandlerFunc {
	return func(c *gin.Context) {
		page := c.Query("id")
		if strings.HasPrefix(page, "s_") {
			c.AbortWithError(404, fmt.Errorf("error was not found"))
			return
		}
		articles := pkg.GetAllArticlesFiles()
		for _, a := range articles {
			fmt.Println(a, page)
			if a == page {
				c.HTML(200, page+".html", nil)
				return
			}
		}
		c.AbortWithError(404, fmt.Errorf("error was not found"))
	}
}

// This page must return list of course
func (r *Router) mainPage(c *gin.Context) {
	c.HTML(200, "s_home.html", nil)
}

func (r *Router) courseArticles(c *gin.Context) {
	courseId := c.Query("id")
	fmt.Println(courseId)
	if courseId == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("course id is not present"))
		return
	}
	list, err := r.service.ArticleService.GetCourseArticles(c, courseId)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.HTML(200, "s_articles.html", list)
}

func (r *Router) coursesList(c *gin.Context) {
	list, err := r.service.CourseService.GetAllCourses(c)
	fmt.Println(list)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.HTML(200, "s_courses.html", list)
}

func (r *Router) adminCourseList(c *gin.Context) {
	list, err := r.service.CourseService.GetAllCourses(c)
	fmt.Println(list)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.HTML(200, "s_admin_courses.html", map[string]any{
		"List": list,
		"Key":  c.Query("key"),
	})
}

func (r *Router) deleteCourse(c *gin.Context) {
	courseId := c.Query("id")
	err := r.service.CourseService.DeleteCourse(c, courseId)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Redirect(301, fmt.Sprintf("/admin/list?key=%s", c.Query("key")))
}

func (r *Router) adminArticleList(c *gin.Context) {
	courseId := c.Query("id")
	fmt.Println(courseId)
	if courseId == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("course id is not present"))
		return
	}
	list, err := r.service.ArticleService.GetCourseArticles(c, courseId)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.HTML(200, "s_admin_articles.html", map[string]any{
		"List":    list,
		"Key":     c.Query("key"),
		"TitleId": c.Query("id"),
	})
}

func (r *Router) deleteArticle(c *gin.Context) {
	articleId := c.Query("id")
	err := r.service.ArticleService.DeleteById(c, articleId)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Redirect(301, fmt.Sprintf("/admin/course?id=%s&key=%s", c.Query("course"), c.Query("key")))
}
