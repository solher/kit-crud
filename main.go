package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/go-kit/kit/endpoint"
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

	// Business domain.
	var service library.Service
	{
		service = library.NewService()
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
		fmt.Println(err)
		os.Exit(-1)
	}

	srv := library.MakeGRPCServer(ctx, endpoints)
	s := grpc.NewServer()
	pb.RegisterLibraryServer(s, srv)

	if err := s.Serve(ln); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
