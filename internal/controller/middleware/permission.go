package middleware

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func Permission(c *gin.Context) {
	key := c.Query("key")
	if key != os.Getenv("SECRET") {
		c.AbortWithError(403, fmt.Errorf("invalid key"))
		return
	}
	c.Next()
}
