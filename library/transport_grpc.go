package library

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/solher/kit-crud/pb"
	"golang.org/x/net/context"
)

func MakeGRPCServer(ctx context.Context, endpoints Endpoints) pb.LibraryServer {
	opts := []grpctransport.ServerOption{}
	return &grpcServer{
		createDocument: grpctransport.NewServer(
			ctx,
			endpoints.CreateDocumentEndpoint,
			DecodeGRPCCreateDocumentRequest,
			EncodeGRPCCreateDocumentResponse,
			opts...,
		),
		findDocuments: grpctransport.NewServer(
			ctx,
			endpoints.FindDocumentsEndpoint,
			DecodeGRPCFindDocumentsRequest,
			EncodeGRPCFindDocumentsResponse,
			opts...,
		),
		findDocumentsByID: grpctransport.NewServer(
			ctx,
			endpoints.FindDocumentsByIDEndpoint,
			DecodeGRPCFindDocumentsByIDRequest,
			EncodeGRPCFindDocumentsByIDResponse,
			opts...,
		),
		replaceDocumentByID: grpctransport.NewServer(
			ctx,
			endpoints.ReplaceDocumentByIDEndpoint,
			DecodeGRPCReplaceDocumentByIDRequest,
			EncodeGRPCReplaceDocumentByIDResponse,
			opts...,
		),
		deleteDocumentsByID: grpctransport.NewServer(
			ctx,
			endpoints.DeleteDocumentsByIDEndpoint,
			DecodeGRPCDeleteDocumentsByIDRequest,
			EncodeGRPCDeleteDocumentsByIDResponse,
			opts...,
		),
	}
}

type grpcServer struct {
	createDocument      grpctransport.Handler
	findDocuments       grpctransport.Handler
	findDocumentsByID   grpctransport.Handler
	replaceDocumentByID grpctransport.Handler
	deleteDocumentsByID grpctransport.Handler
}

func (s *grpcServer) CreateDocument(ctx context.Context, req *pb.CreateDocumentRequest) (*pb.CreateDocumentReply, error) {
	_, rep, err := s.createDocument.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateDocumentReply), nil
}

func (s *grpcServer) FindDocuments(ctx context.Context, req *pb.FindDocumentsRequest) (*pb.FindDocumentsReply, error) {
	_, rep, err := s.findDocuments.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.FindDocumentsReply), nil
}

func (s *grpcServer) FindDocumentsById(ctx context.Context, req *pb.FindDocumentsByIdRequest) (*pb.FindDocumentsByIdReply, error) {
	_, rep, err := s.findDocumentsByID.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.FindDocumentsByIdReply), nil
}

func (s *grpcServer) ReplaceDocumentById(ctx context.Context, req *pb.ReplaceDocumentByIdRequest) (*pb.ReplaceDocumentByIdReply, error) {
	_, rep, err := s.replaceDocumentByID.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ReplaceDocumentByIdReply), nil
}

func (s *grpcServer) DeleteDocumentsById(ctx context.Context, req *pb.DeleteDocumentsByIdRequest) (*pb.DeleteDocumentsByIdReply, error) {
	_, rep, err := s.deleteDocumentsByID.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteDocumentsByIdReply), nil
}

func DecodeGRPCCreateDocumentRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateDocumentRequest)
	return createDocumentRequest{
		UserID:   req.UserId,
		Document: toDocument(req.Document),
	}, nil
}

func EncodeGRPCCreateDocumentResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(createDocumentResponse)
	return &pb.CreateDocumentReply{
		Document: toPBDocument(res.Document),
		Err:      toPBError(res.Err),
	}, nil
}

func DecodeGRPCFindDocumentsRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.FindDocumentsRequest)
	return findDocumentsRequest{
		UserID: req.UserId,
	}, nil
}

func EncodeGRPCFindDocumentsResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(findDocumentsResponse)
	return &pb.FindDocumentsReply{
		Documents: toPBDocuments(res.Documents),
		Err:       toPBError(res.Err),
	}, nil
}

func DecodeGRPCFindDocumentsByIDRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.FindDocumentsByIdRequest)
	return findDocumentsByIDRequest{
		UserID: req.UserId,
		IDs:    req.Ids,
	}, nil
}

func EncodeGRPCFindDocumentsByIDResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(findDocumentsByIDResponse)
	return &pb.FindDocumentsByIdReply{
		Documents: toPBDocuments(res.Documents),
		Err:       toPBError(res.Err),
	}, nil
}

func DecodeGRPCReplaceDocumentByIDRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReplaceDocumentByIdRequest)
	return replaceDocumentByIDRequest{
		UserID:   req.UserId,
		ID:       req.Id,
		Document: toDocument(req.Document),
	}, nil
}

func EncodeGRPCReplaceDocumentByIDResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(replaceDocumentByIDResponse)
	return &pb.ReplaceDocumentByIdReply{
		Document: toPBDocument(res.Document),
		Err:      toPBError(res.Err),
	}, nil
}

func DecodeGRPCDeleteDocumentsByIDRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.DeleteDocumentsByIdRequest)
	return deleteDocumentsByIDRequest{
		UserID: req.UserId,
		IDs:    req.Ids,
	}, nil
}

func EncodeGRPCDeleteDocumentsByIDResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(deleteDocumentsByIDResponse)
	return &pb.DeleteDocumentsByIdReply{
		Documents: toPBDocuments(res.Documents),
		Err:       toPBError(res.Err),
	}, nil
}

func toPBError(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

func toPBDocument(document *Document) *pb.Document {
	return toPBDocuments([]Document{*document})[0]
}

func toPBDocuments(documents []Document) []*pb.Document {
	pbDocuments := make([]*pb.Document, len(documents))
	for i, n := range documents {
		pbDocuments[i] = &pb.Document{
			Id:      n.ID,
			UserId:  n.UserID,
			Content: n.Content,
		}
	}
	return pbDocuments
}

func toDocument(pbDocument *pb.Document) *Document {
	return &toDocuments([]*pb.Document{pbDocument})[0]
}

func toDocuments(pbDocuments []*pb.Document) []Document {
	documents := make([]Document, len(pbDocuments))
	for i, n := range pbDocuments {
		documents[i] = Document{
			ID:      n.Id,
			UserID:  n.UserId,
			Content: n.Content,
		}
	}
	return documents
}
