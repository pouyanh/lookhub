package metrics

import (
	"context"
	"errors"
	"net"
	"net/http"
	"net/http/pprof"

	"github.com/janstoon/toolbox/handywares"
	"github.com/janstoon/toolbox/tricks"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/sync/errgroup"

	"github.com/pouyanh/lookhub/settings"
)

type server struct {
	listeners []net.Listener

	allowedOrigins []string
}

func newServer(apiSs settings.API, listeners []net.Listener) server {
	return server{
		listeners: listeners,

		allowedOrigins: apiSs.Cors.OriginWhitelist,
	}
}

func (s server) Run(ctx context.Context) error {
	handler := s.handler()

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

func (s server) handler() http.Handler {
	mux := http.NewServeMux()
	s.attachHandlers(mux)

	var gmw handywares.HttpMiddlewareStack
	gmw = gmw.
		Push(handywares.HttpPanicRecoverMiddleware()).
		Push(handywares.HttpCrossOriginResourceSharingPolicyMiddleware(
			handywares.CorsAllowOrigins(s.allowedOrigins...),
			handywares.CorsAllowMethods(
				http.MethodHead,
				http.MethodGet,
				http.MethodOptions,
			),
			handywares.CorsAllowHeaders("*"),
		))

	return gmw(mux)
}

func (s server) attachHandlers(mux *http.ServeMux) {
	mux.Handle("/metrics", promhttp.Handler())

	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
}
