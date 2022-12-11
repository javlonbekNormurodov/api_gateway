package main

import (
	"github.com/gin-gonic/gin"
	"github.com/javlonbekNormurodov/api_gateway/api"
	"github.com/javlonbekNormurodov/api_gateway/api/handlers"
	"github.com/javlonbekNormurodov/api_gateway/config"
	"github.com/javlonbekNormurodov/api_gateway/pkg/logger"
	"github.com/javlonbekNormurodov/api_gateway/services"
)

func main() {
	cfg := config.Load()

	grpcSvcs, err := services.NewGrpcClients(cfg)
	if err != nil {
		panic(err)
	}

	var loggerLevel = new(string)
	*loggerLevel = logger.LevelDebug

	switch cfg.Environment {
	case config.DebugMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case config.TestMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		*loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)
	}

	log := logger.NewLogger("albatta_go_api_gateway", *loggerLevel)
	defer func() {
		err := logger.Cleanup(log)
		if err != nil {
			return
		}
	}()

	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	h := handlers.NewHandler(cfg, log, grpcSvcs)

	api.SetUpAPI(r, h, cfg)

	if err := r.Run(cfg.HTTPPort); err != nil {
		return
	}
}
