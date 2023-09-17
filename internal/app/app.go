package app

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/modaniru/html-template-drawer/internal/controller"
)

func Run() {
	configureLogger()
	configureDotEnv()
	engine := gin.New()
	r := controller.NewRouter(engine)
	router := r.GetRouter()
	router.Run(":" + os.Getenv("PORT"))
}
