package library

import "golang.org/x/net/context"

type Service interface {
	CreateDocument(ctx context.Context, userID string, document *Document) (*Document, error)
	FindDocuments(ctx context.Context, userID string) ([]Document, error)
	FindDocumentsByID(ctx context.Context, userID string, ids []string) ([]Document, error)
	ReplaceDocumentByID(ctx context.Context, userID string, id string, document *Document) (*Document, error)
	DeleteDocumentsByID(ctx context.Context, userID string, ids []string) ([]Document, error)
}

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) CreateDocument(ctx context.Context, userID string, document *Document) (*Document, error) {
	return nil, nil
}

func (s *service) FindDocuments(ctx context.Context, userID string) ([]Document, error) {
	return []Document{{ID: "ID", UserID: "admin", Content: "toto"}}, nil
}

func (s *service) FindDocumentsByID(ctx context.Context, userID string, ids []string) ([]Document, error) {
	return nil, nil
}

func (s *service) ReplaceDocumentByID(ctx context.Context, userID string, id string, document *Document) (*Document, error) {
	return nil, nil
}

func (s *service) DeleteDocumentsByID(ctx context.Context, userID string, ids []string) ([]Document, error) {
	return nil, nil
}
