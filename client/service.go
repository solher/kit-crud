package client

type Service interface {
	CreateDocument(userID string, document *Document) (*Document, error)
	FindDocuments(userID string) ([]Document, error)
	FindDocumentsByID(userID string, ids []string) ([]Document, error)
	ReplaceDocumentByID(userID string, id string, document *Document) (*Document, error)
	DeleteDocumentsByID(userID string, ids []string) ([]Document, error)
}
