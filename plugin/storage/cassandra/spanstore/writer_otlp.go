package spanstore

import (
	"context"
	"time"

	"github.com/jaegertracing/jaeger/pkg/cassandra"
	"github.com/jaegertracing/jaeger/pkg/metrics"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

type OTLPSpanWriter struct {
	*SpanWriter
}

func NewOTLPSpanWriter(
	session cassandra.Session,
	writeCacheTTL time.Duration,
	metricsFactory metrics.Factory,
	logger *zap.Logger,
	options ...Option,
) *OTLPSpanWriter {
	sp := NewSpanWriter(session, writeCacheTTL, metricsFactory, logger, options...)
	return &OTLPSpanWriter{
		sp,
	}
}

func (s *OTLPSpanWriter) WriteOTLPSpan(ctx context.Context, td ptrace.Traces) error {
	return nil
}
