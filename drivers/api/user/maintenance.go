package user

import (
	"github.com/go-openapi/runtime/middleware"

	"gitlab.snapp.ir/pouyanh/lookhub/config"
	"gitlab.snapp.ir/pouyanh/lookhub/drivers/api/user/restful/models"
	"gitlab.snapp.ir/pouyanh/lookhub/drivers/api/user/restful/restapi/operations/maintenance"
)

func (s server) handleHealthCheck(_ maintenance.HealthCheckParams) middleware.Responder {
	return maintenance.NewHealthCheckOK().WithPayload(&models.ServiceHealth{
		Status:    models.ServiceStatusReady,
		Commit:    config.Commit,
		Version:   config.Version,
		BuildDate: config.BuildDate,
	})
}
