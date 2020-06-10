package routes

import (
	config "github.com/fingerpeople/guppy-rest/src/app/v1/api/configuration/controller"
	guppy "github.com/fingerpeople/guppy-rest/src/app/v1/api/guppy/controller"
	health "github.com/fingerpeople/guppy-rest/src/app/v1/api/health/controller"
	"github.com/fingerpeople/guppy-rest/src/middleware"
	"github.com/gin-gonic/gin"
)

// VERSION ...
const VERSION = "v1"

// V1RouterLoader types
type V1RouterLoader struct {
	Middleware middleware.DefaultMiddleware
	Health     health.HealthControllerInterface
	Guppy      guppy.GuppyControllerInterface
	Config     config.ConfigurationControllerInterface
}

// V1RouterLoaderHandler ...
func V1RouterLoaderHandler() *V1RouterLoader {
	return &V1RouterLoader{
		Health: health.HealthControllerHandler(),
		Guppy:  guppy.GuppyControllerHandler(),
		Config: config.ConfigurationControllerHandler(),
	}
}

// V1RouterLoaderInterface ...
type V1RouterLoaderInterface interface {
	V1Routes(router *gin.Engine)
}

// V1Routes Params
// @router: gin.Engine
func (rLoader *V1RouterLoader) V1Routes(router *gin.Engine) {
	rLoader.initDocs(router)
	rLoader.initHealth(router)
	rLoader.initGuppy(router)
	rLoader.initConfig(router)
}
