package cmd

import (
	"context"

	"github.com/gogf/gf/contrib/metric/otelmetric/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gmetric"
	"github.com/oldme-git/36hour/app/gateway/internal/controller/account"
	"github.com/oldme-git/36hour/app/gateway/internal/controller/login"
	"github.com/oldme-git/36hour/app/gateway/internal/controller/seat"
	"github.com/oldme-git/36hour/app/gateway/internal/logic/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/exporters/prometheus"
)

var (
	meter = gmetric.GetGlobalProvider().Meter(gmetric.MeterOption{
		Instrument:        "gateway",
		InstrumentVersion: "v1.0",
	})
	counter = meter.MustCounter(
		"gateway.counter",
		gmetric.MetricOption{
			Help: "Gateway for Counter usage",
			Unit: "bytes",
		},
	)
	upDownCounter = meter.MustUpDownCounter(
		"gateway.updown_counter",
		gmetric.MetricOption{
			Help: "Gateway for UpDownCounter usage",
			Unit: "%",
		},
	)
	histogram = meter.MustHistogram(
		"gateway.histogram",
		gmetric.MetricOption{
			Help:    "Gateway for histogram usage",
			Unit:    "ms",
			Buckets: []float64{0, 10, 20, 50, 100, 500, 1000, 2000, 5000, 10000},
		},
	)
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http gateway server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// Prometheus exporter to export metrics as Prometheus format.
			exporter, err := prometheus.New(
				prometheus.WithoutCounterSuffixes(),
				prometheus.WithoutUnits(),
			)
			if err != nil {
				g.Log().Fatal(ctx, err)
			}

			// OpenTelemetry provider.
			provider := otelmetric.MustProvider(otelmetric.WithReader(exporter))
			provider.SetAsGlobal()
			defer provider.Shutdown(ctx)

			// Counter.
			counter.Inc(ctx)
			counter.Add(ctx, 10)

			// UpDownCounter.
			upDownCounter.Inc(ctx)
			upDownCounter.Add(ctx, 10)
			upDownCounter.Dec(ctx)

			// Record values for histogram.
			histogram.Record(1)
			histogram.Record(20)
			histogram.Record(30)
			histogram.Record(101)
			histogram.Record(2000)
			histogram.Record(9000)
			histogram.Record(20000)

			s := g.Server()
			s.BindHandler("/metrics", ghttp.WrapH(promhttp.Handler()))
			s.Group("/", func(group *ghttp.RouterGroup) {
				// 允许跨域
				// group.Middleware(ghttp.MiddlewareCORS)
				group.Middleware(middleware.Response)

				group.Group("/v1", func(group *ghttp.RouterGroup) {
					// App v1 路由
					group.Group("/", func(group *ghttp.RouterGroup) {
						group.Bind(login.NewV1())
						group.Bind(seat.NewV1())
					})

					// Auth Middleware
					group.Group("/", func(group *ghttp.RouterGroup) {
						group.Middleware(middleware.Auth)
						group.Bind(account.NewV1())
					})
				})
			})
			s.Run()
			return nil
		},
	}
)
