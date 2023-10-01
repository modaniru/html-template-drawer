package app

import (
	"database/sql"
	log "log/slog"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/modaniru/html-template-drawer/internal/controller"
	"github.com/modaniru/html-template-drawer/internal/service"
	"github.com/modaniru/html-template-drawer/internal/storage"
)

func Run() {
	configureLogger()
	configureDotEnv()
	db, err := sql.Open("postgres", os.Getenv("DB_SOURCE"))
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	err = db.Ping()
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	storage := storage.NewStorage(db)
	service := service.CreateService(storage)
	engine := gin.New()
	r := controller.NewRouter(engine, service)
	router := r.GetRouter()
	router.Run(":" + os.Getenv("PORT"))
}
