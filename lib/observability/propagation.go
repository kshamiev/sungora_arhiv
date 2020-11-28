package observability

import (
	"net/http"
	"strings"

	"go.opencensus.io/plugin/ochttp/propagation/b3"
	"go.opencensus.io/trace"
	"go.opencensus.io/trace/propagation"
)

const headerName = "Uber-Trace-Id"

type OpenTracingHTTPFormat struct {
}

func NewOpenTracingB3Propagation() propagation.HTTPFormat {
	return &OpenTracingHTTPFormat{}
}

func (n *OpenTracingHTTPFormat) SpanContextFromRequest(req *http.Request) (trace.SpanContext, bool) {
	// Uber-Trace-Id=[4c482a9866071599:7fce86b544364257:4c482a9866071599:1]
	// {trace-id}:{span-id}:{parent-span-id}:{flags}
	// parent-span-id is deprecated, ignoring in here
	// https://www.jaegertracing.io/docs/1.17/client-libraries/

	ts := strings.Split(req.Header.Get(headerName), ":")
	if len(ts) < 4 {
		return trace.SpanContext{}, false
		// "The trace header must have at least 4 values
	}
	// From here, use the b3 methods.
	traceID, ok := b3.ParseTraceID(ts[0])
	if !ok {
		return trace.SpanContext{}, false
	}
	spanID, ok := b3.ParseSpanID(ts[1])
	if !ok {
		return trace.SpanContext{}, false
	}
	sampled, _ := b3.ParseSampled(ts[3])
	return trace.SpanContext{
		TraceID:      traceID,
		SpanID:       spanID,
		TraceOptions: sampled,
	}, true
}

func (n *OpenTracingHTTPFormat) SpanContextToRequest(sc trace.SpanContext, req *http.Request) {
	encodedValue := sc.TraceID.String() + ":" + sc.SpanID.String() + ":0:"
	var sampled string
	if sc.IsSampled() {
		sampled = "1"
	} else {
		sampled = "0"
	}
	encodedValue += sampled
	req.Header.Set(headerName, encodedValue)
}
