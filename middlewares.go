package main

import (
	"github.com/go-kit/kit/endpoint"
	stdopentracing "github.com/opentracing/opentracing-go"
	"golang.org/x/net/context"
)

func EndpointTracingMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		defer func() {
			if err != nil {
				if span := stdopentracing.SpanFromContext(ctx); span != nil {
					span.SetTag("error", err)
				}
			}
		}()
		return next(ctx, request)
	}
}
