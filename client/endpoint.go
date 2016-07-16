package client

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

type Endpoints struct {
	CreateDocumentEndpoint      endpoint.Endpoint
	FindDocumentsEndpoint       endpoint.Endpoint
	FindDocumentsByIDEndpoint   endpoint.Endpoint
	ReplaceDocumentByIDEndpoint endpoint.Endpoint
	DeleteDocumentsByIDEndpoint endpoint.Endpoint
}

func (e Endpoints) CreateDocument(userID string, document *Document) (*Document, error) {
	req := CreateDocumentRequest{
		UserID:   userID,
		Document: document,
	}
	res, err := e.CreateDocumentEndpoint(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return res.(CreateDocumentResponse).Document,
		res.(CreateDocumentResponse).Err
}

func (e Endpoints) FindDocuments(userID string) ([]Document, error) {
	req := FindDocumentsRequest{
		UserID: userID,
	}
	res, err := e.FindDocumentsEndpoint(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return res.(FindDocumentsResponse).Documents,
		res.(FindDocumentsResponse).Err
}

func (e Endpoints) FindDocumentsByID(userID string, ids []string) ([]Document, error) {
	req := FindDocumentsByIDRequest{
		UserID: userID,
		IDs:    ids,
	}
	res, err := e.FindDocumentsByIDEndpoint(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return res.(FindDocumentsByIDResponse).Documents,
		res.(FindDocumentsByIDResponse).Err
}

func (e Endpoints) ReplaceDocumentByID(userID string, id string, document *Document) (*Document, error) {
	req := ReplaceDocumentByIDRequest{
		UserID:   userID,
		ID:       id,
		Document: document,
	}
	res, err := e.ReplaceDocumentByIDEndpoint(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return res.(ReplaceDocumentByIDResponse).Document,
		res.(ReplaceDocumentByIDResponse).Err
}

func (e Endpoints) DeleteDocumentsByID(userID string, ids []string) ([]Document, error) {
	req := DeleteDocumentsByIDRequest{
		UserID: userID,
		IDs:    ids,
	}
	res, err := e.DeleteDocumentsByIDEndpoint(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return res.(DeleteDocumentsByIDResponse).Documents,
		res.(DeleteDocumentsByIDResponse).Err
}

type CreateDocumentRequest struct {
	UserID   string
	Document *Document
}

type CreateDocumentResponse struct {
	Document *Document
	Err      error `json:"err"`
}

func MakeCreateDocumentEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateDocumentRequest)
		document, err := s.CreateDocument(req.UserID, req.Document)
		return CreateDocumentResponse{Document: document, Err: err}, nil
	}
}

type FindDocumentsRequest struct {
	UserID string
}

type FindDocumentsResponse struct {
	Documents []Document `json:"documents"`
	Err       error      `json:"err"`
}

func MakeFindDocumentsEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(FindDocumentsRequest)
		documents, err := s.FindDocuments(req.UserID)
		return FindDocumentsResponse{Documents: documents, Err: err}, nil
	}
}

type FindDocumentsByIDRequest struct {
	UserID string
	IDs    []string
}

type FindDocumentsByIDResponse struct {
	Documents []Document `json:"documents"`
	Err       error      `json:"err"`
}

func MakeFindDocumentsByIDEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(FindDocumentsByIDRequest)
		documents, err := s.FindDocumentsByID(req.UserID, req.IDs)
		return FindDocumentsByIDResponse{Documents: documents, Err: err}, nil
	}
}

type ReplaceDocumentByIDRequest struct {
	UserID   string
	ID       string
	Document *Document
}

type ReplaceDocumentByIDResponse struct {
	Document *Document `json:"document"`
	Err      error     `json:"err"`
}

func MakeReplaceDocumentByIDEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(ReplaceDocumentByIDRequest)
		document, err := s.ReplaceDocumentByID(req.UserID, req.ID, req.Document)
		return ReplaceDocumentByIDResponse{Document: document, Err: err}, nil
	}
}

type DeleteDocumentsByIDRequest struct {
	UserID string
	IDs    []string
}

type DeleteDocumentsByIDResponse struct {
	Documents []Document `json:"documents"`
	Err       error      `json:"err"`
}

func MakeDeleteDocumentsByIDEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteDocumentsByIDRequest)
		documents, err := s.DeleteDocumentsByID(req.UserID, req.IDs)
		return DeleteDocumentsByIDResponse{Documents: documents, Err: err}, nil
	}
}
