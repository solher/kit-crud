package client

import (
	"github.com/solher/kit-crud/pb"
	"golang.org/x/net/context"
)

type Service interface {
	CreateDocument(ctx context.Context, userID string, document *pb.Document) (*pb.Document, error)
	FindDocuments(ctx context.Context, userID string) ([]*pb.Document, error)
	FindDocumentsByID(ctx context.Context, userID string, ids []string) ([]*pb.Document, error)
	ReplaceDocumentByID(ctx context.Context, userID string, id string, document *pb.Document) (*pb.Document, error)
	DeleteDocumentsByID(ctx context.Context, userID string, ids []string) ([]*pb.Document, error)
}
