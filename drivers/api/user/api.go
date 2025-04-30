package user

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/janstoon/toolbox/handywares"
	"github.com/janstoon/toolbox/tricks"
	"go.opentelemetry.io/otel"
	"golang.org/x/sync/errgroup"

	"github.com/pouyanh/lookhub/drivers/api/user/restful/restapi"
	"github.com/pouyanh/lookhub/drivers/api/user/restful/restapi/operations"
	"github.com/pouyanh/lookhub/drivers/api/user/restful/restapi/operations/dnslv"
	"github.com/pouyanh/lookhub/drivers/api/user/restful/restapi/operations/maintenance"
	"github.com/pouyanh/lookhub/settings"
)

var tracer = otel.Tracer("github.com/pouyanh/lookhub/drivers/api/user")

type server struct {
	apps      []any
	listeners []net.Listener

	allowedOrigins []string
}

func newServer(apiSs settings.API, listeners []net.Listener, apps []any) server {
	return server{
		apps:      apps,
		listeners: listeners,

		allowedOrigins: apiSs.Cors.OriginWhitelist,
	}
}

func (s server) Run(ctx context.Context) error {
	handler := s.handler(s.apiOptions()...)

	servers := tricks.Map(s.listeners, func(src net.Listener) *http.Server {
		return &http.Server{
			Handler: handler,
		}
	})

	wg := errgroup.Group{}
	for k := range s.listeners {
		func(k int) {
			wg.Go(func() error {
				if err := servers[k].Serve(s.listeners[k]); !errors.Is(err, http.ErrServerClosed) {
					return err
				}

				return nil
			})

			wg.Go(func() error {
				<-ctx.Done()

				return servers[k].Shutdown(context.TODO())
			})
		}(k)
	}

	return wg.Wait()
}

func (s server) handler(opts ...tricks.Option[operations.LookHubUserAPI]) http.Handler {
	spec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewLookHubUserAPI(spec)
	s.attachAccessProviders(api)
	s.attachMiddlewares(api)
	s.attachHandlers(api)

	var gmw handywares.HttpMiddlewareStack
	restapi.GlobalMiddleware = gmw.
		Push(handywares.HttpOpenTelemetryMiddleware(
			tracer, api.Context(),
			handywares.OtelHttpSpanNamePrefix(spec.Spec().Info.Title),
			handywares.OtelHttpOperationIdException(s.noTraceOperationIds(api.Context())...),
		)).
		Push(handywares.HttpPanicRecoverMiddleware()).
		Push(handywares.HttpCrossOriginResourceSharingPolicyMiddleware(
			handywares.CorsAllowOrigins(s.allowedOrigins...),
			handywares.CorsAllowMethods(
				http.MethodHead,
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete,
			),
			handywares.CorsAllowHeaders("*"),
		))

	return restapi.ConfigureAPI(api, opts...)
}

// attachAccessProviders supposed to attach access utilities: authenticators, authorizers
func (server) attachAccessProviders(_ *operations.LookHubUserAPI) {}

// attachMiddlewares supposed to attach application-specific middlewares
// like rate limiter
func (server) attachMiddlewares(_ *operations.LookHubUserAPI) {}

func (s server) attachHandlers(api *operations.LookHubUserAPI) {
	// Maintenance
	api.MaintenanceHealthCheckHandler = maintenance.HealthCheckHandlerFunc(s.handleHealthCheck)

	// DNSLV
	api.DnslvDnslvLookupHandler = dnslv.DnslvLookupHandlerFunc(s.handleDNSLVLookup)
}

func (server) noTraceOperationIds(mctx *middleware.Context) []string {
	oids := make([]string, 0)
	// mr, _ := mctx.LookupRoute(maintenance.NewHealthCheckParams().HTTPRequest)
	// oids = append(xOpIds, mr.Operation.ID)
	oids = append(oids, "healthCheck")

	return oids
}

func (server) apiOptions() []tricks.Option[operations.LookHubUserAPI] {
	return []tricks.Option[operations.LookHubUserAPI]{
		tricks.InPlaceOption[operations.LookHubUserAPI](func(s *operations.LookHubUserAPI) {
			s.UseSwaggerUI()
			s.JSONConsumer = runtime.JSONConsumer()

			s.JSONProducer = runtime.JSONProducer()
		}),
		tricks.InPlaceOption[operations.LookHubUserAPI](func(s *operations.LookHubUserAPI) {
			s.PreServerShutdown = func() {}
			s.ServerShutdown = func() {}
		}),
	}
}
