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
	"google.golang.org/grpc/metadata"
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
			append(
				opts,
				grpctransport.ClientBefore(
					addGRPCAnnotations(target),
					opentracing.ToGRPCRequest(tracer, logger),
				),
			)...,
		).Endpoint()
		createDocumentEndpoint = EndpointTracingMiddleware(createDocumentEndpoint)
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
			append(
				opts,
				grpctransport.ClientBefore(
					addGRPCAnnotations(target),
					opentracing.ToGRPCRequest(tracer, logger),
				),
			)...,
		).Endpoint()
		findDocumentsEndpoint = EndpointTracingMiddleware(findDocumentsEndpoint)
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
			append(
				opts,
				grpctransport.ClientBefore(
					addGRPCAnnotations(target),
					opentracing.ToGRPCRequest(tracer, logger),
				),
			)...,
		).Endpoint()
		findDocumentsByIDEndpoint = EndpointTracingMiddleware(findDocumentsByIDEndpoint)
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
			append(
				opts,
				grpctransport.ClientBefore(
					addGRPCAnnotations(target),
					opentracing.ToGRPCRequest(tracer, logger),
				),
			)...,
		).Endpoint()
		replaceDocumentByIDEndpoint = EndpointTracingMiddleware(replaceDocumentByIDEndpoint)
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
			append(
				opts,
				grpctransport.ClientBefore(
					addGRPCAnnotations(target),
					opentracing.ToGRPCRequest(tracer, logger),
				),
			)...,
		).Endpoint()
		deleteDocumentsByIDEndpoint = EndpointTracingMiddleware(deleteDocumentsByIDEndpoint)
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

func addGRPCAnnotations(target string) func(ctx context.Context, md *metadata.MD) context.Context {
	return func(ctx context.Context, md *metadata.MD) context.Context {
		if span := stdopentracing.SpanFromContext(ctx); span != nil {
			span = span.SetTag("transport", "gRPC")
			span = span.SetTag("target", target)
			ctx = stdopentracing.ContextWithSpan(ctx, span)
		}
		return ctx
	}
}
