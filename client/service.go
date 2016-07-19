package client

import "github.com/solher/kit-crud/pb"

type Service interface {
	CreateDocument(userID string, document *pb.Document) (*pb.Document, error)
	FindDocuments(userID string) ([]*pb.Document, error)
	FindDocumentsByID(userID string, ids []string) ([]*pb.Document, error)
	ReplaceDocumentByID(userID string, id string, document *pb.Document) (*pb.Document, error)
	DeleteDocumentsByID(userID string, ids []string) ([]*pb.Document, error)
}
