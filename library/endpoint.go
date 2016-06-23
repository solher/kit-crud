package library

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

type createDocumentRequest struct {
	UserID   string
	Document *Document
}

type createDocumentResponse struct {
	Document *Document
	Err      error `json:"err"`
}

func (r createDocumentResponse) error() error { return r.Err }

func makeCreateDocumentEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(createDocumentRequest)
		document, err := s.CreateDocument(req.UserID, req.Document)
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

func (r findDocumentsResponse) error() error { return r.Err }

func makeFindDocumentsEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(findDocumentsRequest)
		documents, err := s.FindDocuments(req.UserID)
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

func (r findDocumentsByIDResponse) error() error { return r.Err }

func makeFindDocumentsByIDEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(findDocumentsByIDRequest)
		documents, err := s.FindDocumentsByID(req.UserID, req.IDs)
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

func (r replaceDocumentByIDResponse) error() error { return r.Err }

func makeReplaceDocumentByIDEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(replaceDocumentByIDRequest)
		document, err := s.ReplaceDocumentByID(req.UserID, req.ID, req.Document)
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

func (r deleteDocumentsByIDResponse) error() error { return r.Err }

func makeDeleteDocumentsByIDEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteDocumentsByIDRequest)
		documents, err := s.DeleteDocumentsByID(req.UserID, req.IDs)
		return deleteDocumentsByIDResponse{Documents: documents, Err: err}, nil
	}
}
