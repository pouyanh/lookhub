package lutel

import "github.com/prometheus/client_golang/prometheus"

var (
	RequestCnt = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "snp_lookhub_requests_total",
		Help: "Number of requests",
	})

	CacheHitCnt = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "snp_lookhub_cache_hits_total",
		Help: "Number of requests responded from cache",
	})

	CacheMissCnt = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "snp_lookhub_cache_misses_total",
		Help: "Number of requests not responded from cache",
	})
)

func init() {
	prometheus.MustRegister(
		RequestCnt,
		CacheHitCnt,
		CacheMissCnt,
	)
}
