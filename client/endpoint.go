package client

import (
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/solher/kit-crud/pb"
	"golang.org/x/net/context"
)

type Endpoints struct {
	CreateDocumentEndpoint      endpoint.Endpoint
	FindDocumentsEndpoint       endpoint.Endpoint
	FindDocumentsByIDEndpoint   endpoint.Endpoint
	ReplaceDocumentByIDEndpoint endpoint.Endpoint
	DeleteDocumentsByIDEndpoint endpoint.Endpoint
}

func (e Endpoints) CreateDocument(ctx context.Context, userID string, document *pb.Document) (*pb.Document, error) {
	req := &pb.CreateDocumentRequest{
		UserId:   userID,
		Document: document,
	}
	response, err := e.CreateDocumentEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	res := response.(*pb.CreateDocumentReply)
	return res.Document, toError(res.Err)
}

func (e Endpoints) FindDocuments(ctx context.Context, userID string) ([]*pb.Document, error) {
	req := &pb.FindDocumentsRequest{
		UserId: userID,
	}
	response, err := e.FindDocumentsEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	res := response.(*pb.FindDocumentsReply)
	return res.Documents, toError(res.Err)
}

func (e Endpoints) FindDocumentsByID(ctx context.Context, userID string, ids []string) ([]*pb.Document, error) {
	req := &pb.FindDocumentsByIdRequest{
		UserId: userID,
		Ids:    ids,
	}
	response, err := e.FindDocumentsByIDEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	res := response.(*pb.FindDocumentsByIdReply)
	return res.Documents, toError(res.Err)
}

func (e Endpoints) ReplaceDocumentByID(ctx context.Context, userID string, id string, document *pb.Document) (*pb.Document, error) {
	req := &pb.ReplaceDocumentByIdRequest{
		UserId:   userID,
		Id:       id,
		Document: document,
	}
	response, err := e.ReplaceDocumentByIDEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	res := response.(*pb.ReplaceDocumentByIdReply)
	return res.Document, toError(res.Err)
}

func (e Endpoints) DeleteDocumentsByID(ctx context.Context, userID string, ids []string) ([]*pb.Document, error) {
	req := &pb.DeleteDocumentsByIdRequest{
		UserId: userID,
		Ids:    ids,
	}
	response, err := e.DeleteDocumentsByIDEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	res := response.(*pb.DeleteDocumentsByIdReply)
	return res.Documents, toError(res.Err)
}

func MakeCreateDocumentEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.CreateDocumentRequest)
		document, err := s.CreateDocument(ctx, req.UserId, req.Document)
		return &pb.CreateDocumentReply{
			Document: document,
			Err:      toPBError(err),
		}, nil
	}
}

func MakeFindDocumentsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.FindDocumentsRequest)
		documents, err := s.FindDocuments(ctx, req.UserId)
		return &pb.FindDocumentsReply{
			Documents: documents,
			Err:       toPBError(err),
		}, nil
	}
}

func MakeFindDocumentsByIDEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.FindDocumentsByIdRequest)
		documents, err := s.FindDocumentsByID(ctx, req.UserId, req.Ids)
		return &pb.FindDocumentsByIdReply{
			Documents: documents,
			Err:       toPBError(err),
		}, nil
	}
}

func MakeReplaceDocumentByIDEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.ReplaceDocumentByIdRequest)
		document, err := s.ReplaceDocumentByID(ctx, req.UserId, req.Id, req.Document)
		return &pb.ReplaceDocumentByIdReply{
			Document: document,
			Err:      toPBError(err),
		}, nil
	}
}

func MakeDeleteDocumentsByIDEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.DeleteDocumentsByIdRequest)
		documents, err := s.DeleteDocumentsByID(ctx, req.UserId, req.Ids)
		return &pb.DeleteDocumentsByIdReply{
			Documents: documents,
			Err:       toPBError(err),
		}, nil
	}
}

func toError(err string) error {
	if len(err) == 0 {
		return nil
	}
	return errors.New(err)
}

func toPBError(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
