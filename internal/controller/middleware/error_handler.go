package middleware

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

var (
	errorPage = template.Must(template.ParseFiles("resources/template/error.html"))
)

func ErrorHandler(c *gin.Context) {
	c.Next()
	if len(c.Errors) == 0 {
		return
	}
	for _, err := range c.Errors {
		errorPage.Execute(c.Writer, map[string]any{
			"error":  err.Err.Error(),
			"status": c.Writer.Status(),
		})
		break
	}
}
