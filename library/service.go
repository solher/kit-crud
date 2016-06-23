package library

type Service interface {
	CreateDocument(userID string, document *Document) (*Document, error)
	FindDocuments(userID string) ([]Document, error)
	FindDocumentsByID(userID string, ids []string) ([]Document, error)
	ReplaceDocumentByID(userID string, id string, document *Document) (*Document, error)
	DeleteDocumentsByID(userID string, ids []string) ([]Document, error)
}

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) CreateDocument(userID string, document *Document) (*Document, error) {
	return nil, nil
}

func (s *service) FindDocuments(userID string) ([]Document, error) {
	return []Document{{ID: "ID", UserID: "admin", Content: "toto"}}, nil
}

func (s *service) FindDocumentsByID(userID string, ids []string) ([]Document, error) {
	return nil, nil
}

func (s *service) ReplaceDocumentByID(userID string, id string, document *Document) (*Document, error) {
	return nil, nil
}

func (s *service) DeleteDocumentsByID(userID string, ids []string) ([]Document, error) {
	return nil, nil
}
