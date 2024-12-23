package restapi

import (
	"crypto/tls"

	"github.com/go-openapi/runtime/middleware"
)

var (
	TlsConfigurator  = func(*tls.Config) {}
	Middleware       = middleware.PassthroughBuilder
	GlobalMiddleware = middleware.PassthroughBuilder
)
