package app

import (
	"fmt"
	"os"
	"showcase/config"
	"showcase/repository"
	"showcase/route"
	"showcase/service"

	"github.com/gin-gonic/gin"
)

var router = gin.New()

func StartApplication() {
	repo := repository.NewRepo(config.PSQL.DB, config.GORM.DB)
	app := service.NewService(repo)
	route.RegisterApi(router, app)

	port := os.Getenv("APP_PORT")
	router.Run(fmt.Sprintf(":%s", port))
}
