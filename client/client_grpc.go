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
		service := grpcClient(instance, conn, tracer, logger)
		endpoint := makeEndpoint(service)

		return endpoint, conn, nil
	}
}

func grpcClient(target string, conn *grpc.ClientConn, tracer stdopentracing.Tracer, logger log.Logger) Service {
	opts := []grpctransport.ClientOption{}

	var createDocumentEndpoint endpoint.Endpoint
	{
		createDocumentEndpoint = grpctransport.NewClient(
			conn,
			"Library",
			"CreateDocument",
			nopCodec,
			nopCodec,
			pb.CreateDocumentReply{},
			append(opts, grpctransport.ClientBefore(opentracing.ToGRPCRequest(tracer, logger)))...,
		).Endpoint()
		createDocumentEndpoint = AddGRPCClientAnnotations(target)(createDocumentEndpoint)
		createDocumentEndpoint = opentracing.TraceClient(tracer, "CreateDocument")(createDocumentEndpoint)
	}
	var findDocumentsEndpoint endpoint.Endpoint
	{
		findDocumentsEndpoint = grpctransport.NewClient(
			conn,
			"Library",
			"FindDocuments",
			nopCodec,
			nopCodec,
			pb.FindDocumentsReply{},
			append(opts, grpctransport.ClientBefore(opentracing.ToGRPCRequest(tracer, logger)))...,
		).Endpoint()
		findDocumentsEndpoint = AddGRPCClientAnnotations(target)(findDocumentsEndpoint)
		findDocumentsEndpoint = opentracing.TraceClient(tracer, "FindDocuments")(findDocumentsEndpoint)
	}
	var findDocumentsByIDEndpoint endpoint.Endpoint
	{
		findDocumentsByIDEndpoint = grpctransport.NewClient(
			conn,
			"Library",
			"FindDocumentsById",
			nopCodec,
			nopCodec,
			pb.FindDocumentsByIdReply{},
			append(opts, grpctransport.ClientBefore(opentracing.ToGRPCRequest(tracer, logger)))...,
		).Endpoint()
		findDocumentsByIDEndpoint = AddGRPCClientAnnotations(target)(findDocumentsByIDEndpoint)
		findDocumentsByIDEndpoint = opentracing.TraceClient(tracer, "FindDocumentsById")(findDocumentsByIDEndpoint)
	}
	var replaceDocumentByIDEndpoint endpoint.Endpoint
	{
		replaceDocumentByIDEndpoint = grpctransport.NewClient(
			conn,
			"Library",
			"ReplaceDocumentById",
			nopCodec,
			nopCodec,
			pb.ReplaceDocumentByIdReply{},
			append(opts, grpctransport.ClientBefore(opentracing.ToGRPCRequest(tracer, logger)))...,
		).Endpoint()
		replaceDocumentByIDEndpoint = AddGRPCClientAnnotations(target)(replaceDocumentByIDEndpoint)
		replaceDocumentByIDEndpoint = opentracing.TraceClient(tracer, "ReplaceDocumentById")(replaceDocumentByIDEndpoint)
	}
	var deleteDocumentsByIDEndpoint endpoint.Endpoint
	{
		deleteDocumentsByIDEndpoint = grpctransport.NewClient(
			conn,
			"Library",
			"DeleteDocumentsById",
			nopCodec,
			nopCodec,
			pb.DeleteDocumentsByIdReply{},
			append(opts, grpctransport.ClientBefore(opentracing.ToGRPCRequest(tracer, logger)))...,
		).Endpoint()
		deleteDocumentsByIDEndpoint = AddGRPCClientAnnotations(target)(deleteDocumentsByIDEndpoint)
		deleteDocumentsByIDEndpoint = opentracing.TraceClient(tracer, "DeleteDocumentsById")(deleteDocumentsByIDEndpoint)
	}

	return Endpoints{
		CreateDocumentEndpoint:      createDocumentEndpoint,
		FindDocumentsEndpoint:       findDocumentsEndpoint,
		FindDocumentsByIDEndpoint:   findDocumentsByIDEndpoint,
		ReplaceDocumentByIDEndpoint: replaceDocumentByIDEndpoint,
		DeleteDocumentsByIDEndpoint: deleteDocumentsByIDEndpoint,
	}
}

func AddGRPCClientAnnotations(target string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			span := stdopentracing.SpanFromContext(ctx)
			if span != nil {
				span = span.SetTag("transport", "gRPC")
				span = span.SetTag("target", target)
				ctx = stdopentracing.ContextWithSpan(ctx, span)
			}
			response, err := next(ctx, request)
			if err != nil && span != nil {
				span = span.SetTag("error", err)
				ctx = stdopentracing.ContextWithSpan(ctx, span)
				return response, err
			}
			err = errorer(response)
			if err != nil && span != nil {
				span = span.SetTag("error", err)
				ctx = stdopentracing.ContextWithSpan(ctx, span)
				return response, nil
			}
			return response, nil
		}
	}
}

func errorer(response interface{}) error {
	var str string
	switch res := response.(type) {
	case *pb.CreateDocumentReply:
		str = res.Err
	case *pb.FindDocumentsReply:
		str = res.Err
	case *pb.FindDocumentsByIdReply:
		str = res.Err
	case *pb.ReplaceDocumentByIdReply:
		str = res.Err
	case *pb.DeleteDocumentsByIdReply:
		str = res.Err
	default:
		str = "unexpected response type"
	}
	return toError(str)
}
