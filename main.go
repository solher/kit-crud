package main

import (
	"flag"
	"net"
	"os"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/solher/kit-crud/library"
	"github.com/solher/kit-crud/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var (
		grpcAddr = flag.String("grpc.addr", ":8082", "gRPC (HTTP) listen address")
	)
	flag.Parse()

	// Logging domain.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)
		logger = log.NewContext(logger).With("caller", log.DefaultCaller)
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
	}
	var findDocumentsEndpoint endpoint.Endpoint
	{
		findDocumentsEndpoint = library.MakeFindDocumentsEndpoint(service)
	}
	var findDocumentsByIDEndpoint endpoint.Endpoint
	{
		findDocumentsByIDEndpoint = library.MakeFindDocumentsByIDEndpoint(service)
	}
	var replaceDocumentByIDEndpoint endpoint.Endpoint
	{
		replaceDocumentByIDEndpoint = library.MakeReplaceDocumentByIDEndpoint(service)
	}
	var deleteDocumentsByIDEndpoint endpoint.Endpoint
	{
		deleteDocumentsByIDEndpoint = library.MakeDeleteDocumentsByIDEndpoint(service)
	}
	endpoints := library.Endpoints{
		CreateDocumentEndpoint:      createDocumentEndpoint,
		FindDocumentsEndpoint:       findDocumentsEndpoint,
		FindDocumentsByIDEndpoint:   findDocumentsByIDEndpoint,
		ReplaceDocumentByIDEndpoint: replaceDocumentByIDEndpoint,
		DeleteDocumentsByIDEndpoint: deleteDocumentsByIDEndpoint,
	}

	ctx := context.Background()

	ln, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		logger.Log("err", err)
		os.Exit(1)
	}

	srv := library.MakeGRPCServer(ctx, endpoints)
	s := grpc.NewServer()
	pb.RegisterLibraryServer(s, srv)

	logger.Log("msg", "listening on "+*grpcAddr+" (gRPC)")

	if err := s.Serve(ln); err != nil {
		logger.Log("err", err)
		os.Exit(1)
	}
}
