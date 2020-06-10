package config

import (
	"github.com/fingerpeople/guppy-rest/src/app/v1/utility/guppy"
	"github.com/fingerpeople/guppy-rest/src/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRouter server router configuration
func SetupRouter() *gin.Engine {
	// setup default config
	defaultConfig()
	defaultMiddleware := middleware.DefaultMiddleware{}
	router := gin.Default()
	router.Use(defaultMiddleware.CORSMiddleware())
	return router
}

func defaultConfig() {
	guppyConfig := guppy.GuppyLibraryHandler()
	guppyConfig.CheckUserConfig()
}
