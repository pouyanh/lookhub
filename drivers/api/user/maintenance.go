package user

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/pouyanh/lookhub/config"
	"github.com/pouyanh/lookhub/drivers/api/user/restful/models"
	"github.com/pouyanh/lookhub/drivers/api/user/restful/restapi/operations/maintenance"
)

func (s server) handleHealthCheck(_ maintenance.HealthCheckParams) middleware.Responder {
	return maintenance.NewHealthCheckOK().WithPayload(&models.ServiceHealth{
		Status:    models.ServiceStatusReady,
		Commit:    config.Commit,
		Version:   config.Version,
		BuildDate: config.BuildDate,
	})
}
