package main

import (
	"github.com/gin-gonic/gin"
	"test_1/internal/libs/ext_api"
)

func createGINRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)

	engine := gin.New()
	engine.Use(
		ext_api.RPSLimiter())

	return engine
}
