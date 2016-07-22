package client

import (
	"io"

	"github.com/go-kit/kit/log"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/solher/kit-crud/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var nopCodec = func(_ context.Context, r interface{}) (interface{}, error) {
	return r, nil
}

func NewGRPC(consulAddr string, tracer stdopentracing.Tracer, logger log.Logger) (Service, error) {
	var (
		endpoints = Endpoints{}
	)
	{
		factory := grpcFactory(MakeCreateDocumentEndpoint, tracer, logger)
		endpoint, _, err := factory(consulAddr)
		// defer conn.Close()
		if err != nil {
			return nil, err
		}
		endpoints.CreateDocumentEndpoint = endpoint
	}
	{
		factory := grpcFactory(MakeFindDocumentsEndpoint, tracer, logger)
		endpoint, _, err := factory(consulAddr)
		// defer conn.Close()
		if err != nil {
			return nil, err
		}
		endpoints.FindDocumentsEndpoint = endpoint
	}
	{
		factory := grpcFactory(MakeFindDocumentsByIDEndpoint, tracer, logger)
		endpoint, _, err := factory(consulAddr)
		// defer conn.Close()
		if err != nil {
			return nil, err
		}
		endpoints.FindDocumentsByIDEndpoint = endpoint
	}
	{
		factory := grpcFactory(MakeReplaceDocumentByIDEndpoint, tracer, logger)
		endpoint, _, err := factory(consulAddr)
		// defer conn.Close()
		if err != nil {
			return nil, err
		}
		endpoints.ReplaceDocumentByIDEndpoint = endpoint
	}
	{
		factory := grpcFactory(MakeDeleteDocumentsByIDEndpoint, tracer, logger)
		endpoint, _, err := factory(consulAddr)
		// defer conn.Close()
		if err != nil {
			return nil, err
		}
		endpoints.DeleteDocumentsByIDEndpoint = endpoint
	}

	return endpoints, nil
}

func grpcFactory(makeEndpoint func(Service) endpoint.Endpoint, tracer stdopentracing.Tracer, logger log.Logger) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}
		service := grpcClient(conn, tracer, logger)
		endpoint := makeEndpoint(service)

		return endpoint, conn, nil
	}
}

func grpcClient(conn *grpc.ClientConn, tracer stdopentracing.Tracer, logger log.Logger) Service {
	opts := []grpctransport.ClientOption{}
	e := Endpoints{}

	e.CreateDocumentEndpoint = grpctransport.NewClient(
		conn,
		"Library",
		"CreateDocument",
		nopCodec,
		nopCodec,
		pb.CreateDocumentReply{},
		append(opts, grpctransport.ClientBefore(opentracing.FromGRPCRequest(tracer, "CreateDocument", logger)))...,
	).Endpoint()

	e.FindDocumentsEndpoint = grpctransport.NewClient(
		conn,
		"Library",
		"FindDocuments",
		nopCodec,
		nopCodec,
		pb.FindDocumentsReply{},
		append(opts, grpctransport.ClientBefore(opentracing.FromGRPCRequest(tracer, "FindDocuments", logger)))...,
	).Endpoint()

	e.FindDocumentsByIDEndpoint = grpctransport.NewClient(
		conn,
		"Library",
		"FindDocumentsById",
		nopCodec,
		nopCodec,
		pb.FindDocumentsByIdReply{},
		append(opts, grpctransport.ClientBefore(opentracing.FromGRPCRequest(tracer, "FindDocumentsById", logger)))...,
	).Endpoint()

	e.ReplaceDocumentByIDEndpoint = grpctransport.NewClient(
		conn,
		"Library",
		"ReplaceDocumentById",
		nopCodec,
		nopCodec,
		pb.ReplaceDocumentByIdReply{},
		append(opts, grpctransport.ClientBefore(opentracing.FromGRPCRequest(tracer, "ReplaceDocumentById", logger)))...,
	).Endpoint()

	e.DeleteDocumentsByIDEndpoint = grpctransport.NewClient(
		conn,
		"Library",
		"DeleteDocumentsById",
		nopCodec,
		nopCodec,
		pb.DeleteDocumentsByIdReply{},
		append(opts, grpctransport.ClientBefore(opentracing.FromGRPCRequest(tracer, "DeleteDocumentsById", logger)))...,
	).Endpoint()

	return e
}
