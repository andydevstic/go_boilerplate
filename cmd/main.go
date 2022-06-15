package main

import (
	"fmt"

	"github.com/andydevstic/boilerplate-backend/config"
	"github.com/andydevstic/boilerplate-backend/core"
	"github.com/andydevstic/boilerplate-backend/db"
	"github.com/andydevstic/boilerplate-backend/shared/middlewares"
	_ "github.com/andydevstic/boilerplate-backend/shared/utils"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func main() {
	config, err := config.GetConfig(".")
	if err != nil {
		log.Error().Msg(fmt.Sprintf("Failed to read config %s", err))
	}

	if config.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	dbConn, err := db.ConnectDb(config)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("failed to establish database connection: %s", err.Error()))
	}

	defer dbConn.Close()

	core.GenerateAppState(dbConn)

	log.Info().Msg("Database connected successfully!")

	app := gin.Default()

	app.NoRoute(middlewares.NoRouteHandler)

	apiRouter := app.Group("/api")

	router.NewHealth().Route(apiRouter)
	router.NewAuth().Route(apiRouter)
	router.NewUser().Route(apiRouter)
	router.NewLink().Route(apiRouter)

	err = app.Run(fmt.Sprintf("0.0.0.0:%d", config.Port))

	if err != nil {
		log.Error().Msg(fmt.Sprintf("failed to start http server: %s", err.Error()))
	}
}
