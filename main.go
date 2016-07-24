package main

import (
	"flag"
	"net"
	"os"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/solher/kit-crud/library"

	zipkin "github.com/openzipkin/zipkin-go-opentracing"

	"github.com/solher/kit-crud/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var (
		grpcAddr   = flag.String("grpc.addr", ":8082", "gRPC (HTTP) listen address")
		zipkinAddr = flag.String("zipkin.addr", "", "Enable Zipkin tracing via a Scribe server host:port")
	)
	flag.Parse()

	// Logging domain.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)
		logger = log.NewContext(logger).With("caller", log.DefaultCaller)
	}

	// Tracing domain.
	var tracer stdopentracing.Tracer
	{
		if *zipkinAddr != "" {
			logger := log.NewContext(logger).With("tracer", "Zipkin")
			logger.Log("msg", "sending trace to "+*zipkinAddr)
			collector, err := zipkin.NewScribeCollector(
				*zipkinAddr,
				3*time.Second,
				zipkin.ScribeLogger(logger),
			)
			if err != nil {
				logger.Log("err", err)
				os.Exit(1)
			}
			tracer, err = zipkin.NewTracer(
				zipkin.NewRecorder(collector, false, "kit-crud:8082", "Library"),
			)
			if err != nil {
				logger.Log("err", err)
				os.Exit(1)
			}
		} else {
			logger := log.NewContext(logger).With("tracer", "none")
			logger.Log("msg", "tracing disabled")
			tracer = stdopentracing.GlobalTracer() // no-op
		}
	}

	// Business domain.
	var service library.Service
	{
		service = library.NewService()
		service = library.ServiceLoggingMiddleware(logger)(service)
	}

	// Endpoint domain.
	var createDocumentEndpoint endpoint.Endpoint
	{
		createDocumentEndpoint = library.MakeCreateDocumentEndpoint(service)
		createDocumentEndpoint = opentracing.TraceServer(tracer, "CreateDocument")(createDocumentEndpoint)
	}
	var findDocumentsEndpoint endpoint.Endpoint
	{
		findDocumentsEndpoint = library.MakeFindDocumentsEndpoint(service)
		findDocumentsEndpoint = TraceTransportBoundaries(findDocumentsEndpoint)
		// findDocumentsEndpoint = opentracing.TraceServer(tracer, "FindDocuments")(findDocumentsEndpoint)
	}
	var findDocumentsByIDEndpoint endpoint.Endpoint
	{
		findDocumentsByIDEndpoint = library.MakeFindDocumentsByIDEndpoint(service)
		findDocumentsByIDEndpoint = opentracing.TraceServer(tracer, "FindDocumentsByID")(findDocumentsByIDEndpoint)
	}
	var replaceDocumentByIDEndpoint endpoint.Endpoint
	{
		replaceDocumentByIDEndpoint = library.MakeReplaceDocumentByIDEndpoint(service)
		replaceDocumentByIDEndpoint = opentracing.TraceServer(tracer, "ReplaceDocumentByID")(replaceDocumentByIDEndpoint)
	}
	var deleteDocumentsByIDEndpoint endpoint.Endpoint
	{
		deleteDocumentsByIDEndpoint = library.MakeDeleteDocumentsByIDEndpoint(service)
		deleteDocumentsByIDEndpoint = opentracing.TraceServer(tracer, "DeleteDocumentsByID")(deleteDocumentsByIDEndpoint)
	}
	endpoints := library.Endpoints{
		CreateDocumentEndpoint:      createDocumentEndpoint,
		FindDocumentsEndpoint:       findDocumentsEndpoint,
		FindDocumentsByIDEndpoint:   findDocumentsByIDEndpoint,
		ReplaceDocumentByIDEndpoint: replaceDocumentByIDEndpoint,
		DeleteDocumentsByIDEndpoint: deleteDocumentsByIDEndpoint,
	}

	// Transport domain.
	ctx := context.Background()

	ln, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		logger.Log("err", err)
		os.Exit(1)
	}

	srv := library.MakeGRPCServer(ctx, endpoints, tracer, logger)
	s := grpc.NewServer()
	pb.RegisterLibraryServer(s, srv)

	logger.Log("msg", "listening on "+*grpcAddr+" (gRPC)")

	if err := s.Serve(ln); err != nil {
		logger.Log("err", err)
		os.Exit(1)
	}
}

func TraceTransportBoundaries(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		span := stdopentracing.SpanFromContext(ctx)
		if span != nil {
			span.LogEvent("Transport domain ends")
			ctx = stdopentracing.ContextWithSpan(ctx, span)
			defer func() {
				span.LogEvent("Transport domain begins")
			}()
		}
		return next(ctx, request)
	}
}
