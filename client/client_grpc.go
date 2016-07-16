package client

import (
	"io"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/solher/kit-crud/pb"
	"google.golang.org/grpc"
)

func NewGRPC(consulAddr string) (Service, error) {
	var (
		endpoints = Endpoints{}
	)
	{
		factory := grpcFactory(MakeCreateDocumentEndpoint)
		endpoint, _, err := factory("localhost:8082")
		// defer conn.Close()
		if err != nil {
			return nil, err
		}
		endpoints.CreateDocumentEndpoint = endpoint
	}
	{
		factory := grpcFactory(MakeFindDocumentsEndpoint)
		endpoint, _, err := factory("localhost:8082")
		// defer conn.Close()
		if err != nil {
			return nil, err
		}
		endpoints.FindDocumentsEndpoint = endpoint
	}
	{
		factory := grpcFactory(MakeFindDocumentsByIDEndpoint)
		endpoint, _, err := factory("localhost:8082")
		// defer conn.Close()
		if err != nil {
			return nil, err
		}
		endpoints.FindDocumentsByIDEndpoint = endpoint
	}
	{
		factory := grpcFactory(MakeReplaceDocumentByIDEndpoint)
		endpoint, _, err := factory("localhost:8082")
		// defer conn.Close()
		if err != nil {
			return nil, err
		}
		endpoints.ReplaceDocumentByIDEndpoint = endpoint
	}
	{
		factory := grpcFactory(MakeDeleteDocumentsByIDEndpoint)
		endpoint, _, err := factory("localhost:8082")
		// defer conn.Close()
		if err != nil {
			return nil, err
		}
		endpoints.DeleteDocumentsByIDEndpoint = endpoint
	}

	return endpoints, nil
}

func grpcFactory(makeEndpoint func(Service) endpoint.Endpoint) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}
		service := grpcClient(conn)
		endpoint := makeEndpoint(service)

		return endpoint, conn, nil
	}
}

func grpcClient(conn *grpc.ClientConn) Service {
	opts := []grpctransport.ClientOption{}
	e := Endpoints{}

	e.CreateDocumentEndpoint = grpctransport.NewClient(
		conn,
		"Library",
		"CreateDocument",
		encodeGRPCCreateDocumentRequest,
		decodeGRPCCreateDocumentResponse,
		pb.CreateDocumentReply{},
		opts...,
	).Endpoint()

	e.FindDocumentsEndpoint = grpctransport.NewClient(
		conn,
		"Library",
		"FindDocuments",
		encodeGRPCFindDocumentsRequest,
		decodeGRPCFindDocumentsResponse,
		pb.FindDocumentsReply{},
		opts...,
	).Endpoint()

	e.FindDocumentsByIDEndpoint = grpctransport.NewClient(
		conn,
		"Library",
		"FindDocumentsById",
		encodeGRPCFindDocumentsByIDRequest,
		decodeGRPCFindDocumentsByIDResponse,
		pb.FindDocumentsByIdReply{},
		opts...,
	).Endpoint()

	e.ReplaceDocumentByIDEndpoint = grpctransport.NewClient(
		conn,
		"Library",
		"ReplaceDocumentById",
		encodeGRPCReplaceDocumentByIDRequest,
		decodeGRPCReplaceDocumentByIDResponse,
		pb.ReplaceDocumentByIdReply{},
		opts...,
	).Endpoint()

	e.DeleteDocumentsByIDEndpoint = grpctransport.NewClient(
		conn,
		"Library",
		"DeleteDocumentsById",
		encodeGRPCDeleteDocumentsByIDRequest,
		decodeGRPCDeleteDocumentsByIDResponse,
		pb.DeleteDocumentsByIdReply{},
		opts...,
	).Endpoint()

	return e
}
