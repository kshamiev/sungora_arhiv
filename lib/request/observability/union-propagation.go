package observability

import (
	"net/http"

	"go.opencensus.io/plugin/ochttp/propagation/b3"
	"go.opencensus.io/trace"
	"go.opencensus.io/trace/propagation"
)

type HTTPFormat struct {
	b3 b3.HTTPFormat
	ot OpenTracingHTTPFormat
}

func NewHTTPFormat() propagation.HTTPFormat {
	return &HTTPFormat{}
}

func (f *HTTPFormat) SpanContextFromRequest(req *http.Request) (sc trace.SpanContext, ok bool) {
	if sc, ok := f.b3.SpanContextFromRequest(req); ok {
		return sc, true
	}
	if sc, ok := f.ot.SpanContextFromRequest(req); ok {
		return sc, true
	}
	return trace.SpanContext{}, false
}
func (f *HTTPFormat) SpanContextToRequest(sc trace.SpanContext, req *http.Request) {
	f.b3.SpanContextToRequest(sc, req)
	f.ot.SpanContextToRequest(sc, req)
}
