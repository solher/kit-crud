package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/solher/kit-crud/library"

	"sourcegraph.com/sourcegraph/appdash"
	appdashot "sourcegraph.com/sourcegraph/appdash/opentracing"

	"github.com/solher/kit-crud/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var (
		grpcAddr    = flag.String("grpc.addr", ":8082", "gRPC (HTTP) listen address")
		appdashAddr = flag.String("appdash.addr", "", "Enable Appdash tracing via server host:port")
	)
	flag.Parse()

	exitCode := 0
	defer func() {
		os.Exit(exitCode)
	}()

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
		if *appdashAddr != "" {
			logger := log.NewContext(logger).With("tracer", "Appdash")
			logger.Log("msg", "sending trace to "+*appdashAddr)
			tracer = appdashot.NewTracer(appdash.NewRemoteCollector(*appdashAddr))

		} else {
			logger := log.NewContext(logger).With("tracer", "none")
			logger.Log("msg", "tracing disabled")
			tracer = stdopentracing.GlobalTracer() // no-op
		}
	}

	// Business domain.
	var service library.Service
	{
		store := library.NewRepository(tracer)
		service = library.NewService(store)
		service = library.ServiceTracingMiddleware(service)
	}

	// Endpoint domain.
	var createDocumentEndpoint endpoint.Endpoint
	{
		createDocumentEndpoint = library.MakeCreateDocumentEndpoint(service)
		createDocumentEndpoint = EndpointTracingMiddleware(createDocumentEndpoint)
	}
	var findDocumentsEndpoint endpoint.Endpoint
	{
		findDocumentsEndpoint = library.MakeFindDocumentsEndpoint(service)
		findDocumentsEndpoint = EndpointTracingMiddleware(findDocumentsEndpoint)
	}
	var findDocumentsByIDEndpoint endpoint.Endpoint
	{
		findDocumentsByIDEndpoint = library.MakeFindDocumentsByIDEndpoint(service)
		findDocumentsByIDEndpoint = EndpointTracingMiddleware(findDocumentsByIDEndpoint)
	}
	var replaceDocumentByIDEndpoint endpoint.Endpoint
	{
		replaceDocumentByIDEndpoint = library.MakeReplaceDocumentByIDEndpoint(service)
		replaceDocumentByIDEndpoint = EndpointTracingMiddleware(replaceDocumentByIDEndpoint)
	}
	var deleteDocumentsByIDEndpoint endpoint.Endpoint
	{
		deleteDocumentsByIDEndpoint = library.MakeDeleteDocumentsByIDEndpoint(service)
		deleteDocumentsByIDEndpoint = EndpointTracingMiddleware(deleteDocumentsByIDEndpoint)
	}
	endpoints := library.Endpoints{
		CreateDocumentEndpoint:      createDocumentEndpoint,
		FindDocumentsEndpoint:       findDocumentsEndpoint,
		FindDocumentsByIDEndpoint:   findDocumentsByIDEndpoint,
		ReplaceDocumentByIDEndpoint: replaceDocumentByIDEndpoint,
		DeleteDocumentsByIDEndpoint: deleteDocumentsByIDEndpoint,
	}

	// Mechanical domain.
	ctx := context.Background()
	errc := make(chan error)

	// Transport domain.
	s := library.MakeGRPCServer(ctx, endpoints, tracer, logger)
	server := grpc.NewServer()
	pb.RegisterLibraryServer(server, s)

	conn, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		logger.Log("err", err)
		exitCode = 1
		return
	}
	defer conn.Close()
	logger.Log("msg", "listening on "+*grpcAddr+" (gRPC)")
	go func() {
		if err := server.Serve(conn); err != nil {
			errc <- err
			return
		}
	}()

	// Interrupt handler.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		logger.Log(
			"signal", fmt.Sprintf("%s", <-c),
			"msg", "gracefully shutting down",
		)
		errc <- nil
	}()

	if err := <-errc; err != nil {
		logger.Log("err", err)
		exitCode = 1
	}
}
