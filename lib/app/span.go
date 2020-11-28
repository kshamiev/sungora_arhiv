package app

import (
	"context"

	"go.opencensus.io/trace"
)

type Span struct {
	Span *trace.Span
}

func NewStartSpan(ctx context.Context) *Span {
	s := &Span{}
	kind, _, fun, _ := TraceAtom(2)
	_, s.Span = trace.StartSpan(ctx, fun)
	s.StringAttribute("location", kind)
	return s
}

func NewStartSpanName(ctx context.Context, name string) *Span {
	s := &Span{}
	_, s.Span = trace.StartSpan(ctx, name)
	s.StringAttribute("location", Trace(2))
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

func (s *Span) Int64Attribute(key string, value int64) {
	s.Span.AddAttributes(trace.Int64Attribute(key, value))
}

func (s *Span) End() {
	s.Span.End()
}
