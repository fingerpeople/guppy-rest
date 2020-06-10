package router

import (
	v1 "github.com/fingerpeople/guppy-rest/src/app/v1/routes"
	"github.com/gin-gonic/gin"
)

// LoadRouter params
// @routers: gin.Engine
func LoadRouter(routers *gin.Engine) {
	version1 := v1.V1RouterLoaderHandler()
	version1.V1Routes(routers)
}
