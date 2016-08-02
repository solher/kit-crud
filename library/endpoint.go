package library

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

type errorer interface {
	error() error
}

type Endpoints struct {
	CreateDocumentEndpoint      endpoint.Endpoint
	FindDocumentsEndpoint       endpoint.Endpoint
	FindDocumentsByIDEndpoint   endpoint.Endpoint
	ReplaceDocumentByIDEndpoint endpoint.Endpoint
	DeleteDocumentsByIDEndpoint endpoint.Endpoint
}

func (e Endpoints) CreateDocument(ctx context.Context, userID string, document *Document) (*Document, error) {
	req := createDocumentRequest{
		UserID:   userID,
		Document: document,
	}
	res, err := e.CreateDocumentEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(createDocumentResponse).Document,
		res.(createDocumentResponse).Err
}

func (e Endpoints) FindDocuments(ctx context.Context, userID string) ([]Document, error) {
	req := findDocumentsRequest{
		UserID: userID,
	}
	res, err := e.FindDocumentsEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(findDocumentsResponse).Documents,
		res.(findDocumentsResponse).Err
}

func (e Endpoints) FindDocumentsByID(ctx context.Context, userID string, ids []string) ([]Document, error) {
	req := findDocumentsByIDRequest{
		UserID: userID,
		IDs:    ids,
	}
	res, err := e.FindDocumentsByIDEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(findDocumentsByIDResponse).Documents,
		res.(findDocumentsByIDResponse).Err
}

func (e Endpoints) ReplaceDocumentByID(ctx context.Context, userID string, id string, document *Document) (*Document, error) {
	req := replaceDocumentByIDRequest{
		UserID:   userID,
		ID:       id,
		Document: document,
	}
	res, err := e.ReplaceDocumentByIDEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(replaceDocumentByIDResponse).Document,
		res.(replaceDocumentByIDResponse).Err
}

func (e Endpoints) DeleteDocumentsByID(ctx context.Context, userID string, ids []string) ([]Document, error) {
	req := deleteDocumentsByIDRequest{
		UserID: userID,
		IDs:    ids,
	}
	res, err := e.DeleteDocumentsByIDEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(deleteDocumentsByIDResponse).Documents,
		res.(deleteDocumentsByIDResponse).Err
}

type createDocumentRequest struct {
	UserID   string
	Document *Document
}

type createDocumentResponse struct {
	Document *Document
	Err      error `json:"err"`
}

func (r *createDocumentResponse) error() error { return r.Err }

func MakeCreateDocumentEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createDocumentRequest)
		document, err := s.CreateDocument(ctx, req.UserID, req.Document)
		return createDocumentResponse{Document: document, Err: err}, nil
	}
}

type findDocumentsRequest struct {
	UserID string
}

type findDocumentsResponse struct {
	Documents []Document `json:"documents"`
	Err       error      `json:"err"`
}

func (r *findDocumentsResponse) error() error { return r.Err }

func MakeFindDocumentsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(findDocumentsRequest)
		documents, err := s.FindDocuments(ctx, req.UserID)
		return findDocumentsResponse{Documents: documents, Err: err}, nil
	}
}

type findDocumentsByIDRequest struct {
	UserID string
	IDs    []string
}

type findDocumentsByIDResponse struct {
	Documents []Document `json:"documents"`
	Err       error      `json:"err"`
}

func (r *findDocumentsByIDResponse) error() error { return r.Err }

func MakeFindDocumentsByIDEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(findDocumentsByIDRequest)
		documents, err := s.FindDocumentsByID(ctx, req.UserID, req.IDs)
		return findDocumentsByIDResponse{Documents: documents, Err: err}, nil
	}
}

type replaceDocumentByIDRequest struct {
	UserID   string
	ID       string
	Document *Document
}

type replaceDocumentByIDResponse struct {
	Document *Document `json:"document"`
	Err      error     `json:"err"`
}

func (r *replaceDocumentByIDResponse) error() error { return r.Err }

func MakeReplaceDocumentByIDEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(replaceDocumentByIDRequest)
		document, err := s.ReplaceDocumentByID(ctx, req.UserID, req.ID, req.Document)
		return replaceDocumentByIDResponse{Document: document, Err: err}, nil
	}
}

type deleteDocumentsByIDRequest struct {
	UserID string
	IDs    []string
}

type deleteDocumentsByIDResponse struct {
	Documents []Document `json:"documents"`
	Err       error      `json:"err"`
}

func (r *deleteDocumentsByIDResponse) error() error { return r.Err }

func MakeDeleteDocumentsByIDEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteDocumentsByIDRequest)
		documents, err := s.DeleteDocumentsByID(ctx, req.UserID, req.IDs)
		return deleteDocumentsByIDResponse{Documents: documents, Err: err}, nil
	}
}
