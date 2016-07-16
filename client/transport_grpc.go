package client

import (
	"errors"

	"github.com/solher/kit-crud/pb"
	"golang.org/x/net/context"
)

func encodeGRPCCreateDocumentRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(CreateDocumentRequest)
	return &pb.CreateDocumentRequest{
		UserId:   req.UserID,
		Document: toPBDocument(req.Document),
	}, nil
}

func decodeGRPCCreateDocumentResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(*pb.CreateDocumentReply)
	return CreateDocumentResponse{
		Document: toDocument(res.Document),
		Err:      toError(res.Err),
	}, nil
}

func encodeGRPCFindDocumentsRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(FindDocumentsRequest)
	return &pb.FindDocumentsRequest{
		UserId: req.UserID,
	}, nil
}

func decodeGRPCFindDocumentsResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(*pb.FindDocumentsReply)
	return FindDocumentsResponse{
		Documents: toDocuments(res.Documents),
		Err:       toError(res.Err),
	}, nil
}

func encodeGRPCFindDocumentsByIDRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(FindDocumentsByIDRequest)
	return &pb.FindDocumentsByIdRequest{
		UserId: req.UserID,
		Ids:    req.IDs,
	}, nil
}

func decodeGRPCFindDocumentsByIDResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(*pb.FindDocumentsByIdReply)
	return FindDocumentsByIDResponse{
		Documents: toDocuments(res.Documents),
		Err:       toError(res.Err),
	}, nil
}

func encodeGRPCReplaceDocumentByIDRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(ReplaceDocumentByIDRequest)
	return &pb.ReplaceDocumentByIdRequest{
		UserId: req.UserID,
		Id:     req.ID,
	}, nil
}

func decodeGRPCReplaceDocumentByIDResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(*pb.ReplaceDocumentByIdReply)
	return ReplaceDocumentByIDResponse{
		Document: toDocument(res.Document),
		Err:      toError(res.Err),
	}, nil
}

func encodeGRPCDeleteDocumentsByIDRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(DeleteDocumentsByIDRequest)
	return &pb.DeleteDocumentsByIdRequest{
		UserId: req.UserID,
		Ids:    req.IDs,
	}, nil
}

func decodeGRPCDeleteDocumentsByIDResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(*pb.DeleteDocumentsByIdReply)
	return DeleteDocumentsByIDResponse{
		Documents: toDocuments(res.Documents),
		Err:       toError(res.Err),
	}, nil
}

func toError(err string) error {
	if len(err) == 0 {
		return nil
	}
	return errors.New(err)
}

func toPBDocument(document *Document) *pb.Document {
	return toPBDocuments([]Document{*document})[0]
}

func toPBDocuments(documents []Document) []*pb.Document {
	pbDocuments := make([]*pb.Document, len(documents))
	for i, n := range documents {
		pbDocuments[i] = &pb.Document{
			Id:      n.ID,
			UserId:  n.UserID,
			Content: n.Content,
		}
	}
	return pbDocuments
}

func toDocument(pbDocument *pb.Document) *Document {
	return &toDocuments([]*pb.Document{pbDocument})[0]
}

func toDocuments(pbDocuments []*pb.Document) []Document {
	documents := make([]Document, len(pbDocuments))
	for i, n := range pbDocuments {
		documents[i] = Document{
			ID:      n.Id,
			UserID:  n.UserId,
			Content: n.Content,
		}
	}
	return documents
}
