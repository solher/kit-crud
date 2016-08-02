package library

import (
	"errors"
	"time"

	stdopentracing "github.com/opentracing/opentracing-go"
	otext "github.com/opentracing/opentracing-go/ext"
	"golang.org/x/net/context"
)

type Repository struct {
	tracer stdopentracing.Tracer
}

func NewRepository(tracer stdopentracing.Tracer) *Repository {
	return &Repository{
		tracer: tracer,
	}
}

func (s *Repository) Run(ctx context.Context, query string) (string, error) {
	operationName := "database"
	var clientSpan stdopentracing.Span
	if parentSpan := stdopentracing.SpanFromContext(ctx); parentSpan != nil {
		clientSpan = s.tracer.StartSpan(
			operationName,
			stdopentracing.ChildOf(parentSpan.Context()),
		)
	} else {
		clientSpan = s.tracer.StartSpan(operationName)
	}
	defer clientSpan.Finish()
	otext.SpanKindRPCClient.Set(clientSpan)
	clientSpan.SetTag("query", query)
	clientSpan.SetTag("target", "cassandra:9042")

	// Database call simulation
	time.Sleep(2 * time.Millisecond)

	err := errors.New("db error")
	if err != nil {
		clientSpan.SetTag("error", err.Error())
		return "", err
	}

	return "", nil
}
