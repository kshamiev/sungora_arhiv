package observability

import (
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
	"google.golang.org/grpc"
)

var (
	// Client View
	ClientSentBytesDistribution = &view.View{
		Name:        "http/client/sent_bytes",
		Measure:     ochttp.ClientSentBytes,
		Aggregation: ochttp.DefaultSizeDistribution,
		Description: "Total bytes sent in request body (not including headers), by HTTP method and response status",
		TagKeys:     []tag.Key{ochttp.KeyClientMethod, ochttp.KeyClientStatus, ochttp.KeyClientPath},
	}

	ClientRoundtripLatencyDistribution = &view.View{
		Name:        "http/client/roundtrip_latency",
		Measure:     ochttp.ClientRoundtripLatency,
		Aggregation: ochttp.DefaultLatencyDistribution,
		Description: "End-to-end latency, by HTTP method and response status",
		TagKeys:     []tag.Key{ochttp.KeyClientMethod, ochttp.KeyClientStatus, ochttp.KeyClientPath},
	}

	ClientCompletedCount = &view.View{
		Name:        "http/client/completed_count",
		Measure:     ochttp.ClientRoundtripLatency,
		Aggregation: view.Count(),
		Description: "Count of completed requests, by HTTP method and response status",
		TagKeys:     []tag.Key{ochttp.KeyClientMethod, ochttp.KeyClientStatus, ochttp.KeyClientPath},
	}
	// Server View

	ServerRequestCountView = &view.View{
		Name:        "http/server/response_count",
		Description: "Server response count by status code",
		TagKeys:     []tag.Key{ochttp.Method, ochttp.Path, ochttp.StatusCode},
		Measure:     ochttp.ServerLatency,
		Aggregation: view.Count(),
	}
	ServerLatencyView = &view.View{
		Name:        "http/server/latency",
		Description: "Latency distribution of HTTP requests",
		Measure:     ochttp.ServerLatency,
		TagKeys:     []tag.Key{ochttp.Method, ochttp.Path, ochttp.StatusCode},
		Aggregation: ochttp.DefaultLatencyDistribution,
	}
)
var AllHTTPViews = []*view.View{
	ClientSentBytesDistribution,
	ClientRoundtripLatencyDistribution,
	ClientCompletedCount,
	ServerRequestCountView,
	ServerLatencyView,
}

func TelemetryInterceptor() grpc.ServerOption {
	return grpc.StatsHandler(
		&ocgrpc.ServerHandler{
			IsPublicEndpoint: false,
		},
	)
}
