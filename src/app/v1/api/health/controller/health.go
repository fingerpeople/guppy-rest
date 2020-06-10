package controller

import (
	"github.com/fingerpeople/guppy-rest/src/app/v1/api/health/service"
	"github.com/fingerpeople/guppy-rest/src/app/v1/utility/rest"
	"github.com/gin-gonic/gin"
)

// HealthController types
type HealthController struct {
	Service service.HealthServiceInterface
}

// HealthControllerHandler ...
func HealthControllerHandler() *HealthController {
	return &HealthController{
		Service: service.HealthServiceHandler(),
	}
}

// HealthControllerInterface ...
type HealthControllerInterface interface {
	Health(context *gin.Context)
}

// Health params
// @contex: gin Context
func (ctrl *HealthController) Health(context *gin.Context) {
	data := ctrl.Service.HealthService()
	rest.SuccessResponse(context, data, nil, "")
	return
}
