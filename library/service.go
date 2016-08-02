package library

import (
	"time"

	"golang.org/x/net/context"
)

type (
	DatabaseRunner interface {
		Run(ctx context.Context, query string) (string, error)
	}

	Service interface {
		CreateDocument(ctx context.Context, userID string, document *Document) (*Document, error)
		FindDocuments(ctx context.Context, userID string) ([]Document, error)
		FindDocumentsByID(ctx context.Context, userID string, ids []string) ([]Document, error)
		ReplaceDocumentByID(ctx context.Context, userID string, id string, document *Document) (*Document, error)
		DeleteDocumentsByID(ctx context.Context, userID string, ids []string) ([]Document, error)
	}
)

type service struct {
	store DatabaseRunner
}

func NewService(store DatabaseRunner) *service {
	return &service{
		store: store,
	}
}

func (s *service) CreateDocument(ctx context.Context, userID string, document *Document) (*Document, error) {
	return nil, nil
}

func (s *service) FindDocuments(ctx context.Context, userID string) ([]Document, error) {
	time.Sleep(100 * time.Microsecond)
	if _, err := s.store.Run(ctx, "SELECT * FROM documents;"); err != nil {
		return nil, Unexpected(err)
	}
	time.Sleep(100 * time.Microsecond)
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
