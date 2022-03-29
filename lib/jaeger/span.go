package jaeger

import (
	"context"

	"sample/lib/errs"

	"go.opencensus.io/trace"
)

type Span struct {
	Span *trace.Span
}

func NewSpan(ctx context.Context) *Span {
	s := &Span{}
	t, _, fun, _ := errs.Trace(2)
	_, s.Span = trace.StartSpan(ctx, fun)
	s.StringAttribute("location", t)
	return s
}

func NewSpanName(ctx context.Context, name string) *Span {
	s := &Span{}
	_, s.Span = trace.StartSpan(ctx, name)
	t, _, _, _ := errs.Trace(2)
	s.StringAttribute("location", t)
	return s
}

func (s *Span) StringAttribute(key, value string) {
	s.Span.AddAttributes(trace.StringAttribute(key, value))
}

func (s *Span) BoolAttribute(key string, value bool) {
	s.Span.AddAttributes(trace.BoolAttribute(key, value))
}

func (s *Span) Float64Attribute(key string, value float64) {
	s.Span.AddAttributes(trace.Float64Attribute(key, value))
}

func (s *Span) Float32Attribute(key string, value float32) {
	s.Span.AddAttributes(trace.Float64Attribute(key, float64(value)))
}

func (s *Span) Int64Attribute(key string, value int64) {
	s.Span.AddAttributes(trace.Int64Attribute(key, value))
}

func (s *Span) IntAttribute(key string, value int) {
	s.Span.AddAttributes(trace.Int64Attribute(key, int64(value)))
}

func (s *Span) End() {
	s.Span.End()
}
