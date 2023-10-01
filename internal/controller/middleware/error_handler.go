package middleware

import (
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next()
	if len(c.Errors) == 0 {
		return
	}
	for _, err := range c.Errors {
		c.HTML(400, "s_error.html", map[string]any{
			"error":  err.Err.Error(),
			"status": c.Writer.Status(),
		})
		break
	}
}
