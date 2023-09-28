package app

import (
	"database/sql"
	log "log/slog"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/modaniru/html-template-drawer/internal/controller"
	"github.com/modaniru/html-template-drawer/internal/storage"
)

func Run() {
	configureLogger()
	configureDotEnv()
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5555/postgres?sslmode=disable")
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	err = db.Ping()
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	engine := gin.New()
	r := controller.NewRouter(engine, storage.NewStorage(db))
	router := r.GetRouter()
	router.Run(":" + os.Getenv("PORT"))
}
